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
		fmt.Println("redis host: ", myEnv["REDIS_HOSTNAME"])
		return myEnv["REDIS_HOSTNAME"]
	}
	fmt.Println("redis host: ", importHostname)
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
			return "default-redis-password"
		}
		fmt.Println("redis pass: ", myEnv["REDIS_PASSWORD"])
		return myEnv["REDIS_PASSWORD"]
	}
	fmt.Println("redis pass: ", importRedisPassword)
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
		fmt.Println("redis port: ", myEnv["REDIS_PORT"])
		return myEnv["REDIS_PORT"]
	}
	fmt.Println("redis port: ", importRedisPort)
	return importRedisPort
}
