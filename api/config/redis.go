package config

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
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
		return "192.168.0.32"
	}
	return importHostname

}

func getRedisPassword() string {
	importRedisPassword, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		return "n4th4n43l"
	}
	return importRedisPassword
}

func getRedisPort() string {
	importRedisPort, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		return "6399"
	}
	return importRedisPort
}
