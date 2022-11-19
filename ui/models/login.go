package models

import (
	"ui/config"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// Render index template
	return c.Render("login", fiber.Map{
		"API_SERVICE": config.API_SERVICE,
	})
}
