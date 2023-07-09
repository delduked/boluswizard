package services

import (
	"time"
)

func ValidateUserInput(input string) error {
	_, err := time.ParseDuration(input)
	if err != nil {
		ErrorLogger <- err
		return err
	}
	return err
}
