package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var redisHostname = getRedisHostname()
var redisPassword = getRedisPassword()
var redisPort = getRedisPort()

var RedisCtx = context.Background()
var Rdb = redis.NewClient(&redis.Options{
	Addr:     redisHostname + ":" + redisPort,
	Password: redisPassword, // no password set
	DB:       0,             // use default DB
})

func getRedisHostname() string {
	importHostname, ok := os.LookupEnv("REDIS_HOSTNAME")
	if !ok {
		var myEnv map[string]string
		var err error
		myEnv, err = godotenv.Read("../.env")
		if err != nil {
			fmt.Println("Error reading .env file")
			return "localhost"
		}
		return myEnv["REDIS_HOSTNAME"]
	}
	return importHostname

}

func getRedisPassword() string {
	importRedisPassword, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		var myEnv map[string]string
		var err error
		myEnv, err = godotenv.Read("../.env")
		if err != nil {
			fmt.Println("Error reading .env file")
			return "n4th4n43l"
		}
		return myEnv["REDIS_PASSWORD"]
	}
	return importRedisPassword
}
func getRedisPort() string {
	importRedisPort, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		var myEnv map[string]string
		var err error
		myEnv, err = godotenv.Read("../.env")
		if err != nil {
			fmt.Println("Error reading .env file")
			return "6399"
		}
		return myEnv["REDIS_PORT"]
	}
	return importRedisPort
}
