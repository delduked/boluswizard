package services

import (
	"api/config"
	"api/types"
	"fmt"
	"time"

	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func SaveCorrection(value *types.Correction, uid string) error {
	var err error
	key := (uuid.New()).String()

	// Set some fields.
	_, err = config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		field := uid + "::" + key + "::correction"
		value.Key = key
		rdb.HSet(config.RedisCtx, field, "Key", key)
		rdb.HSet(config.RedisCtx, field, "Bg", value.Bg)
		rdb.HSet(config.RedisCtx, field, "Carbs", value.Carbs)
		rdb.HSet(config.RedisCtx, field, "Bolus", value.Bolus)
		rdb.HSet(config.RedisCtx, field, "TimeStamp", time.Now().UnixNano())
		return nil
	})

	if err != nil {
		ErrorLogger <- err
		return err
	}

	return err
}

// Get a single or multiple corrections specific to the logged in user
func Corrections(uid string) ([]types.Correction, error) {

	savedCorrections := []types.Correction{}
	var keyedField types.Correction

	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		ErrorLogger <- err
		return savedCorrections, err
	}

	for _, j := range length {

		parse := strings.Split(j, "::")

		if len(parse) == 3 {
			if parse[0] == uid && parse[2] == "correction" {
				err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&keyedField)
				if err != nil {
					ErrorLogger <- err
					return savedCorrections, err
				}
				savedField := types.Correction{
					Key:       parse[1],
					Bg:        keyedField.Bg,
					Carbs:     keyedField.Carbs,
					Bolus:     keyedField.Bolus,
					TimeStamp: keyedField.TimeStamp,
				}
				savedCorrections = append(savedCorrections, savedField)
			}
		}
	}

	SortCorrection(savedCorrections)

	return savedCorrections, err
}

func Correction(uid string, key string) (types.Correction, error) {

	var keyedField types.Correction

	err := config.Rdb.HGetAll(config.RedisCtx, uid+"::"+key+"::correction").Scan(&keyedField)
	if err != nil {
		ErrorLogger <- err
		return keyedField, err
	}
	return keyedField, nil
}

// Save a single or multiple corrections
func SaveCorrections(value []types.Correction, uid string) error {
	fmt.Println("value")
	fmt.Println(value)
	var err error
	var key string

	for _, j := range value {

		key = (uuid.New()).String()
		field := uid + "::" + key + "::correction"

		_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {

			rdb.HSet(config.RedisCtx, field, "Key", key)
			rdb.HSet(config.RedisCtx, field, "Bg", j.Bg)
			rdb.HSet(config.RedisCtx, field, "Carbs", j.Carbs)
			rdb.HSet(config.RedisCtx, field, "Bolus", j.Bolus)
			rdb.HSet(config.RedisCtx, field, "TimeStamp", time.Now().UnixNano())

			return nil
		})
		if err != nil {
			ErrorLogger <- err
			return err
		}
	}

	return err

}

// function to get a range ocrrections from the current logged in user between a specified period of time.
func CorrectionRange(req types.CorrectionRange, uid string) ([]types.Correction, error) {
	var result []types.Correction
	var err error

	// Get entire list of corrections for the specific user
	corrections, err := Corrections(uid)
	if err != nil {
		ErrorLogger <- err
		return result, err
	}

	for _, j := range corrections {

		// determine if current correction value is in between the requested unix nano times
		above := j.TimeStamp >= req.Start
		below := j.TimeStamp <= req.End

		if above && below {
			result = append(result, j)
		}
	}
	return result, err
}
