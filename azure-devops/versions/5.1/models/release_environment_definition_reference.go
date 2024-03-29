// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ReleaseEnvironmentDefinitionReference Reference to release environment resource.
// swagger:model ReleaseEnvironmentDefinitionReference
type ReleaseEnvironmentDefinitionReference struct {

	// ID of the release definition that contains the release environment definition.
	DefinitionID int32 `json:"definitionId,omitempty"`

	// ID of the release environment definition.
	EnvironmentDefinitionID int32 `json:"environmentDefinitionId,omitempty"`
}

// Validate validates this release environment definition reference
func (m *ReleaseEnvironmentDefinitionReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseEnvironmentDefinitionReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseEnvironmentDefinitionReference) UnmarshalBinary(b []byte) error {
	var res ReleaseEnvironmentDefinitionReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
