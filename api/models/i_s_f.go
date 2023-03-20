// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ISF i s f
//
// swagger:model ISF
type ISF struct {

	// end
	End string `json:"End,omitempty"`

	// key
	Key string `json:"Key,omitempty"`

	// sensitivity
	Sensitivity float64 `json:"Sensitivity,omitempty"`

	// start
	Start string `json:"Start,omitempty"`
}

// Validate validates this i s f
func (m *ISF) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this i s f based on context it is used
func (m *ISF) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ISF) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ISF) UnmarshalBinary(b []byte) error {
	var res ISF
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
