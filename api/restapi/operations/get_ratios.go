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

// GetRatiosHandlerFunc turns a function with the right signature into a get ratios handler
type GetRatiosHandlerFunc func(GetRatiosParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRatiosHandlerFunc) Handle(params GetRatiosParams) middleware.Responder {
	return fn(params)
}

// GetRatiosHandler interface for that can handle valid get ratios params
type GetRatiosHandler interface {
	Handle(GetRatiosParams) middleware.Responder
}

// NewGetRatios creates a new http.Handler for the get ratios operation
func NewGetRatios(ctx *middleware.Context, handler GetRatiosHandler) *GetRatios {
	return &GetRatios{Context: ctx, Handler: handler}
}

/*
	GetRatios swagger:route GET /wizard/Ratios getRatios

Get Ratios
*/
type GetRatios struct {
	Context *middleware.Context
	Handler GetRatiosHandler
}

func (o *GetRatios) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRatiosParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetRatiosOKBody get ratios o k body
//
// swagger:model GetRatiosOKBody
type GetRatiosOKBody struct {

	// data
	Data []*models.CarbRatio `json:"Data"`

	// error
	Error interface{} `json:"Error,omitempty"`

	// status
	Status int64 `json:"Status,omitempty"`
}

// Validate validates this get ratios o k body
func (o *GetRatiosOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRatiosOKBody) validateData(formats strfmt.Registry) error {
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
					return ve.ValidateName("getRatiosOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getRatiosOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get ratios o k body based on the context it is used
func (o *GetRatiosOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRatiosOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRatiosOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getRatiosOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRatiosOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRatiosOKBody) UnmarshalBinary(b []byte) error {
	var res GetRatiosOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
