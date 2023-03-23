package controllers

import (
	"boluswizard/models"
	"boluswizard/restapi/operations"
	"boluswizard/restapi/services"

	"github.com/go-openapi/runtime/middleware"
)

// Save a single or multiple ISFs with an array of ISF values
func SaveISFs(params operations.CreateISFsParams) middleware.Responder {
	// only accept POST requests
	uid, err := services.VerifyCredentialsWithToken(params.AuthToken)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	var data []models.ISF
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

	uid, err := services.VerifyCredentialsWithToken(params.AuthToken)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	isf, err := services.ISFs(uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	var data []*models.ISF
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
