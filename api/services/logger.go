package services

import (
	"fmt"
	"log"
)

var ErrorLogger = make(chan any)

func ErrorLog() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Defer:", fmt.Sprintf("%v", r))
		}
	}()

	for {
		err := <-ErrorLogger
		log.Println("Error:", fmt.Sprintf("%v", err))
	}
}
