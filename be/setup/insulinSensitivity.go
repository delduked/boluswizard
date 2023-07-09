package setup

type InsulinSensitivity struct {
	Start       string  `json:"Start" redis:"Start"`
	End         string  `json:"End" redis:"End"`
	Sensitivity float64 `json:"Sensitivity" redis:"Sensitivity"`
}

var ISF = []InsulinSensitivity{
	{
		Start:       "00h00m",
		End:         "05h59m",
		Sensitivity: 1 / 1.32,
	},
	{
		Start:       "06h00m",
		End:         "23h59m",
		Sensitivity: 1 / 1.4,
	},
}
