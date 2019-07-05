// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestExtensionField test extension field
// swagger:model TestExtensionField
type TestExtensionField struct {

	// field
	Field *TestExtensionFieldDetails `json:"field,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this test extension field
func (m *TestExtensionField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestExtensionField) validateField(formats strfmt.Registry) error {

	if swag.IsZero(m.Field) { // not required
		return nil
	}

	if m.Field != nil {
		if err := m.Field.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("field")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestExtensionField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestExtensionField) UnmarshalBinary(b []byte) error {
	var res TestExtensionField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}