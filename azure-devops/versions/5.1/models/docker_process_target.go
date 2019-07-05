// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DockerProcessTarget Represents the target for the docker build process.
// swagger:model DockerProcessTarget
type DockerProcessTarget struct {
	DesignerProcessTarget
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DockerProcessTarget) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DesignerProcessTarget
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DesignerProcessTarget = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DockerProcessTarget) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.DesignerProcessTarget)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this docker process target
func (m *DockerProcessTarget) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DesignerProcessTarget
	if err := m.DesignerProcessTarget.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DockerProcessTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DockerProcessTarget) UnmarshalBinary(b []byte) error {
	var res DockerProcessTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
