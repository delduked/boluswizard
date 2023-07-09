package setup

type activeInsulin struct {
	Duration string
}

type History struct {
	Bolus float64
	Time  int64
}

var ActiveInsulin = activeInsulin{
	Duration: "02h00m",
}

var BolusHistory = []History{
	{
		Bolus: 10.0,
		Time:  1652115300000000000,
		//     1351700038292387000
		//     1652155200000000000
	},
	{
		Bolus: 10.0,
		Time:  1652118420000000000, //1:47pm
		//     1351700038292387000
	},
	{
		Bolus: 10.0,
		Time:  1652118000000000000, //1:40pm
		//     1351700038292387000
	},
	{
		Bolus: 10.0,
		Time:  1652118420000000000, //1:47pm
		//     1351700038292387000
	},
}
