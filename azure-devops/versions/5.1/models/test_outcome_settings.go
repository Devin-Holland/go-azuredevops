// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestOutcomeSettings Test outcome settings
// swagger:model TestOutcomeSettings
type TestOutcomeSettings struct {

	// Value to configure how test outcomes for the same tests across suites are shown
	SyncOutcomeAcrossSuites bool `json:"syncOutcomeAcrossSuites,omitempty"`
}

// Validate validates this test outcome settings
func (m *TestOutcomeSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestOutcomeSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestOutcomeSettings) UnmarshalBinary(b []byte) error {
	var res TestOutcomeSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
