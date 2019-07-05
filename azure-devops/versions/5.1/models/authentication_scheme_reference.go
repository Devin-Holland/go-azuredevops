// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AuthenticationSchemeReference authentication scheme reference
// swagger:model AuthenticationSchemeReference
type AuthenticationSchemeReference struct {

	// inputs
	Inputs map[string]string `json:"inputs,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this authentication scheme reference
func (m *AuthenticationSchemeReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationSchemeReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationSchemeReference) UnmarshalBinary(b []byte) error {
	var res AuthenticationSchemeReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
