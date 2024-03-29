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

// TestFailureDetails test failure details
// swagger:model TestFailureDetails
type TestFailureDetails struct {

	// count
	Count int32 `json:"count,omitempty"`

	// test results
	TestResults []*TestCaseResultIdentifier `json:"testResults"`
}

// Validate validates this test failure details
func (m *TestFailureDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTestResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestFailureDetails) validateTestResults(formats strfmt.Registry) error {

	if swag.IsZero(m.TestResults) { // not required
		return nil
	}

	for i := 0; i < len(m.TestResults); i++ {
		if swag.IsZero(m.TestResults[i]) { // not required
			continue
		}

		if m.TestResults[i] != nil {
			if err := m.TestResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestFailureDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestFailureDetails) UnmarshalBinary(b []byte) error {
	var res TestFailureDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
