// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProjectAvatar Contains the image data for project avatar.
// swagger:model ProjectAvatar
type ProjectAvatar struct {

	// The avatar image represented as a byte array.
	Image []strfmt.Base64 `json:"image"`
}

// Validate validates this project avatar
func (m *ProjectAvatar) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectAvatar) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	for i := 0; i < len(m.Image); i++ {

		// Format "byte" (base64 string) is already validated when unmarshalled

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectAvatar) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectAvatar) UnmarshalBinary(b []byte) error {
	var res ProjectAvatar
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
