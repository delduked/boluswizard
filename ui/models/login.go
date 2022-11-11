package models

import (
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// Render index template
	return c.Render("login", nil)
}
