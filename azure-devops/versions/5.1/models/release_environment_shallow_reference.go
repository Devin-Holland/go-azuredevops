// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReleaseEnvironmentShallowReference release environment shallow reference
// swagger:model ReleaseEnvironmentShallowReference
type ReleaseEnvironmentShallowReference struct {

	// Gets the links to related resources, APIs, and views for the release environment.
	Links *ReferenceLinks `json:"_links,omitempty"`

	// Gets the unique identifier of release environment.
	ID int32 `json:"id,omitempty"`

	// Gets or sets the name of the release environment.
	Name string `json:"name,omitempty"`

	// Gets the REST API url to access the release environment.
	URL string `json:"url,omitempty"`
}

// Validate validates this release environment shallow reference
func (m *ReleaseEnvironmentShallowReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseEnvironmentShallowReference) validateLinks(formats strfmt.Registry) error {

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
func (m *ReleaseEnvironmentShallowReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseEnvironmentShallowReference) UnmarshalBinary(b []byte) error {
	var res ReleaseEnvironmentShallowReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
