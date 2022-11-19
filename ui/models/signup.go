package models

import (
	"ui/config"

	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	// Render index template
	return c.Render("signup", fiber.Map{
		"API_SERVICE": config.API_SERVICE,
	})
}
