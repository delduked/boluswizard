package setup

type carbRatio struct {
	Start string
	End   string
	Ratio float64
}

var Ratios = []carbRatio{
	{
		Start: "00h00m",
		End:   "11h29m",
		Ratio: 1 / 8.0,
	},
	{
		Start: "11h30m",
		End:   "15h59m",
		Ratio: 1 / 6.5,
	},
	{
		Start: "16h00m",
		End:   "23h59m",
		Ratio: 1 / 8.0,
	},
}
