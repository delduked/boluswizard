package controllers

import (
	"api/services"
	"api/types"

	"github.com/gofiber/fiber/v2"
)

// Get a Single duration value for a specific user
func Duration(c *fiber.Ctx) error {
	uid := c.Locals("Uid").(string)

	duration, err := services.GetDuration(uid)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	res := types.Response{
		Status: fiber.StatusOK,
		Error:  err,
		Data:   duration,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)
}

// Save a single insulin duration value for a specific user
func SaveDuration(c *fiber.Ctx) error {
	var err error
	// only accept POST requests
	c.Accepts("application/json")

	uid := c.Locals("Uid").(string)

	// check if the request body matches the requires body type
	body := new(types.ActiveInsulinDuration)
	if err = c.BodyParser(body); err != nil {
		services.ErrorLogger <- err
		c.Status(fiber.StatusBadRequest)
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	// check if the duration string is in the required format ex: 02h00m or 13h15m
	err = services.ValidateUserInput(body.Duration)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	// save the duration in redis
	err = services.SaveDuration(uid, body)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return services.Response(res, c)
	}

	// construct the response
	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)

}
