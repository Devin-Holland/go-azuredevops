// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestCaseReference Test Case Reference
// swagger:model TestCaseReference
type TestCaseReference struct {

	// Identity to whom the test case is assigned
	AssignedTo *IdentityRef `json:"assignedTo,omitempty"`

	// Test Case Id
	ID int32 `json:"id,omitempty"`

	// Test Case Name
	Name string `json:"name,omitempty"`

	// State of the test case work item
	State string `json:"state,omitempty"`
}

// Validate validates this test case reference
func (m *TestCaseReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCaseReference) validateAssignedTo(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignedTo) { // not required
		return nil
	}

	if m.AssignedTo != nil {
		if err := m.AssignedTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignedTo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCaseReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCaseReference) UnmarshalBinary(b []byte) error {
	var res TestCaseReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
