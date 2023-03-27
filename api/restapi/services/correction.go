package services

import (
	"boluswizard/models"
	"boluswizard/restapi/config"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func BolusWizard(bg, carbs float64, uid string) (models.CorrectionResponse, error) {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	var errors = make(chan error)
	var done = make(chan bool) // semaphore

	var bgTarget models.Target
	var isf float64
	var carbRatio float64
	var activeInsulinDuration models.ActiveInsulinDuration

	// get the list of BG Targets for the specific user
	go func() {
		var err error
		bgTarget, err = getBGTarget(t, now, uid)
		errors <- err
		done <- true
	}()

	// get all the insulin sensitivity factor value for the specific user through out the day
	go func() {
		var err error
		isf, err = getInsulinSensitivityFactor(t, now, uid)
		errors <- err
		done <- true
	}()

	// get the list of all carb ratios through out the day for the specific user
	go func() {
		var err error
		carbRatio, err = getCarbRatio(t, now, uid)
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
			return models.CorrectionResponse{}, n
		}
	}

	// calculate the blood sugar correction for the specific user based on the save bg targets and insulin sensitivity factor through the day
	bgCorrection := getBgCorrection(bg, bgTarget, isf)

	// calculdate the carb correction based on the carb ratios saved for the specific user
	carbCorrection := getCarbCorrection(carbs, carbRatio)

	// calculate the amount of active insulin based on the users bolus history, insulin sensitivity factor and active inuslin duration
	activeInsulin, err := getActiveInsulin(t, now, uid, activeInsulinDuration.Duration)
	if err != nil {
		ErrorLogger <- err
		return models.CorrectionResponse{}, err
	}

	// round the value to two decimal places
	totalCorrection := roundFloat(carbCorrection+bgCorrection-activeInsulin, 2)
	if *totalCorrection < 0 {
		*totalCorrection = 0
	}

	// construct the response
	res := models.CorrectionResponse{
		BgCorrection:           roundFloat(bgCorrection, 2),
		CarbCorrection:         roundFloat(carbCorrection, 2),
		ActiveInsulinReduction: roundFloat(activeInsulin, 2),
		Bolus:                  roundFloat(*totalCorrection, 2),
	}

	fmt.Printf(`Bolus Wizard correction: %v`, res)
	return res, err

}

func unixTime(t time.Time) {

	unixTime := t.UnixNano()
	fmt.Println("")
	fmt.Println("unixTime: ", unixTime)

}
func getBGTarget(t time.Time, now time.Time, uid string) (models.Target, error) {
	var result models.Target
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
func getCarbRatio(t time.Time, now time.Time, uid string) (float64, error) {
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

func getInsulinSensitivityFactor(t time.Time, now time.Time, uid string) (float64, error) {
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
func getBgCorrection(bg float64, bgTarget models.Target, isf float64) float64 {
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
func roundFloat(val float64, precision uint) *float64 {
	ratio := math.Pow(10, float64(precision))
	res := math.Round(val*ratio) / ratio
	return &res
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

// Get a single or multiple corrections specific to the logged in user
func Corrections(uid string) ([]*models.Correction, error) {

	savedCorrections := []*models.Correction{}
	var keyedField models.Correction

	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		ErrorLogger <- err
		return savedCorrections, err
	}

	for _, j := range length {

		parse := strings.Split(j, "::")

		if len(parse) == 3 {
			if parse[0] == uid && parse[2] == "correction" {
				err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&keyedField)
				if err != nil {
					ErrorLogger <- err
					return savedCorrections, err
				}
				savedField := &models.Correction{
					Key:       parse[1],
					Bg:        keyedField.Bg,
					Carbs:     keyedField.Carbs,
					Bolus:     keyedField.Bolus,
					TimeStamp: keyedField.TimeStamp,
				}
				savedCorrections = append(savedCorrections, savedField)
			}
		}
	}

	SortCorrection(savedCorrections)

	return savedCorrections, err
}

func calculateRemainingActiveInsulin(bolus float64, usedDurationOfActiveInsulin float64, activeInsulinDuration float64) float64 {

	a := activeInsulinDuration - usedDurationOfActiveInsulin
	b := a / activeInsulinDuration
	result := bolus * b
	return result
}

// function to get a range ocrrections from the current logged in user between a specified period of time.
func CorrectionRange(req models.CorrectionRange, uid string) ([]models.Correction, error) {
	var result []models.Correction
	var err error

	// Get entire list of corrections for the specific user
	corrections, err := Corrections(uid)
	if err != nil {
		ErrorLogger <- err
		return result, err
	}

	for _, j := range corrections {

		// determine if current correction value is in between the requested unix nano times
		above := j.TimeStamp >= req.Start
		below := j.TimeStamp <= req.End

		if above && below {
			result = append(result, *j)
		}
	}
	return result, err
}

// Save a single or multiple corrections
func SaveCorrections(value []models.Correction, uid string) error {

	var err error
	var key string

	for _, j := range value {

		key = (uuid.New()).String()
		field := uid + "::" + key + "::correction"

		_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {

			rdb.HSet(config.RedisCtx, field, "Key", key)
			rdb.HSet(config.RedisCtx, field, "Bg", j.Bg)
			rdb.HSet(config.RedisCtx, field, "Carbs", j.Carbs)
			rdb.HSet(config.RedisCtx, field, "Bolus", j.Bolus)
			rdb.HSet(config.RedisCtx, field, "TimeStamp", time.Now().UnixNano())

			return nil
		})
		if err != nil {
			ErrorLogger <- err
			return err
		}
	}

	return err

}
