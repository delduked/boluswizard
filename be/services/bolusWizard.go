package services

import (
	"api/types"
	"fmt"
	"math"
	"time"
)

// Calculate correction bolus
func BolusWizard(value *types.CorrectionRequest, uid string) (types.CorrectionResponse, error) {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	var errors = make(chan error)
	var done = make(chan bool) // semaphore

	var bgTarget types.Target
	var isf float64
	var carbRatio float64
	var activeInsulinDuration types.ActiveInsulinDuration

	// get the list of BG Targets for the specific user
	go func() {
		var err error
		bgTarget, err = GetBGTarget(t, now, uid)
		errors <- err
		done <- true
	}()

	// get all the insulin sensitivity factor value for the specific user through out the day
	go func() {
		var err error
		isf, err = GetInsulinSensitivityFactor(t, now, uid)
		errors <- err
		done <- true
	}()

	// get the list of all carb ratios through out the day for the specific user
	go func() {
		var err error
		carbRatio, err = GetCarbRatio(t, now, uid)
		errors <- err
		done <- true
	}()

	// get active insulin duration value for the specific user
	go func() {
		var err error
		activeInsulinDuration, err = GetDuration(uid)
		errors <- err
		done <- true
	}()

	go func() {
		<-done // semaphore
		<-done // semaphore
		<-done // semaphore
		<-done // semaphore
		close(errors)
	}()

	// check if any error were placed on the error channel, if so respond with an error message
	for n := range errors {
		if n != nil {
			ErrorLogger <- n
			return types.CorrectionResponse{}, n
		}
	}

	// calculate the blood sugar correction for the specific user based on the save bg targets and insulin sensitivity factor through the day
	bgCorrection := getBgCorrection(value.Bg, bgTarget, isf)

	// calculdate the carb correction based on the carb ratios saved for the specific user
	carbCorrection := getCarbCorrection(value.Carbs, carbRatio)

	// calculate the amount of active insulin based on the users bolus history, insulin sensitivity factor and active inuslin duration
	activeInsulin, err := getActiveInsulin(t, now, uid, activeInsulinDuration.Duration)
	if err != nil {
		ErrorLogger <- err
		return types.CorrectionResponse{}, err
	}

	// round the value to two decimal places
	totalCorrection := RoundFloat(carbCorrection+bgCorrection-activeInsulin, 2)
	if totalCorrection < 0 {
		totalCorrection = 0
	}

	// construct the response
	res := types.CorrectionResponse{
		BgCorrection:           RoundFloat(bgCorrection, 2),
		CarbCorrection:         RoundFloat(carbCorrection, 2),
		ActiveInsulinReduction: RoundFloat(activeInsulin, 2),
		Bolus:                  RoundFloat(totalCorrection, 2),
	}

	fmt.Printf(`Bolus Wizard correction: %v`, res)
	return res, err

}

func unixTime(t time.Time) {

	unixTime := t.UnixNano()
	fmt.Println("")
	fmt.Println("unixTime: ", unixTime)

}
func GetBGTarget(t time.Time, now time.Time, uid string) (types.Target, error) {
	var result types.Target
	var err error

	targets, err := Targets(uid)

	for _, j := range targets {

		start, err := time.ParseDuration(j.Start)
		if err != nil {
			return result, err
		}
		end, err := time.ParseDuration(j.End)
		if err != nil {
			return result, err
		}

		// floor add startnanoseconds
		nanoFloorStart := t.UnixNano() + start.Nanoseconds()
		nanoFloorEnd := t.UnixNano() + end.Nanoseconds()

		// current time
		nanoCurrent := now.UnixNano()
		above := nanoFloorStart <= nanoCurrent
		below := nanoFloorEnd >= nanoCurrent

		if above && below {
			result = j
			return result, err
		}
	}
	err = fmt.Errorf("No BG Target found")
	return result, err
}
func GetCarbRatio(t time.Time, now time.Time, uid string) (float64, error) {
	var result float64
	var err error

	ratios, err := Ratios(uid)

	for _, j := range ratios {

		start, err := time.ParseDuration(j.Start)
		if err != nil {
			return result, err
		}
		end, err := time.ParseDuration(j.End)
		if err != nil {
			return result, err
		}

		// current date starting at 00:00:00 in unix nanoseconds
		nanoFloorStart := t.UnixNano() + start.Nanoseconds()
		nanoFloorEnd := t.UnixNano() + end.Nanoseconds()

		// current time
		nanoCurrent := now.UnixNano()

		above := nanoFloorStart <= nanoCurrent
		below := nanoFloorEnd >= nanoCurrent

		if above && below {
			result = j.Ratio
			return result, err
		}
	}
	err = fmt.Errorf("No Carb Ratio found")
	return result, err
}

