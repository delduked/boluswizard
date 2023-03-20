// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"boluswizard/models"
)

// NewCreateISFsParams creates a new CreateISFsParams object
//
// There are no default values defined in the spec.
func NewCreateISFsParams() CreateISFsParams {

	return CreateISFsParams{}
}

// CreateISFsParams contains all the bound params for the create i s fs operation
// typically these are obtained from a http.Request
//
// swagger:parameters createISFs
type CreateISFsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	AuthToken string
	/*
	  Required: true
	  In: body
	*/
	TargetRatios []*models.ISF
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateISFsParams() beforehand.
func (o *CreateISFsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthToken(r.Header[http.CanonicalHeaderKey("authToken")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.ISF
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("targetRatios", "body", ""))
			} else {
				res = append(res, errors.NewParseError("targetRatios", "body", "", err))
			}
		} else {

			// validate array of body objects
			for i := range body {
				if body[i] == nil {
					continue
				}
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.TargetRatios = body
			}
		}
	} else {
		res = append(res, errors.Required("targetRatios", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthToken binds and validates parameter AuthToken from header.
func (o *CreateISFsParams) bindAuthToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
