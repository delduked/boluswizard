package controllers

import (
	"boluswizard/models"
	"boluswizard/restapi/operations"
	"boluswizard/restapi/services"

	"github.com/go-openapi/runtime/middleware"
)

// Save a single or multiple BG targets with an array of values
func SaveTargets(params operations.CreateTargetsParams) middleware.Responder {

	uid, err := services.VerifyCredentialsWithToken(params.AuthToken)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	var body []models.Target
	for i, j := range params.TargetRatios {
		body[i] = *j
	}

	// save the targets in redis
	err = services.SaveTargets(body, uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	// construct the response
	response := &operations.CreateTargetsOKBody{
		Status: 200,
		Error:  nil,
	}
	return operations.NewCreateTargetsOK().WithPayload(response)
}

// Get a single or multiple BG Targets with an array of Target values
func Targets(params operations.GetTargetsParams) middleware.Responder {

	uid, err := services.VerifyCredentialsWithToken(params.AuthToken)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	targets, err := services.Targets(uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	var data []*models.Target
	for i, j := range targets {
		data[i] = &j
	}

	response := &operations.GetTargetsOKBody{
		Status: 200,
		Data:   data,
	}
	return operations.NewGetTargetsOK().WithPayload(response)
}
