package models

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	// Render index template
	return c.Render("home", fiber.Map{
		"Username": c.Locals("Username"),
	})
}
