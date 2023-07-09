package controllers

import (
	"api/services"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	log.Println("Context:", fmt.Sprintf("%v", c))
	go services.ErrorLog()
	return c.Next()

}
