// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestConfiguration Test configuration
// swagger:model TestConfiguration
type TestConfiguration struct {
	TestConfigurationCreateUpdateParameters

	// Id of the configuration
	ID int32 `json:"id,omitempty"`

	// Id of the test configuration variable
	Project *TeamProjectReference `json:"project,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TestConfiguration) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TestConfigurationCreateUpdateParameters
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TestConfigurationCreateUpdateParameters = aO0

	// now for regular properties
	var propsTestConfiguration struct {
		ID int32 `json:"id,omitempty"`

		Project *TeamProjectReference `json:"project,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsTestConfiguration); err != nil {
		return err
	}
	m.ID = propsTestConfiguration.ID

	m.Project = propsTestConfiguration.Project

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TestConfiguration) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TestConfigurationCreateUpdateParameters)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTestConfiguration struct {
		ID int32 `json:"id,omitempty"`

		Project *TeamProjectReference `json:"project,omitempty"`
	}
	propsTestConfiguration.ID = m.ID

	propsTestConfiguration.Project = m.Project

	jsonDataPropsTestConfiguration, errTestConfiguration := swag.WriteJSON(propsTestConfiguration)
	if errTestConfiguration != nil {
		return nil, errTestConfiguration
	}
	_parts = append(_parts, jsonDataPropsTestConfiguration)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this test configuration
func (m *TestConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TestConfigurationCreateUpdateParameters
	if err := m.TestConfigurationCreateUpdateParameters.Validate(formats); err != nil {
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

func (m *TestConfiguration) validateProject(formats strfmt.Registry) error {

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
func (m *TestConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestConfiguration) UnmarshalBinary(b []byte) error {
	var res TestConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
