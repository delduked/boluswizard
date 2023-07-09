package services

import (
	"api/config"
	"api/types"
	"fmt"

	"strings"

	"github.com/go-redis/redis/v8"
)

// Check if the user exists in the redis database
func UsersExists(username string) error {

	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		ErrorLogger <- err
		return err
	}

	// Users only have two identifiers in the field name
	// the Uid and the type ex: 345s-vda54-at4e-a22h::username
	for _, j := range length {

		parse := strings.Split(j, "::")

		if len(parse) == 2 && parse[1] == "username" {

			data := config.Rdb.HGet(config.RedisCtx, j, "Username")
			if data.Err() != nil {
				ErrorLogger <- data.Err()
				return data.Err()
			}

			if data.Val() == username {
				return fmt.Errorf("User account already exists.")
			}
		}
	}

	return err
}

// Verifiy if the username has asking to login is passing the correct password associated with the user attempting to login
func VerifyUsernameAndPassword(user types.Users) (bool, types.Users, error) {

	var keyedField types.Users
	var verifiedUsers types.Users

	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		ErrorLogger <- err
		return false, verifiedUsers, err
	}

	// Users only have two identifiers in the field name
	// the Uid and the type ex: 345s-vda54-at4e-a22h::username

	for _, j := range length {

		parse := strings.Split(j, "::")

		if len(parse) == 2 {
			if parse[1] == "username" {
				err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&keyedField)
				if err != nil {
					ErrorLogger <- err
					return false, verifiedUsers, err
				}

				if keyedField.Username == user.Username && keyedField.Password == user.Password {
					verifiedUsers = keyedField
					verifiedUsers.Uid = parse[0]
					return true, verifiedUsers, err
				}
			}
		}
	}

	return false, verifiedUsers, fmt.Errorf("Incorrect username or password")
}

// Get list of all users saved in the redis database
func User(uid string) (types.Users, error) {

	// Users only have two identifiers in the field name
	// the Uid and the type ex: 345s-vda54-at4e-a22h::username

	var keyedField types.Users

	err := config.Rdb.HGetAll(config.RedisCtx, uid+"::username").Scan(&keyedField)
	if err != nil {
		ErrorLogger <- err
		return keyedField, err
	}

	return keyedField, err
}

// Allow a new user to sign up
func SaveUser(value *types.Users) error {
	var err error

	// Set some fields.
	_, err = config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		field := value.Uid + "::username"
		rdb.HSet(config.RedisCtx, field, "Username", value.Username)
		rdb.HSet(config.RedisCtx, field, "Password", value.Password)
		return nil
	})

	if err != nil {
		ErrorLogger <- err
		return err
	}

	return err

}
