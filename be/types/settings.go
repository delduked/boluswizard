package types

type Target struct {
	Key   string  `json:"Key" redis:"Key"`
	End   string  `json:"End" redis:"End"`
	Start string  `json:"Start" redis:"Start"`
	High  float64 `json:"High" redis:"High"`
	Low   float64 `json:"Low" redis:"Low"`
}

type ActiveInsulin struct {
	Key      string `json:"Key" redis:"Key"`
	Duration string `json:"Duration" redis:"Duration"`
}

type CarbRatio struct {
	Key   string  `json:"Key" redis:"Key"`
	End   string  `json:"End" redis:"End"`
	Start string  `json:"Start" redis:"Start"`
	Ratio float64 `json:"Ratio" redis:"Ratio"`
}
type InsulinSensitivity struct {
	Key         string  `json:"Key" redis:"Key"`
	End         string  `json:"End" redis:"End"`
	Start       string  `json:"Start" redis:"Start"`
	Sensitivity float64 `json:"Sensitivity" redis:"Sensitivity"`
}
