package controllers

import (
	"api/services"
	"api/types"

	"github.com/gofiber/fiber/v2"
)

// Get All Corrections
func Corrections(c *fiber.Ctx) error {
	uid := c.Locals("Uid").(string)

	corrections, err := services.Corrections(uid)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return services.Response(res, c)
	}

	res := types.Response{
		Status: fiber.StatusOK,
		Error:  err,
		Data:   corrections,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)
}

// Get a Single Correction
func Correction(c *fiber.Ctx) error {
	uid := c.Locals("Uid").(string)
	key := c.Params("Key")

	correction, err := services.Correction(uid, key)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	res := types.Response{
		Status: fiber.StatusOK,
		Error:  err,
		Data:   correction,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)
}

// Get a range of corrections specified by the user
func CorrectionRange(c *fiber.Ctx) error {
	uid := c.Locals("Uid").(string)

	body := new(types.CorrectionRange)
	if err := c.QueryParser(body); err != nil {
		services.ErrorLogger <- err
		c.Status(fiber.StatusBadRequest)
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	corrections, err := services.CorrectionRange(*body, uid)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return services.Response(res, c)
	}

	res := types.Response{
		Status: fiber.StatusOK,
		Error:  err,
		Data:   corrections,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)
}

// Save a single or multiple corrections
func SaveCorrections(c *fiber.Ctx) error {
	// only accept POST requests
	c.Accepts("application/json")

	uid := c.Locals("Uid").(string) // place holder for user account when jwt is finsihed

	// check if the request body matches the requires body type
	body := new([]types.Correction)
	if err := c.BodyParser(body); err != nil {
		services.ErrorLogger <- err
		c.Status(fiber.StatusBadRequest)
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	// save the corrections in redis
	err := services.SaveCorrections(*body, uid)
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
