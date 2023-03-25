package services

import (
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
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
		log.Println("Log :", fmt.Sprintf("%v", err), strfmt.DateTime(time.Now().UTC()))
	}
}
