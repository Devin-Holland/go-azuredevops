// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ArtifactSourceReference artifact source reference
// swagger:model ArtifactSourceReference
type ArtifactSourceReference struct {

	// ID of the artifact source.
	ID string `json:"id,omitempty"`

	// Name of the artifact source.
	Name string `json:"name,omitempty"`
}

// Validate validates this artifact source reference
func (m *ArtifactSourceReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactSourceReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactSourceReference) UnmarshalBinary(b []byte) error {
	var res ArtifactSourceReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
