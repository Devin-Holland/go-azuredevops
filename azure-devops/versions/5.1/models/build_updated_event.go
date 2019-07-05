// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BuildUpdatedEvent build updated event
// swagger:model BuildUpdatedEvent
type BuildUpdatedEvent struct {
	RealtimeBuildEvent

	// build
	Build *Build `json:"build,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BuildUpdatedEvent) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RealtimeBuildEvent
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RealtimeBuildEvent = aO0

	// now for regular properties
	var propsBuildUpdatedEvent struct {
		Build *Build `json:"build,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsBuildUpdatedEvent); err != nil {
		return err
	}
	m.Build = propsBuildUpdatedEvent.Build

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BuildUpdatedEvent) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.RealtimeBuildEvent)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsBuildUpdatedEvent struct {
		Build *Build `json:"build,omitempty"`
	}
	propsBuildUpdatedEvent.Build = m.Build

	jsonDataPropsBuildUpdatedEvent, errBuildUpdatedEvent := swag.WriteJSON(propsBuildUpdatedEvent)
	if errBuildUpdatedEvent != nil {
		return nil, errBuildUpdatedEvent
	}
	_parts = append(_parts, jsonDataPropsBuildUpdatedEvent)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this build updated event
func (m *BuildUpdatedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RealtimeBuildEvent
	if err := m.RealtimeBuildEvent.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildUpdatedEvent) validateBuild(formats strfmt.Registry) error {

	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {
		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildUpdatedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildUpdatedEvent) UnmarshalBinary(b []byte) error {
	var res BuildUpdatedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}