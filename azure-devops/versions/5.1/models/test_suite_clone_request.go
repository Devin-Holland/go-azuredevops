// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestSuiteCloneRequest Test suite clone request
// swagger:model TestSuiteCloneRequest
type TestSuiteCloneRequest struct {

	// Clone options for cloning the test suite.
	CloneOptions *CloneOptions `json:"cloneOptions,omitempty"`

	// Suite id under which, we have to clone the suite.
	DestinationSuiteID int32 `json:"destinationSuiteId,omitempty"`

	// Destination suite project name.
	DestinationSuiteProjectName string `json:"destinationSuiteProjectName,omitempty"`
}

// Validate validates this test suite clone request
func (m *TestSuiteCloneRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloneOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestSuiteCloneRequest) validateCloneOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.CloneOptions) { // not required
		return nil
	}

	if m.CloneOptions != nil {
		if err := m.CloneOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestSuiteCloneRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestSuiteCloneRequest) UnmarshalBinary(b []byte) error {
	var res TestSuiteCloneRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}