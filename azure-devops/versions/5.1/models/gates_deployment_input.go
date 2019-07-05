// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GatesDeploymentInput gates deployment input
// swagger:model GatesDeploymentInput
type GatesDeploymentInput struct {
	BaseDeploymentInput

	// Gates minimum success duration.
	MinimumSuccessDuration int32 `json:"minimumSuccessDuration,omitempty"`

	// Gates sampling interval.
	SamplingInterval int32 `json:"samplingInterval,omitempty"`

	// Gates stabilization time.
	StabilizationTime int32 `json:"stabilizationTime,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GatesDeploymentInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseDeploymentInput
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseDeploymentInput = aO0

	// now for regular properties
	var propsGatesDeploymentInput struct {
		MinimumSuccessDuration int32 `json:"minimumSuccessDuration,omitempty"`

		SamplingInterval int32 `json:"samplingInterval,omitempty"`

		StabilizationTime int32 `json:"stabilizationTime,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsGatesDeploymentInput); err != nil {
		return err
	}
	m.MinimumSuccessDuration = propsGatesDeploymentInput.MinimumSuccessDuration

	m.SamplingInterval = propsGatesDeploymentInput.SamplingInterval

	m.StabilizationTime = propsGatesDeploymentInput.StabilizationTime

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GatesDeploymentInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BaseDeploymentInput)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsGatesDeploymentInput struct {
		MinimumSuccessDuration int32 `json:"minimumSuccessDuration,omitempty"`

		SamplingInterval int32 `json:"samplingInterval,omitempty"`

		StabilizationTime int32 `json:"stabilizationTime,omitempty"`
	}
	propsGatesDeploymentInput.MinimumSuccessDuration = m.MinimumSuccessDuration

	propsGatesDeploymentInput.SamplingInterval = m.SamplingInterval

	propsGatesDeploymentInput.StabilizationTime = m.StabilizationTime

	jsonDataPropsGatesDeploymentInput, errGatesDeploymentInput := swag.WriteJSON(propsGatesDeploymentInput)
	if errGatesDeploymentInput != nil {
		return nil, errGatesDeploymentInput
	}
	_parts = append(_parts, jsonDataPropsGatesDeploymentInput)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this gates deployment input
func (m *GatesDeploymentInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseDeploymentInput
	if err := m.BaseDeploymentInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GatesDeploymentInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatesDeploymentInput) UnmarshalBinary(b []byte) error {
	var res GatesDeploymentInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
