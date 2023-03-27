package controllers

import (
	"boluswizard/models"
	"boluswizard/restapi/operations"
	"boluswizard/restapi/services"
	"boluswizard/restapi/tools"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// Save a single or multiple ISFs with an array of ISF values
func SaveISFs(params operations.CreateISFsParams) middleware.Responder {
	// only accept POST requests
	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	//var data []models.ISF
	data := make([]models.ISF, len(params.ISFs))
	for i, j := range params.ISFs {
		data[i] = *j
	}

	// save the targets in redis
	err = services.SaveISFs(data, uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	// construct the response
	response := &operations.CreateISFsOKBody{
		Status: 200,
		Error:  nil,
	}

	return operations.NewCreateISFsOK().WithPayload(response)
}

// Get all ISFs saved for the specific logged in user
func ISFs(params operations.GetISFsParams) middleware.Responder {

	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	isf, err := services.ISFs(uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	data := make([]*models.ISF, len(isf))
	for i, j := range isf {
		data[i] = &j
	}

	// construct the response
	response := &operations.GetISFsOKBody{
		Status: 200,
		Error:  err,
		Data:   data,
	}
	return operations.NewGetISFsOK().WithPayload(response)
}
