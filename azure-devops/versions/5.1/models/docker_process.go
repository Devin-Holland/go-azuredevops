// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DockerProcess docker process
// swagger:model DockerProcess
type DockerProcess struct {
	BuildProcess

	// target
	Target *DockerProcessTarget `json:"target,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DockerProcess) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuildProcess
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuildProcess = aO0

	// now for regular properties
	var propsDockerProcess struct {
		Target *DockerProcessTarget `json:"target,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsDockerProcess); err != nil {
		return err
	}
	m.Target = propsDockerProcess.Target

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DockerProcess) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BuildProcess)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsDockerProcess struct {
		Target *DockerProcessTarget `json:"target,omitempty"`
	}
	propsDockerProcess.Target = m.Target

	jsonDataPropsDockerProcess, errDockerProcess := swag.WriteJSON(propsDockerProcess)
	if errDockerProcess != nil {
		return nil, errDockerProcess
	}
	_parts = append(_parts, jsonDataPropsDockerProcess)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this docker process
func (m *DockerProcess) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuildProcess
	if err := m.BuildProcess.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DockerProcess) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DockerProcess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DockerProcess) UnmarshalBinary(b []byte) error {
	var res DockerProcess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}