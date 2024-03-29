// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ArtifactVersionQueryResult artifact version query result
// swagger:model ArtifactVersionQueryResult
type ArtifactVersionQueryResult struct {

	// Gets or sets the list for artifact versions of artifact version query result.
	ArtifactVersions []*ArtifactVersion `json:"artifactVersions"`
}

// Validate validates this artifact version query result
func (m *ArtifactVersionQueryResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifactVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArtifactVersionQueryResult) validateArtifactVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.ArtifactVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.ArtifactVersions); i++ {
		if swag.IsZero(m.ArtifactVersions[i]) { // not required
			continue
		}

		if m.ArtifactVersions[i] != nil {
			if err := m.ArtifactVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("artifactVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactVersionQueryResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactVersionQueryResult) UnmarshalBinary(b []byte) error {
	var res ArtifactVersionQueryResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
