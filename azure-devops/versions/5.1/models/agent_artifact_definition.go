// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AgentArtifactDefinition agent artifact definition
// swagger:model AgentArtifactDefinition
type AgentArtifactDefinition struct {

	// Gets or sets the artifact definition alias.
	Alias string `json:"alias,omitempty"`

	// Gets or sets the artifact type.
	// Enum: [xamlBuild build jenkins fileShare nuget tfsOnPrem gitHub tfGit externalTfsBuild custom tfvc]
	ArtifactType interface{} `json:"artifactType,omitempty"`

	// Gets or sets the artifact definition details.
	Details string `json:"details,omitempty"`

	// Gets or sets the name of artifact definition.
	Name string `json:"name,omitempty"`

	// Gets or sets the version of artifact definition.
	Version string `json:"version,omitempty"`
}

// Validate validates this agent artifact definition
func (m *AgentArtifactDefinition) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AgentArtifactDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentArtifactDefinition) UnmarshalBinary(b []byte) error {
	var res AgentArtifactDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
