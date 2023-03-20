// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateISFsHandlerFunc turns a function with the right signature into a create i s fs handler
type CreateISFsHandlerFunc func(CreateISFsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateISFsHandlerFunc) Handle(params CreateISFsParams) middleware.Responder {
	return fn(params)
}

// CreateISFsHandler interface for that can handle valid create i s fs params
type CreateISFsHandler interface {
	Handle(CreateISFsParams) middleware.Responder
}

// NewCreateISFs creates a new http.Handler for the create i s fs operation
func NewCreateISFs(ctx *middleware.Context, handler CreateISFsHandler) *CreateISFs {
	return &CreateISFs{Context: ctx, Handler: handler}
}

/*
	CreateISFs swagger:route POST /ISF createISFs

Create ISF
*/
type CreateISFs struct {
	Context *middleware.Context
	Handler CreateISFsHandler
}

func (o *CreateISFs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateISFsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateISFsOKBody create i s fs o k body
//
// swagger:model CreateISFsOKBody
type CreateISFsOKBody struct {

	// data
	Data interface{} `json:"Data,omitempty"`

	// error
	Error interface{} `json:"Error,omitempty"`

	// status
	Status int64 `json:"Status,omitempty"`
}

// Validate validates this create i s fs o k body
func (o *CreateISFsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create i s fs o k body based on context it is used
func (o *CreateISFsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateISFsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateISFsOKBody) UnmarshalBinary(b []byte) error {
	var res CreateISFsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
