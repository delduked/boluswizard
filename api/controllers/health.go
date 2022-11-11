package controllers

import (
	"api/services"
	"api/types"

	"github.com/gofiber/fiber/v2"
)

// Get status of API
func Health(c *fiber.Ctx) error {

	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)

}
