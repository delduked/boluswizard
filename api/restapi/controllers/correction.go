package controllers

import (
	"boluswizard/models"
	"boluswizard/restapi/operations"
	"boluswizard/restapi/services"
	"boluswizard/restapi/tools"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

func Newcorrection(params operations.NewCorrectionParams) middleware.Responder {
	// do some logic in the order you want
	// for example do a check of the jwt token in the header

	// decrypt the token and get the uid
	// if the service that determines the token to be invalid by throwing an error
	// return the middleware.ResponderFunc(services.Error) or whatever function you make

	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	correction, err := services.BolusWizard(params.Bg, params.Carbs, uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	response := &operations.NewCorrectionOKBody{
		Status: 200,
		Error:  nil,
		Data:   &correction,
	}

	return operations.NewNewCorrectionOK().WithPayload(response)
}

// Get All Corrections
func Corrections(params operations.GetCorrectionsParams) middleware.Responder {

	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	corrections, err := services.Corrections(uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	response := &operations.GetCorrectionsOKBody{
		Status: 200,
		Error:  nil,
		Data:   corrections,
	}

	return operations.NewGetCorrectionsOK().WithPayload(response)
}

// Get a range of corrections specified by the user
func CorrectionRange(params operations.CorrectionRangeParams) middleware.Responder {

	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	body := models.CorrectionRange{
		Start: params.Start,
		End:   params.End,
	}

	CorrectionRange, err := services.CorrectionRange(body, uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	var data []*models.Correction
	for i, j := range CorrectionRange {
		data[i] = &j
	}

	response := &operations.CorrectionRangeOKBody{
		Status: 200,
		Error:  nil,
		Data:   data,
	}

	return operations.NewCorrectionRangeOK().WithPayload(response)

}

// Save a single or multiple corrections
func SaveCorrections(params operations.CreateCorrectionsParams) middleware.Responder {

	cookies, err := tools.GetCookies(params.HTTPRequest, "wizardToken")
	if err != nil {
		return services.NewError(http.StatusUnauthorized, err)
	}

	uid, err := services.VerifyCredentialsWithToken(cookies["wizardToken"])
	if err != nil {
		return tools.DeleteCookie("wizardToken", err)
	}

	var data []models.Correction
	for i, j := range params.Corrections {
		data[i] = *j
	}

	// save the corrections in redis
	err = services.SaveCorrections(data, uid)
	if err != nil {
		return middleware.ResponderFunc(services.Error)
	}

	// construct the response
	response := &operations.CorrectionRangeOKBody{
		Status: 200,
		Error:  nil,
	}

	return operations.NewCorrectionRangeOK().WithPayload(response)

}
