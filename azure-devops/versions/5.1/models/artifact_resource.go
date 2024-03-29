// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ArtifactResource artifact resource
// swagger:model ArtifactResource
type ArtifactResource struct {

	// links
	Links *ReferenceLinks `json:"_links,omitempty"`

	// Type-specific data about the artifact.
	Data string `json:"data,omitempty"`

	// A link to download the resource.
	DownloadURL string `json:"downloadUrl,omitempty"`

	// Type-specific properties of the artifact.
	Properties map[string]string `json:"properties,omitempty"`

	// The type of the resource: File container, version control folder, UNC path, etc.
	Type string `json:"type,omitempty"`

	// The full http link to the resource.
	URL string `json:"url,omitempty"`
}

// Validate validates this artifact resource
func (m *ArtifactResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArtifactResource) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactResource) UnmarshalBinary(b []byte) error {
	var res ArtifactResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
