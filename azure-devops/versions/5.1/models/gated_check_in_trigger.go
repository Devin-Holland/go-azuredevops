// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GatedCheckInTrigger Represents a gated check-in trigger.
// swagger:model GatedCheckInTrigger
type GatedCheckInTrigger struct {
	BuildTrigger

	// path filters
	PathFilters []string `json:"pathFilters"`

	// Indicates whether CI triggers should run after the gated check-in succeeds.
	RunContinuousIntegration bool `json:"runContinuousIntegration,omitempty"`

	// Indicates whether to take workspace mappings into account when determining whether a build should run.
	UseWorkspaceMappings bool `json:"useWorkspaceMappings,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GatedCheckInTrigger) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuildTrigger
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuildTrigger = aO0

	// now for regular properties
	var propsGatedCheckInTrigger struct {
		PathFilters []string `json:"pathFilters"`

		RunContinuousIntegration bool `json:"runContinuousIntegration,omitempty"`

		UseWorkspaceMappings bool `json:"useWorkspaceMappings,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsGatedCheckInTrigger); err != nil {
		return err
	}
	m.PathFilters = propsGatedCheckInTrigger.PathFilters

	m.RunContinuousIntegration = propsGatedCheckInTrigger.RunContinuousIntegration

	m.UseWorkspaceMappings = propsGatedCheckInTrigger.UseWorkspaceMappings

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GatedCheckInTrigger) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BuildTrigger)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsGatedCheckInTrigger struct {
		PathFilters []string `json:"pathFilters"`

		RunContinuousIntegration bool `json:"runContinuousIntegration,omitempty"`

		UseWorkspaceMappings bool `json:"useWorkspaceMappings,omitempty"`
	}
	propsGatedCheckInTrigger.PathFilters = m.PathFilters

	propsGatedCheckInTrigger.RunContinuousIntegration = m.RunContinuousIntegration

	propsGatedCheckInTrigger.UseWorkspaceMappings = m.UseWorkspaceMappings

	jsonDataPropsGatedCheckInTrigger, errGatedCheckInTrigger := swag.WriteJSON(propsGatedCheckInTrigger)
	if errGatedCheckInTrigger != nil {
		return nil, errGatedCheckInTrigger
	}
	_parts = append(_parts, jsonDataPropsGatedCheckInTrigger)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this gated check in trigger
func (m *GatedCheckInTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuildTrigger
	if err := m.BuildTrigger.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GatedCheckInTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatedCheckInTrigger) UnmarshalBinary(b []byte) error {
	var res GatedCheckInTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
