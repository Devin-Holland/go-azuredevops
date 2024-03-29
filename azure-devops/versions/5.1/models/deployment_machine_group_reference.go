// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeploymentMachineGroupReference deployment machine group reference
// swagger:model DeploymentMachineGroupReference
type DeploymentMachineGroupReference struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pool
	Pool *TaskAgentPoolReference `json:"pool,omitempty"`

	// project
	Project *ProjectReference `json:"project,omitempty"`
}

// Validate validates this deployment machine group reference
func (m *DeploymentMachineGroupReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePool(formats); err != nil {
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

func (m *DeploymentMachineGroupReference) validatePool(formats strfmt.Registry) error {

	if swag.IsZero(m.Pool) { // not required
		return nil
	}

	if m.Pool != nil {
		if err := m.Pool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pool")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentMachineGroupReference) validateProject(formats strfmt.Registry) error {

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
func (m *DeploymentMachineGroupReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentMachineGroupReference) UnmarshalBinary(b []byte) error {
	var res DeploymentMachineGroupReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
