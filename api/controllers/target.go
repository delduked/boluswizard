package controllers

import (
	"api/services"
	"api/types"

	"github.com/gofiber/fiber/v2"
)

// Save a single or multiple BG targets with an array of values
func SaveTargets(c *fiber.Ctx) error {
	// only accept POST requests
	c.Accepts("application/json")

	uid := c.Locals("Uid").(string) // place holder for user account when jwt is finsihed

	// check if the request body matches the requires body type
	body := new([]types.Target)
	if err := c.BodyParser(body); err != nil {
		services.ErrorLogger <- err
		c.Status(fiber.StatusBadRequest)
		res := types.Response[any]{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	// save the targets in redis
	err := services.SaveTargets(*body, uid)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		res := types.Response[any]{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return services.Response(res, c)
	}

	// construct the response
	res := types.Response[any]{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)
}

// Get a single or multiple BG Targets with an array of Target values
func Targets(c *fiber.Ctx) error {

	uid := c.Locals("Uid").(string)

	targets, err := services.Targets(uid)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		res := types.Response[any]{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return services.Response(res, c)
	}

	res := types.Response[any]{
		Status: fiber.StatusOK,
		Error:  err,
		Data:   targets,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)
}
