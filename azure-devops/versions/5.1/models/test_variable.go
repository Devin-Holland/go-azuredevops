// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestVariable Test Variable
// swagger:model TestVariable
type TestVariable struct {
	TestVariableCreateUpdateParameters

	// Id of the test variable
	ID int32 `json:"id,omitempty"`

	// Id of the test variable
	Project *TeamProjectReference `json:"project,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TestVariable) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TestVariableCreateUpdateParameters
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TestVariableCreateUpdateParameters = aO0

	// now for regular properties
	var propsTestVariable struct {
		ID int32 `json:"id,omitempty"`

		Project *TeamProjectReference `json:"project,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsTestVariable); err != nil {
		return err
	}
	m.ID = propsTestVariable.ID

	m.Project = propsTestVariable.Project

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TestVariable) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TestVariableCreateUpdateParameters)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTestVariable struct {
		ID int32 `json:"id,omitempty"`

		Project *TeamProjectReference `json:"project,omitempty"`
	}
	propsTestVariable.ID = m.ID

	propsTestVariable.Project = m.Project

	jsonDataPropsTestVariable, errTestVariable := swag.WriteJSON(propsTestVariable)
	if errTestVariable != nil {
		return nil, errTestVariable
	}
	_parts = append(_parts, jsonDataPropsTestVariable)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this test variable
func (m *TestVariable) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TestVariableCreateUpdateParameters
	if err := m.TestVariableCreateUpdateParameters.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestVariable) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestVariable) UnmarshalBinary(b []byte) error {
	var res TestVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
