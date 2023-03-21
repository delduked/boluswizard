// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Health health
//
// swagger:model Health
type Health struct {

	// error
	Error interface{} `json:"Error,omitempty"`

	// status
	Status int64 `json:"Status,omitempty"`

	// timestamp
	Timestamp interface{} `json:"Timestamp,omitempty"`
}

// Validate validates this health
func (m *Health) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this health based on context it is used
func (m *Health) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Health) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Health) UnmarshalBinary(b []byte) error {
	var res Health
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
