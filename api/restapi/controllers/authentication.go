package controllers

import (
	"boluswizard/models"
	"boluswizard/restapi/operations"
	"boluswizard/restapi/services"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

func SignUp(params operations.SignUpParams) middleware.Responder {

	body := models.Users{
		Password: params.SignUpData.Password,
		Username: params.SignUpData.Username,
	}

	token, err := services.SignUp(&body)
	if err != nil {
		services.NewError(http.StatusBadRequest, err)
	}

	response := services.Authenticated(body.Username, token)

	return middleware.ResponderFunc(response)

}

func Signin(params operations.SignInParams) middleware.Responder {

	body := models.Users{
		UID:      "",
		Password: params.SignInData.Password,
		Username: params.SignInData.Username,
	}

	validate, user, err := services.VerifyUsernameAndPassword(body)
	//c.Status(fiber.StatusInternalServerError)
	if err != nil {
		services.NewError(http.StatusBadRequest, err)
	}

	if !validate {
		//c.Status(fiber.StatusForbidden)
		if err != nil {
			services.NewError(http.StatusForbidden, err)
		}

	}

	token, err := services.SignIn(user)
	if err != nil {
		services.NewError(http.StatusInternalServerError, err)
	}

	response := services.Authenticated(body.Username, token)

	return middleware.ResponderFunc(response)
}
