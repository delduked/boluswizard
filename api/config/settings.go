package config

import (
	"api/types"
	"os"
)

var ApiPort = getApiPort()
var ActiveInsulin = types.ActiveInsulin{
	Duration: "02h00m",
}

func getApiPort() string {
	importApiPort, ok := os.LookupEnv("API_PORT")
	if !ok {
		return "8080"
	}
	return ":" + importApiPort
}
