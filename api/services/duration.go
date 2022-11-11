package services

import (
	"api/config"
	"api/types"

	"github.com/go-redis/redis/v8"
)

func SaveDuration(uid string, value *types.ActiveInsulinDuration) error {

	var err error

	// check if user input is valid
	err = ValidateUserInput(value.Duration)
	if err != nil {
		return err
	}

	err = deleteDuration(uid)
	if err != nil {
		return err
	}

	// Set field.
	_, err = config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		field := uid + "::duration"
		rdb.HSet(config.RedisCtx, field, "Duration", value.Duration)
		return nil
	})

	if err != nil {
		ErrorLogger <- err
		return err
	}

	return err
}

// Get value of active insulin duration from specific user
func GetDuration(uid string) (types.ActiveInsulinDuration, error) {
	var keyedField types.ActiveInsulinDuration

	data := config.Rdb.HGet(config.RedisCtx, uid+"::duration", "Duration")
	if data.Err() != nil {
		return keyedField, data.Err()
	}

	keyedField.Duration = data.Val()
	return keyedField, nil
}

// Delete active insulin duration saved for a specific user
func deleteDuration(uid string) error {

	err := config.Rdb.Del(config.RedisCtx, uid+"::duration")
	if err.Err() != nil {
		ErrorLogger <- err
		return err.Err()
	}

	return nil
}
