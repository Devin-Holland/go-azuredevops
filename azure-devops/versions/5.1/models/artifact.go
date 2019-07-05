// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Artifact artifact
// swagger:model Artifact
type Artifact struct {

	// Gets or sets alias.
	Alias string `json:"alias,omitempty"`

	// Gets or sets definition reference. e.g. {"project":{"id":"fed755ea-49c5-4399-acea-fd5b5aa90a6c","name":"myProject"},"definition":{"id":"1","name":"mybuildDefinition"},"connection":{"id":"1","name":"myConnection"}}.
	DefinitionReference map[string]ArtifactSourceReference `json:"definitionReference,omitempty"`

	// Indicates whether artifact is primary or not.
	IsPrimary bool `json:"isPrimary,omitempty"`

	// Indicates whether artifact is retained by release or not.
	IsRetained bool `json:"isRetained,omitempty"`

	// Gets or sets type. It can have value as 'Build', 'Jenkins', 'GitHub', 'Nuget', 'Team Build (external)', 'ExternalTFSBuild', 'Git', 'TFVC', 'ExternalTfsXamlBuild'.
	Type string `json:"type,omitempty"`
}

// Validate validates this artifact
func (m *Artifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefinitionReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Artifact) validateDefinitionReference(formats strfmt.Registry) error {

	if swag.IsZero(m.DefinitionReference) { // not required
		return nil
	}

	for k := range m.DefinitionReference {

		if err := validate.Required("definitionReference"+"."+k, "body", m.DefinitionReference[k]); err != nil {
			return err
		}
		if val, ok := m.DefinitionReference[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Artifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Artifact) UnmarshalBinary(b []byte) error {
	var res Artifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
