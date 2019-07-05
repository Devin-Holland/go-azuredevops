// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestResultsUpdateSettings test results update settings
// swagger:model TestResultsUpdateSettings
type TestResultsUpdateSettings struct {

	// FlakySettings defines Flaky Setttings Data.
	FlakySettings *FlakySettings `json:"flakySettings,omitempty"`
}

// Validate validates this test results update settings
func (m *TestResultsUpdateSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlakySettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestResultsUpdateSettings) validateFlakySettings(formats strfmt.Registry) error {

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
func (m *TestResultsUpdateSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultsUpdateSettings) UnmarshalBinary(b []byte) error {
	var res TestResultsUpdateSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
