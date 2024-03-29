// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestResultsSettings test results settings
// swagger:model TestResultsSettings
type TestResultsSettings struct {

	// IsRequired and EmitDefaultValue are passed as false as if users doesn't pass anything, should not come for serialisation and deserialisation.
	FlakySettings *FlakySettings `json:"flakySettings,omitempty"`
}

// Validate validates this test results settings
func (m *TestResultsSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlakySettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestResultsSettings) validateFlakySettings(formats strfmt.Registry) error {

	if swag.IsZero(m.FlakySettings) { // not required
		return nil
	}

	if m.FlakySettings != nil {
		if err := m.FlakySettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flakySettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestResultsSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultsSettings) UnmarshalBinary(b []byte) error {
	var res TestResultsSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
