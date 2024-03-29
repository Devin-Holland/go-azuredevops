// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FlakySettings flaky settings
// swagger:model FlakySettings
type FlakySettings struct {

	// FlakyDetection defines types of detection.
	FlakyDetection *FlakyDetection `json:"flakyDetection,omitempty"`

	// FlakyInSummaryReport defines flaky data should show in summary report or not.
	FlakyInSummaryReport bool `json:"flakyInSummaryReport,omitempty"`

	// ManualMarkUnmarkFlaky defines manual marking unmarking of flaky testcase.
	ManualMarkUnmarkFlaky bool `json:"manualMarkUnmarkFlaky,omitempty"`
}

// Validate validates this flaky settings
func (m *FlakySettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlakyDetection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlakySettings) validateFlakyDetection(formats strfmt.Registry) error {

	if swag.IsZero(m.FlakyDetection) { // not required
		return nil
	}

	if m.FlakyDetection != nil {
		if err := m.FlakyDetection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flakyDetection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlakySettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlakySettings) UnmarshalBinary(b []byte) error {
	var res FlakySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
