// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateCorrectionsHandlerFunc turns a function with the right signature into a create corrections handler
type CreateCorrectionsHandlerFunc func(CreateCorrectionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCorrectionsHandlerFunc) Handle(params CreateCorrectionsParams) middleware.Responder {
	return fn(params)
}

// CreateCorrectionsHandler interface for that can handle valid create corrections params
type CreateCorrectionsHandler interface {
	Handle(CreateCorrectionsParams) middleware.Responder
}

// NewCreateCorrections creates a new http.Handler for the create corrections operation
func NewCreateCorrections(ctx *middleware.Context, handler CreateCorrectionsHandler) *CreateCorrections {
	return &CreateCorrections{Context: ctx, Handler: handler}
}

/*
	CreateCorrections swagger:route POST /corrections createCorrections

Create Corrections
*/
type CreateCorrections struct {
	Context *middleware.Context
	Handler CreateCorrectionsHandler
}

func (o *CreateCorrections) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateCorrectionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}