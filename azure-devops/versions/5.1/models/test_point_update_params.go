// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestPointUpdateParams Test Point Update Parameters
// swagger:model TestPointUpdateParams
type TestPointUpdateParams struct {

	// Id of Test Point to be updated
	ID int32 `json:"id,omitempty"`

	// Reset the Test Point to Active
	IsActive bool `json:"isActive,omitempty"`

	// Results of the test point
	Results *Results `json:"results,omitempty"`

	// Tester of the Test Point
	Tester *IdentityRef `json:"tester,omitempty"`
}

// Validate validates this test point update params
func (m *TestPointUpdateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTester(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestPointUpdateParams) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	if m.Results != nil {
		if err := m.Results.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

func (m *TestPointUpdateParams) validateTester(formats strfmt.Registry) error {

	if swag.IsZero(m.Tester) { // not required
		return nil
	}

	if m.Tester != nil {
		if err := m.Tester.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tester")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestPointUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestPointUpdateParams) UnmarshalBinary(b []byte) error {
	var res TestPointUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
