// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MavenPomGav maven pom gav
// swagger:model MavenPomGav
type MavenPomGav struct {

	// artifact Id
	ArtifactID string `json:"artifactId,omitempty"`

	// group Id
	GroupID string `json:"groupId,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this maven pom gav
func (m *MavenPomGav) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MavenPomGav) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MavenPomGav) UnmarshalBinary(b []byte) error {
	var res MavenPomGav
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
