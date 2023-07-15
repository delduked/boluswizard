package services

import (
	"api/config"
	"api/types"
	"time"

	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// Save a single or multiple ISF values for a specific user
func SaveISFs(value []types.InsulinSensitivity, uid string) error {

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

	// get all targets associated with logged in user
	isfs, err := ISFs(uid)
	if err != nil {
		return err
	}

	// delete all targets associated with logged in user
	err = deleteISFs(isfs, uid)
	if err != nil {
		return err
	}

	for _, j := range value {

		key = (uuid.New()).String()
		field = uid + "::" + key + "::isf"

		_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {

			rdb.HSet(config.RedisCtx, field, "Key", key)
			rdb.HSet(config.RedisCtx, field, "Start", j.Start)
			rdb.HSet(config.RedisCtx, field, "End", j.End)
			rdb.HSet(config.RedisCtx, field, "Sensitivity", j.Sensitivity)

			return nil
		})
		if err != nil {
			ErrorLogger <- err
			return err
		}
	}

	return err
}

// Get all ISF values for a specific user
func ISFs(uid string) ([]types.InsulinSensitivity, error) {

	savedISF := []types.InsulinSensitivity{}
	var keyedField types.InsulinSensitivity

	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		ErrorLogger <- err
		return savedISF, err
	}

	for _, j := range length {

		parse := strings.Split(j, "::")

		if len(parse) == 3 {
			if parse[0] == uid && parse[2] == "isf" {
				err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&keyedField)
				if err != nil {
					ErrorLogger <- err
					return savedISF, err
				}
				savedField := types.InsulinSensitivity{
					Key:         parse[1],
					Start:       keyedField.Start,
					End:         keyedField.End,
					Sensitivity: keyedField.Sensitivity,
				}
				savedISF = append(savedISF, savedField)
			}
		}
	}

	return savedISF, err
}

// Delete all ISF values for a specific user
func deleteISFs(value []types.InsulinSensitivity, uid string) error {

	for _, j := range value {

		err := config.Rdb.Del(config.RedisCtx, uid+"::"+j.Key+"::isf")
		if err.Err() != nil {
			ErrorLogger <- err
			return err.Err()
		}
	}

	return nil
}

// Get the current Insulin sensitivity factor for the current user
func CurrentIsf(uid string) (types.InsulinSensitivity, error) {
	var res types.InsulinSensitivity
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	isf, err := GetInsulinSensitivityFactor(t, now, uid)
	if err != nil {
		return res, err
	}
	res.Sensitivity = RoundFloat(isf,2)
	return res, err
}
