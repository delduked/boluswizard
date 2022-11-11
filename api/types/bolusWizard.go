package types

type CorrectionResponse struct {
	BgCorrection           float64 `json:"BgCorrection" redis:"BgCorrection"`
	CarbCorrection         float64 `json:"CarbCorrection" redis:"CarbCorrection"`
	ActiveInsulinReduction float64 `json:"ActiveInsulinReduction" redis:"ActiveInsulinReduction"`
	Bolus                  float64 `json:"Bolus" redis:"Bolus"`
}
type CorrectionRequest struct {
	Bg    float64 `json:"Bg" redis:"Bg"`
	Carbs float64 `json:"Carbs" redis:"Carbs"`
}
type Correction struct {
	Key       string  `json:"Key" redis:"Key"`
	Bg        float64 `json:"Bg" redis:"Bg"`
	Carbs     float64 `json:"Carbs" redis:"Carbs"`
	Bolus     float64 `json:"Bolus" redis:"Bolus"`
	TimeStamp int64   `json:"TimeStamp" redis:"TimeStamp"`
}
type ActiveInsulinDuration struct {
	Duration string `json:"Duration" redis:"Duration"`
}
type CorrectionRange struct {
	Start int64 `json:"Start" redis:"Start"`
	End   int64 `json:"End" redis:"End"`
}
