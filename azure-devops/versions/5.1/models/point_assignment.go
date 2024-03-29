// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PointAssignment Assignments for the Test Point
// swagger:model PointAssignment
type PointAssignment struct {
	Configuration

	// Name of the Configuration Assigned to the Test Point
	ConfigurationName string `json:"configurationName,omitempty"`

	// Id of the Test Point
	ID int32 `json:"id,omitempty"`

	// Tester Assigned to the Test Point
	Tester *IdentityRef `json:"tester,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PointAssignment) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Configuration
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Configuration = aO0

	// now for regular properties
	var propsPointAssignment struct {
		ConfigurationName string `json:"configurationName,omitempty"`

		ID int32 `json:"id,omitempty"`

		Tester *IdentityRef `json:"tester,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsPointAssignment); err != nil {
		return err
	}
	m.ConfigurationName = propsPointAssignment.ConfigurationName

	m.ID = propsPointAssignment.ID

	m.Tester = propsPointAssignment.Tester

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PointAssignment) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Configuration)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsPointAssignment struct {
		ConfigurationName string `json:"configurationName,omitempty"`

		ID int32 `json:"id,omitempty"`

		Tester *IdentityRef `json:"tester,omitempty"`
	}
	propsPointAssignment.ConfigurationName = m.ConfigurationName

	propsPointAssignment.ID = m.ID

	propsPointAssignment.Tester = m.Tester

	jsonDataPropsPointAssignment, errPointAssignment := swag.WriteJSON(propsPointAssignment)
	if errPointAssignment != nil {
		return nil, errPointAssignment
	}
	_parts = append(_parts, jsonDataPropsPointAssignment)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this point assignment
func (m *PointAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Configuration
	if err := m.Configuration.Validate(formats); err != nil {
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

func (m *PointAssignment) validateTester(formats strfmt.Registry) error {

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
func (m *PointAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PointAssignment) UnmarshalBinary(b []byte) error {
	var res PointAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
