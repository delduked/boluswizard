// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Response response
//
// swagger:model Response
type Response struct {

	// error
	Error interface{} `json:"Error,omitempty"`

	// status
	Status int64 `json:"Status,omitempty"`
}

// Validate validates this response
func (m *Response) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this response based on context it is used
func (m *Response) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Response) UnmarshalBinary(b []byte) error {
	var res Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
