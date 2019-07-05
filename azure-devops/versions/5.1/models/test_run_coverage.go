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

// TestRunCoverage Test Run Code Coverage Details
// swagger:model TestRunCoverage
type TestRunCoverage struct {

	// Last Error
	LastError string `json:"lastError,omitempty"`

	// List of Modules Coverage
	Modules []*ModuleCoverage `json:"modules"`

	// State
	State string `json:"state,omitempty"`

	// Reference of test Run.
	TestRun *ShallowReference `json:"testRun,omitempty"`
}

// Validate validates this test run coverage
func (m *TestRunCoverage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestRunCoverage) validateModules(formats strfmt.Registry) error {

	if swag.IsZero(m.Modules) { // not required
		return nil
	}

	for i := 0; i < len(m.Modules); i++ {
		if swag.IsZero(m.Modules[i]) { // not required
			continue
		}

		if m.Modules[i] != nil {
			if err := m.Modules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestRunCoverage) validateTestRun(formats strfmt.Registry) error {

	if swag.IsZero(m.TestRun) { // not required
		return nil
	}

	if m.TestRun != nil {
		if err := m.TestRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testRun")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestRunCoverage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestRunCoverage) UnmarshalBinary(b []byte) error {
	var res TestRunCoverage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
