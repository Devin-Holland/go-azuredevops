// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestConfigurationReference Test Configuration Reference
// swagger:model TestConfigurationReference
type TestConfigurationReference struct {

	// Id of the configuration
	ID int32 `json:"id,omitempty"`

	// Name of the configuration
	Name string `json:"name,omitempty"`
}

// Validate validates this test configuration reference
func (m *TestConfigurationReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestConfigurationReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestConfigurationReference) UnmarshalBinary(b []byte) error {
	var res TestConfigurationReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}