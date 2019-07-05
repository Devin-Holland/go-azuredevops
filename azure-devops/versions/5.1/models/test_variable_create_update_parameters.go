// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestVariableCreateUpdateParameters Test Variable Create or Update Parameters
// swagger:model TestVariableCreateUpdateParameters
type TestVariableCreateUpdateParameters struct {

	// Description of the test variable
	Description string `json:"description,omitempty"`

	// Name of the test variable
	Name string `json:"name,omitempty"`

	// List of allowed values
	Values []string `json:"values"`
}

// Validate validates this test variable create update parameters
func (m *TestVariableCreateUpdateParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestVariableCreateUpdateParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestVariableCreateUpdateParameters) UnmarshalBinary(b []byte) error {
	var res TestVariableCreateUpdateParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
