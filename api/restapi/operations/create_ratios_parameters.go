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

// NewCreateRatiosParams creates a new CreateRatiosParams object
//
// There are no default values defined in the spec.
func NewCreateRatiosParams() CreateRatiosParams {

	return CreateRatiosParams{}
}

// CreateRatiosParams contains all the bound params for the create ratios operation
// typically these are obtained from a http.Request
//
// swagger:parameters createRatios
type CreateRatiosParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	CarbRatios []*models.CarbRatio
	/*
	  Required: true
	  In: header
	*/
	WizardToken string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateRatiosParams() beforehand.
func (o *CreateRatiosParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.CarbRatio
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("carbRatios", "body", ""))
			} else {
				res = append(res, errors.NewParseError("carbRatios", "body", "", err))
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
				o.CarbRatios = body
			}
		}
	} else {
		res = append(res, errors.Required("carbRatios", "body", ""))
	}

	if err := o.bindWizardToken(r.Header[http.CanonicalHeaderKey("wizardToken")], true, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindWizardToken binds and validates parameter WizardToken from header.
func (o *CreateRatiosParams) bindWizardToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("wizardToken", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("wizardToken", "header", raw); err != nil {
		return err
	}
	o.WizardToken = raw

	return nil
}
