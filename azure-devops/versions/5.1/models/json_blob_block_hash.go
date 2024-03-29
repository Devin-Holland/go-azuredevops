// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// JSONBlobBlockHash BlobBlock hash formatted to be deserialized for symbol service.
// swagger:model JsonBlobBlockHash
type JSONBlobBlockHash struct {

	// Array of hash bytes.
	HashBytes []strfmt.Base64 `json:"hashBytes"`
}

// Validate validates this Json blob block hash
func (m *JSONBlobBlockHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHashBytes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONBlobBlockHash) validateHashBytes(formats strfmt.Registry) error {

	if swag.IsZero(m.HashBytes) { // not required
		return nil
	}

	for i := 0; i < len(m.HashBytes); i++ {

		// Format "byte" (base64 string) is already validated when unmarshalled

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JSONBlobBlockHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONBlobBlockHash) UnmarshalBinary(b []byte) error {
	var res JSONBlobBlockHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
