// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Timestamp timestamp
//
// swagger:model Timestamp
type Timestamp struct {

	// timestamp
	Timestamp string `json:"Timestamp,omitempty"`
}

// Validate validates this timestamp
func (m *Timestamp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this timestamp based on context it is used
func (m *Timestamp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Timestamp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Timestamp) UnmarshalBinary(b []byte) error {
	var res Timestamp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
