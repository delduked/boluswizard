package services

import (
	"api/config"
	"api/types"
	"time"

	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// Save a single or multiple Carb ratios for a specific user
func SaveRatios(value []types.CarbRatio, uid string) error {

	var err error
	var key, field string

	// check if user input is valid
	for _, j := range value {
		err = ValidateUserInput(j.Start)
		if err != nil {
			return err
		}
		err = ValidateUserInput(j.End)
		if err != nil {
			return err
		}
	}

	// get all ratios associated with logged in user
	ratios, err := Ratios(uid)
	if err != nil {
		return err
	}

	// delete all ratios associated with logged in user
	err = deleteRatios(ratios, uid)
	if err != nil {
		return err
	}

	// remakte all ratios associated with the logged in user
	for _, j := range value {

		key = (uuid.New()).String()
		field = uid + "::" + key + "::ratio"

		_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {

			rdb.HSet(config.RedisCtx, field, "Key", key)
			rdb.HSet(config.RedisCtx, field, "Start", j.Start)
			rdb.HSet(config.RedisCtx, field, "End", j.End)
			rdb.HSet(config.RedisCtx, field, "Ratio", j.Ratio)

			return nil
		})
		if err != nil {
			return err
		}
	}

	return err
}

// Get all BG targets for a specific user
func Ratios(uid string) ([]types.CarbRatio, error) {

	savedRatios := []types.CarbRatio{}
	var keyedField types.CarbRatio

	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		ErrorLogger <- err
		return savedRatios, err
	}

	for _, j := range length {

		parse := strings.Split(j, "::")

		if len(parse) == 3 {
			if parse[0] == uid && parse[2] == "ratio" {
				err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&keyedField)
				if err != nil {
					ErrorLogger <- err
					return savedRatios, err
				}
				savedField := types.CarbRatio{
					Key:   parse[1],
					Start: keyedField.Start,
					End:   keyedField.End,
					Ratio: keyedField.Ratio,
				}
				savedRatios = append(savedRatios, savedField)
			}
		}
	}

	return savedRatios, err
}

// Delete all carb ratios for a specific user
func deleteRatios(value []types.CarbRatio, uid string) error {

	for _, j := range value {

		err := config.Rdb.Del(config.RedisCtx, uid+"::"+j.Key+"::ratio")
		if err.Err() != nil {
			ErrorLogger <- err
			return err.Err()
		}
	}

	return nil
}

func CurrentRatio(uid string) (types.CarbRatio, error) {
	now := time.Now()
	var res types.CarbRatio
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	carbRatio, err := GetCarbRatio(t, now, uid)
	if err != nil {
		return res, err
	}
	res.Ratio = RoundFloat(carbRatio, 2)
	return res, err
}
