package models

import (
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	// Render index template
	return c.Render("signup", nil)
}
