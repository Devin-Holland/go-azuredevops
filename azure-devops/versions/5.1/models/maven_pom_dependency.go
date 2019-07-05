// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MavenPomDependency maven pom dependency
// swagger:model MavenPomDependency
type MavenPomDependency struct {
	MavenPomGav

	// optional
	Optional bool `json:"optional,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MavenPomDependency) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MavenPomGav
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MavenPomGav = aO0

	// now for regular properties
	var propsMavenPomDependency struct {
		Optional bool `json:"optional,omitempty"`

		Scope string `json:"scope,omitempty"`

		Type string `json:"type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsMavenPomDependency); err != nil {
		return err
	}
	m.Optional = propsMavenPomDependency.Optional

	m.Scope = propsMavenPomDependency.Scope

	m.Type = propsMavenPomDependency.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MavenPomDependency) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.MavenPomGav)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsMavenPomDependency struct {
		Optional bool `json:"optional,omitempty"`

		Scope string `json:"scope,omitempty"`

		Type string `json:"type,omitempty"`
	}
	propsMavenPomDependency.Optional = m.Optional

	propsMavenPomDependency.Scope = m.Scope

	propsMavenPomDependency.Type = m.Type

	jsonDataPropsMavenPomDependency, errMavenPomDependency := swag.WriteJSON(propsMavenPomDependency)
	if errMavenPomDependency != nil {
		return nil, errMavenPomDependency
	}
	_parts = append(_parts, jsonDataPropsMavenPomDependency)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this maven pom dependency
func (m *MavenPomDependency) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MavenPomGav
	if err := m.MavenPomGav.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MavenPomDependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MavenPomDependency) UnmarshalBinary(b []byte) error {
	var res MavenPomDependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}