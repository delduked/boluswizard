package controllers

import (
	"boluswizard/restapi/operations"
	"boluswizard/restapi/services"

	"github.com/go-openapi/runtime/middleware"
)

func Healthcheck(params operations.HealthCheckParams) middleware.Responder {
	// do some logic in the order you want
	// for example do a check of the jwt token in the header
	return middleware.ResponderFunc(services.Healthcheck)
}
