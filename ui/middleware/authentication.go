package middleware

import (
	"fmt"
	"ui/controllers"

	"github.com/gofiber/fiber/v2"
)

func ValidateToken(c *fiber.Ctx) error {

	bearer := c.Cookies("authToken")
	uid, username, err := controllers.VerifyToken(bearer, c)
	if err != nil {
		c.Status(fiber.StatusForbidden)
		c.Redirect("/error/403") // redirect to login page
	}

	c.Locals("Uid", uid)
	c.Locals("Username", username)

	return c.Next()
}

func VerifyUser(c *fiber.Ctx) error {
	fmt.Println(c.Params("user"))
	fmt.Println(c.Locals("Username"))

	if c.Params("user") != c.Locals("Username") {
		c.ClearCookie("authToken")
		c.Status(fiber.StatusForbidden)
		c.Redirect("/error/401") // redirect to login page
	}

	return c.Next()
}
