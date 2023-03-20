// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Correction correction
//
// swagger:model Correction
type Correction struct {

	// bg
	Bg float64 `json:"Bg,omitempty"`

	// bolus
	Bolus float64 `json:"Bolus,omitempty"`

	// carbs
	Carbs float64 `json:"Carbs,omitempty"`

	// key
	Key string `json:"Key,omitempty"`

	// time stamp
	TimeStamp int64 `json:"TimeStamp,omitempty"`
}

// Validate validates this correction
func (m *Correction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this correction based on context it is used
func (m *Correction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Correction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Correction) UnmarshalBinary(b []byte) error {
	var res Correction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
