package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var apiServiceName = getServiceHostname()
var apiServicePort = getServicePortNumber()
var API_SERVICE = apiServiceName[:len(apiServiceName)-1] + ":" + apiServicePort + "/"

func getServiceHostname() string {
	importHostname, ok := os.LookupEnv("API_HOSTNAME")
	if !ok {
		var myEnv map[string]string
		var err error
		myEnv, err = godotenv.Read("../.env")
		if err != nil {
			fmt.Println("Error reading .env file")
			return "http://localhost/"
		}
		return myEnv["API_HOSTNAME"]
	}
	return importHostname
}
func getServicePortNumber() string {
	importApiPort, ok := os.LookupEnv("API_PORT")
	if !ok {
		var myEnv map[string]string
		var err error
		myEnv, err = godotenv.Read("../.env")
		if err != nil {
			fmt.Println("Error reading .env file")
			return "6399"
		}
		return myEnv["API_PORT"]
	}
	return importApiPort
}
