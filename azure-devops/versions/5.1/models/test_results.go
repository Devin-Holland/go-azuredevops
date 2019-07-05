// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestResults test results
// swagger:model TestResults
type TestResults struct {

	// cloud load test solution Url
	CloudLoadTestSolutionURL string `json:"cloudLoadTestSolutionUrl,omitempty"`

	// counter groups
	CounterGroups []*CounterGroup `json:"counterGroups"`

	// diagnostics
	Diagnostics *Diagnostics `json:"diagnostics,omitempty"`

	// results Url
	ResultsURL string `json:"resultsUrl,omitempty"`
}

// Validate validates this test results
func (m *TestResults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounterGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiagnostics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestResults) validateCounterGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.CounterGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.CounterGroups); i++ {
		if swag.IsZero(m.CounterGroups[i]) { // not required
			continue
		}

		if m.CounterGroups[i] != nil {
			if err := m.CounterGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counterGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestResults) validateDiagnostics(formats strfmt.Registry) error {

	if swag.IsZero(m.Diagnostics) { // not required
		return nil
	}

	if m.Diagnostics != nil {
		if err := m.Diagnostics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diagnostics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestResults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResults) UnmarshalBinary(b []byte) error {
	var res TestResults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}