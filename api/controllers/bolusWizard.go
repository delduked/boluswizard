package controllers

import (
	"api/services"
	"api/types"

	"github.com/gofiber/fiber/v2"
)

// Calculates a carb correction, BG correction and active insulin reduction based on the current intake of carbs, the current BG value and the current amount of active insulin
func Bolus(c *fiber.Ctx) error {
	uid := c.Locals("Uid").(string)

	body := new(types.CorrectionRequest)
	if err := c.QueryParser(body); err != nil {
		services.ErrorLogger <- err
		c.Status(fiber.StatusBadRequest)
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	correction, err := services.BolusWizard(body, uid)
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
		Data:   correction,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)

}
