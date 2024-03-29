// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ArtifactURIQuery Contains criteria for querying work items based on artifact URI.
// swagger:model ArtifactUriQuery
type ArtifactURIQuery struct {

	// List of artifact URIs to use for querying work items.
	ArtifactUris []string `json:"artifactUris"`
}

// Validate validates this artifact Uri query
func (m *ArtifactURIQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactURIQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactURIQuery) UnmarshalBinary(b []byte) error {
	var res ArtifactURIQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
