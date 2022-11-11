package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Log(c *fiber.Ctx) error {

	fmt.Println(c)

	return c.Next()
}
