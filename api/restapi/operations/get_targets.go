// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"boluswizard/models"
)

// GetTargetsHandlerFunc turns a function with the right signature into a get targets handler
type GetTargetsHandlerFunc func(GetTargetsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTargetsHandlerFunc) Handle(params GetTargetsParams) middleware.Responder {
	return fn(params)
}

// GetTargetsHandler interface for that can handle valid get targets params
type GetTargetsHandler interface {
	Handle(GetTargetsParams) middleware.Responder
}

// NewGetTargets creates a new http.Handler for the get targets operation
func NewGetTargets(ctx *middleware.Context, handler GetTargetsHandler) *GetTargets {
	return &GetTargets{Context: ctx, Handler: handler}
}

/*
	GetTargets swagger:route GET /Targets getTargets

Get Targets
*/
type GetTargets struct {
	Context *middleware.Context
	Handler GetTargetsHandler
}

func (o *GetTargets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTargetsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetTargetsOKBody get targets o k body
//
// swagger:model GetTargetsOKBody
type GetTargetsOKBody struct {

	// data
	Data []*models.Target `json:"Data"`

	// error
	Error interface{} `json:"Error,omitempty"`

	// status
	Status int64 `json:"Status,omitempty"`
}

// Validate validates this get targets o k body
func (o *GetTargetsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTargetsOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getTargetsOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getTargetsOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get targets o k body based on the context it is used
func (o *GetTargetsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTargetsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getTargetsOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getTargetsOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTargetsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTargetsOKBody) UnmarshalBinary(b []byte) error {
	var res GetTargetsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
