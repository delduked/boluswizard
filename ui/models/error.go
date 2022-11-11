package models

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorRedirect(c *fiber.Ctx) error {
	return c.Redirect("/error/404")
}
func ErrorPage(c *fiber.Ctx) error {
	// Render index template
	// potentially populate the nil options with a json object to add the specific error code in the future

	status := c.Params("error")

	switch status {
	case "404":
		return c.Render("error", fiber.Map{
			"Status":  status,
			"Message": "The page you are looking for does not exist",
		})
	case "403":
		return c.Render("error", fiber.Map{
			"Status":  status,
			"Message": "Looks like you do not have access to the selected resource.",
		})
	case "401":
		return c.Render("error", fiber.Map{
			"Status":  status,
			"Message": "Looks like you do not have access to the selected resource.",
		})
	case "400":
		return c.Render("error", fiber.Map{
			"Status":  status,
			"Message": "The resource request was made improperly.",
		})
	default:
		return c.Render("error", fiber.Map{
			"Status":  "503",
			"Message": "Oops... looks like there was an error!",
		})
	}
}
