// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReleaseCreatedEvent release created event
// swagger:model ReleaseCreatedEvent
type ReleaseCreatedEvent struct {

	// project
	Project *ProjectReference `json:"project,omitempty"`

	// release
	Release *Release `json:"release,omitempty"`
}

// Validate validates this release created event
func (m *ReleaseCreatedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseCreatedEvent) validateProject(formats strfmt.Registry) error {

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

func (m *ReleaseCreatedEvent) validateRelease(formats strfmt.Registry) error {

	if swag.IsZero(m.Release) { // not required
		return nil
	}

	if m.Release != nil {
		if err := m.Release.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("release")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseCreatedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseCreatedEvent) UnmarshalBinary(b []byte) error {
	var res ReleaseCreatedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}