func GetInsulinSensitivityFactor(t time.Time, now time.Time, uid string) (float64, error) {
	var result float64
	var err error

	isf, err := ISFs(uid)
	if err != nil {
		return result, err
	}

	for _, j := range isf {

		start, err := time.ParseDuration(j.Start)
		if err != nil {
			ErrorLogger <- err
			return result, err
		}
		end, err := time.ParseDuration(j.End)
		if err != nil {
			ErrorLogger <- err
			return result, err
		}

		// floor add startnanoseconds
		nanoFloorStart := t.UnixNano() + start.Nanoseconds()
		nanoFloorEnd := t.UnixNano() + end.Nanoseconds()

		// current time
		nanoCurrent := now.UnixNano()

		above := nanoFloorStart <= nanoCurrent
		below := nanoFloorEnd >= nanoCurrent

		if above && below {
			result = j.Sensitivity
			return result, err
		}
	}
	err = fmt.Errorf("No ISF found")
	return result, err
}
func getBgCorrection(bg float64, bgTarget types.Target, isf float64) float64 {
	var result float64

	above := bg >= bgTarget.High
	below := bg <= bgTarget.Low
	inBetween := bg >= bgTarget.Low && bg <= bgTarget.High

	if above {
		result = isf * (bg - bgTarget.High)
	}
	if below {
		result = -1 * (isf * (bgTarget.Low - bg))
	}
	if inBetween {
		result = 0
	}
	return result
}

func getCarbCorrection(carbs float64, carbRatio float64) float64 {
	result := carbs * carbRatio
	return result
}
func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func getActiveInsulin(t time.Time, now time.Time, uid string, duration string) (float64, error) {
	// look through array of previous boluses
	// usualy active insulin lasts 2 hours, so look through the array
	// and return value which are 2 hours before the current time
	// Find the difference between the current time and when the
	// bolus was given. Divide the difference by two hours
	// to get the percentage of the remain insulin to be effect
	// the user. For al the boluses which happened with in the last 2 hours
	// get the sum of the remaining percentages multipled by the bolus at
	// current time.
	// ex: sum = (12u * (120min - (current time - bolus time))/120min))) + (5u * 0.4)

	var remainingActiveInsulin float64
	remainingActiveInsulin = 0.0
	var err error

	userCorrectionHistory, err := Corrections(uid)
	if err != nil {
		return remainingActiveInsulin, err
	}

	for _, j := range userCorrectionHistory {

		nanoCurrentTime := now.UnixNano()
		bolusTime := j.TimeStamp

		activeInsulinDuration, err := time.ParseDuration(duration)
		if err != nil {
			ErrorLogger <- err
			return remainingActiveInsulin, err
		}

		usedDurationOfActiveInsulin := nanoCurrentTime - bolusTime

		if usedDurationOfActiveInsulin < activeInsulinDuration.Nanoseconds() {

			activeInsulinFormula := calculateRemainingActiveInsulin(
				j.Bolus,
				float64(usedDurationOfActiveInsulin),
				float64(activeInsulinDuration.Nanoseconds()),
			)
			remainingActiveInsulin = remainingActiveInsulin + activeInsulinFormula
		}

	}
	return remainingActiveInsulin, err

}

func calculateRemainingActiveInsulin(bolus float64, usedDurationOfActiveInsulin float64, activeInsulinDuration float64) float64 {

	a := activeInsulinDuration - usedDurationOfActiveInsulin
	b := a / activeInsulinDuration
	result := bolus * b
	return result
}
