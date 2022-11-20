package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var API_SERVICE = serviceUrl()

func serviceUrl() string {
	publicorlocal, ok := getServiceHostname()
	if !ok {
		var apiServicePort = getServicePortNumber()
		return publicorlocal[:len(publicorlocal)-1] + ":" + apiServicePort + "/"
	}
	return publicorlocal
}

func getServiceHostname() (string, bool) {
	importHostname, ok := os.LookupEnv("API_HOSTNAME")
	if !ok {
		var myEnv map[string]string
		var err error
		myEnv, err = godotenv.Read("../.env")
		if err != nil {
			fmt.Println("Error reading .env file")
			return "http://localhost/", false
		}
		return myEnv["API_HOSTNAME"], false
	}
	return importHostname, true
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
