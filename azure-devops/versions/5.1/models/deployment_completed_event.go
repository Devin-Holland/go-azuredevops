// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeploymentCompletedEvent deployment completed event
// swagger:model DeploymentCompletedEvent
type DeploymentCompletedEvent struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// data
	Data map[string]interface{} `json:"data,omitempty"`

	// deployment
	Deployment *Deployment `json:"deployment,omitempty"`

	// environment
	Environment *ReleaseEnvironment `json:"environment,omitempty"`

	// project
	Project *ProjectReference `json:"project,omitempty"`
}

// Validate validates this deployment completed event
func (m *DeploymentCompletedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
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

func (m *DeploymentCompletedEvent) validateDeployment(formats strfmt.Registry) error {

	if swag.IsZero(m.Deployment) { // not required
		return nil
	}

	if m.Deployment != nil {
		if err := m.Deployment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentCompletedEvent) validateEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentCompletedEvent) validateProject(formats strfmt.Registry) error {

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
func (m *DeploymentCompletedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentCompletedEvent) UnmarshalBinary(b []byte) error {
	var res DeploymentCompletedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}