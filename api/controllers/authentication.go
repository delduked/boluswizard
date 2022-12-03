package controllers

import (
	"api/services"
	"api/types"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {

	c.Accepts("application/json")

	body := new(types.Users)
	if err := c.BodyParser(body); err != nil {
		services.ErrorLogger <- err
		c.Status(fiber.StatusBadRequest)
		res := types.Response[any]{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	authToken, err := services.SignUp(body)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		res := types.Response[any]{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "authToken"
	cookie.Value = authToken

	c.Cookie(cookie)

	data := types.SigninData{
		Username: body.Username,
		Token:    authToken,
	}

	res := types.Response[types.SigninData]{
		Status: fiber.StatusOK,
		Error:  nil,
		Data:   data,
	}

	return services.Response(res, c)
}

func VerifyMiddleWare(c *fiber.Ctx) error {
	bearer := c.Cookies("authToken")
	uid, err := services.VerifyCredentialsWithToken(bearer, c)
	if err != nil {
		c.Status(fiber.StatusForbidden)
		res := types.Response[any]{
			Status: fiber.StatusForbidden,
			Error:  err,
		}
		return services.Response(res, c)
	}

	c.Locals("Uid", uid)

	return c.Next()
}
func Signin(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(types.Users)
	if err := c.BodyParser(body); err != nil {
		services.ErrorLogger <- err
		c.Status(fiber.StatusBadRequest)
		res := types.Response[any]{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return services.Response(res, c)
	}

	validate, user, err := services.VerifyUsernameAndPassword(*body)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		res := types.Response[any]{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return services.Response(res, c)
	}

	if !validate {
		c.Status(fiber.StatusForbidden)
		res := types.Response[any]{
			Status: fiber.StatusForbidden,
			Error:  err,
		}
		return services.Response(res, c)
	}

	authToken, err := services.SignIn(user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		res := types.Response[any]{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return services.Response(res, c)
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "authToken"
	cookie.Value = authToken
	cookie.Path = "/"

	c.Cookie(cookie)

	data := types.SigninData{
		Username: body.Username,
		Token:    authToken,
	}

	res := types.Response[any]{
		Status: fiber.StatusOK,
		Error:  err,
		Data:   data,
	}
	c.Status(fiber.StatusOK)
	return services.Response(res, c)
}
