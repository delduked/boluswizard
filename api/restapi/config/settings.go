package config

import (
	"os"
)

var ApiPort = getApiPort()

func getApiPort() string {
	importApiPort, ok := os.LookupEnv("API_PORT")
	if !ok {
		return "8080"
	}
	return ":" + importApiPort
}
