package config

import "os"

var Secret = getSecret()

func getSecret() []byte {
	importApiPort, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return []byte("default-secret-here")
	}
	return []byte(importApiPort)
}
