package controllers

import (
	"boluswizard/restapi/operations"
	"boluswizard/restapi/services"
	"boluswizard/restapi/tools"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// Get a Single duration value for a specific user
func Duration(params operations.GetDurationParams) middleware.Responder {

	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	data, err := services.GetDuration(uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	response := &operations.GetDurationOKBody{
		Status: 200,
		Error:  nil,
		Data:   data.Duration,
	}

	return operations.NewGetDurationOK().WithPayload(response)
}

// Save a single insulin duration value for a specific user
func SaveDuration(params operations.CreateDurationParams) middleware.Responder {

	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	// check if the duration string is in the required format ex: 02h00m or 13h15m
	err = services.ValidateUserInput(params.Duration.Duration)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	// save the duration in redis
	err = services.SaveDuration(uid, params.Duration.Duration)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}
	// construct the response
	response := &operations.GetDurationOKBody{
		Status: 200,
		Error:  nil,
	}

	return operations.NewGetDurationOK().WithPayload(response)

}
