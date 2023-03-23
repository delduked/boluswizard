// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetTargetsParams creates a new GetTargetsParams object
//
// There are no default values defined in the spec.
func NewGetTargetsParams() GetTargetsParams {

	return GetTargetsParams{}
}

// GetTargetsParams contains all the bound params for the get targets operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTargets
type GetTargetsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	AuthToken string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTargetsParams() beforehand.
func (o *GetTargetsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthToken(r.Header[http.CanonicalHeaderKey("authToken")], true, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthToken binds and validates parameter AuthToken from header.
func (o *GetTargetsParams) bindAuthToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("authToken", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("authToken", "header", raw); err != nil {
		return err
	}
	o.AuthToken = raw

	return nil
}
