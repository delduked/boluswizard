package services

import (
	"boluswizard/models"
	"boluswizard/restapi/config"

	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// Get a single BG target value associated to a logged in user
func Target(uid string, key string) (models.Target, error) {

	var keyedField models.Target

	err := config.Rdb.HGetAll(config.RedisCtx, uid+"::"+key+"::target").Scan(&keyedField)
	if err != nil {
		return keyedField, err
	}
	return keyedField, nil
}

// Get all the BG target values associated with a specific user
func Targets(uid string) ([]models.Target, error) {

	savedTargets := []models.Target{}
	var keyedField models.Target

	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		ErrorLogger <- err
		return savedTargets, err
	}

	for _, j := range length {

		parse := strings.Split(j, "::")

		if len(parse) == 3 {

			if parse[0] == uid && parse[2] == "target" {
				err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&keyedField)
				if err != nil {
					ErrorLogger <- err
					return savedTargets, err
				}
				savedField := models.Target{
					Key:   parse[1],
					Start: keyedField.Start,
					End:   keyedField.End,
					High:  keyedField.High,
					Low:   keyedField.Low,
				}
				savedTargets = append(savedTargets, savedField)
			}
		}
	}

	return savedTargets, err
}

// Save a single or multiple BG targets for a specific user
func SaveTargets(value []models.Target, uid string) error {

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
	targets, err := Targets(uid)
	if err != nil {
		return err
	}

	// delete all targets associated with logged in user
	err = deleteTargets(targets, uid)
	if err != nil {
		return err
	}

	// add all targets supplied by user
	for _, j := range value {

		key = (uuid.New()).String()
		field = uid + "::" + key + "::target"

		_, err = config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {

			rdb.HSet(config.RedisCtx, field, "Key", key)
			rdb.HSet(config.RedisCtx, field, "Start", j.Start)
			rdb.HSet(config.RedisCtx, field, "End", j.End)
			rdb.HSet(config.RedisCtx, field, "High", j.High)
			rdb.HSet(config.RedisCtx, field, "Low", j.Low)

			return nil
		})
		if err != nil {
			ErrorLogger <- err
			return err
		}
	}

	return err
}

// Delete all BG targets from a single user
func deleteTargets(value []models.Target, uid string) error {

	for _, j := range value {

		err := config.Rdb.Del(config.RedisCtx, uid+"::"+j.Key+"::target")
		if err.Err() != nil {
			ErrorLogger <- err
			return err.Err()
		}
	}

	return nil
}
