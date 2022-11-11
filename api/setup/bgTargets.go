package setup

import "api/types"

type Target struct {
	Start string
	End   string
	High  float64
	Low   float64
}

var BGTargets = []types.Target{
	{
		Start: "00h00m",
		End:   "05h59m",
		High:  6.6,
		Low:   5.4,
	},
	{
		Start: "06h00m",
		End:   "17h59m",
		High:  5.7,
		Low:   5.2,
	},
	{
		Start: "18h00m",
		End:   "23h59m",
		High:  6.6,
		Low:   5.4,
	},
}
