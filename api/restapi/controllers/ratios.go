package controllers

import (
	"boluswizard/models"
	"boluswizard/restapi/operations"
	"boluswizard/restapi/services"
	"boluswizard/restapi/tools"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// Save a single or multiple carb ratios with an array of values
func SaveRatios(params operations.CreateRatiosParams) middleware.Responder {

	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	var body []models.CarbRatio
	for i, j := range params.CarbRatios {
		body[i] = *j
	}

	// save the targets in redis
	err = services.SaveRatios(body, uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	// construct the response
	response := &operations.CreateRatiosOKBody{
		Status: 200,
		Error:  nil,
	}
	return operations.NewCreateRatiosOK().WithPayload(response)
}

// Get all carb ratios for the specific user logged in
func Ratios(params operations.GetRatiosParams) middleware.Responder {
	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	ratios, err := services.Ratios(uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	var data []*models.CarbRatio
	for i, j := range ratios {
		data[i] = &j
	}

	response := &operations.GetRatiosOKBody{
		Status: 200,
		Error:  err,
		Data:   data,
	}
	return operations.NewGetRatiosOK().WithPayload(response)
}
