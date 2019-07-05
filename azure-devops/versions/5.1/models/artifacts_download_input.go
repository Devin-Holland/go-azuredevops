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

// ArtifactsDownloadInput artifacts download input
// swagger:model ArtifactsDownloadInput
type ArtifactsDownloadInput struct {

	// download inputs
	DownloadInputs []*ArtifactDownloadInputBase `json:"downloadInputs"`
}

// Validate validates this artifacts download input
func (m *ArtifactsDownloadInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownloadInputs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArtifactsDownloadInput) validateDownloadInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.DownloadInputs) { // not required
		return nil
	}

	for i := 0; i < len(m.DownloadInputs); i++ {
		if swag.IsZero(m.DownloadInputs[i]) { // not required
			continue
		}

		if m.DownloadInputs[i] != nil {
			if err := m.DownloadInputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("downloadInputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactsDownloadInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactsDownloadInput) UnmarshalBinary(b []byte) error {
	var res ArtifactsDownloadInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
