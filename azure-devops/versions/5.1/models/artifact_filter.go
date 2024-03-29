// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ArtifactFilter artifact filter
// swagger:model ArtifactFilter
type ArtifactFilter struct {

	// Gets or sets whether a release should be created on build tagging.
	CreateReleaseOnBuildTagging bool `json:"createReleaseOnBuildTagging,omitempty"`

	// Gets or sets the branch for the filter.
	SourceBranch string `json:"sourceBranch,omitempty"`

	// Gets or sets the list of tags for the filter.
	Tags []string `json:"tags"`

	// Gets or sets whether filter should default to build definition branch.
	UseBuildDefinitionBranch bool `json:"useBuildDefinitionBranch,omitempty"`
}

// Validate validates this artifact filter
func (m *ArtifactFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactFilter) UnmarshalBinary(b []byte) error {
	var res ArtifactFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
