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

// ArtifactSourceIdsQueryResult artifact source ids query result
// swagger:model ArtifactSourceIdsQueryResult
type ArtifactSourceIdsQueryResult struct {

	// Gets or sets the list of artifactsourceIds.
	ArtifactSourceIds []*ArtifactSourceID `json:"artifactSourceIds"`
}

// Validate validates this artifact source ids query result
func (m *ArtifactSourceIdsQueryResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifactSourceIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArtifactSourceIdsQueryResult) validateArtifactSourceIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ArtifactSourceIds) { // not required
		return nil
	}

	for i := 0; i < len(m.ArtifactSourceIds); i++ {
		if swag.IsZero(m.ArtifactSourceIds[i]) { // not required
			continue
		}

		if m.ArtifactSourceIds[i] != nil {
			if err := m.ArtifactSourceIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("artifactSourceIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactSourceIdsQueryResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactSourceIdsQueryResult) UnmarshalBinary(b []byte) error {
	var res ArtifactSourceIdsQueryResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
