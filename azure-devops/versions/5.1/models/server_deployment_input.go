// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerDeploymentInput server deployment input
// swagger:model ServerDeploymentInput
type ServerDeploymentInput struct {
	BaseDeploymentInput

	// Gets or sets the parallel execution input.
	ParallelExecution *ExecutionInput `json:"parallelExecution,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ServerDeploymentInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseDeploymentInput
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseDeploymentInput = aO0

	// now for regular properties
	var propsServerDeploymentInput struct {
		ParallelExecution *ExecutionInput `json:"parallelExecution,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsServerDeploymentInput); err != nil {
		return err
	}
	m.ParallelExecution = propsServerDeploymentInput.ParallelExecution

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ServerDeploymentInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BaseDeploymentInput)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsServerDeploymentInput struct {
		ParallelExecution *ExecutionInput `json:"parallelExecution,omitempty"`
	}
	propsServerDeploymentInput.ParallelExecution = m.ParallelExecution

	jsonDataPropsServerDeploymentInput, errServerDeploymentInput := swag.WriteJSON(propsServerDeploymentInput)
	if errServerDeploymentInput != nil {
		return nil, errServerDeploymentInput
	}
	_parts = append(_parts, jsonDataPropsServerDeploymentInput)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this server deployment input
func (m *ServerDeploymentInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseDeploymentInput
	if err := m.BaseDeploymentInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParallelExecution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerDeploymentInput) validateParallelExecution(formats strfmt.Registry) error {

	if swag.IsZero(m.ParallelExecution) { // not required
		return nil
	}

	if m.ParallelExecution != nil {
		if err := m.ParallelExecution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parallelExecution")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerDeploymentInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerDeploymentInput) UnmarshalBinary(b []byte) error {
	var res ServerDeploymentInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}