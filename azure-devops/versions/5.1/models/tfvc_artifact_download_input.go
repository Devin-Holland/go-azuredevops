// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TfvcArtifactDownloadInput tfvc artifact download input
// swagger:model TfvcArtifactDownloadInput
type TfvcArtifactDownloadInput struct {
	ArtifactDownloadInputBase
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TfvcArtifactDownloadInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ArtifactDownloadInputBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ArtifactDownloadInputBase = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TfvcArtifactDownloadInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ArtifactDownloadInputBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tfvc artifact download input
func (m *TfvcArtifactDownloadInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ArtifactDownloadInputBase
	if err := m.ArtifactDownloadInputBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TfvcArtifactDownloadInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TfvcArtifactDownloadInput) UnmarshalBinary(b []byte) error {
	var res TfvcArtifactDownloadInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
