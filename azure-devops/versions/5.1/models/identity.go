// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Identity identity
// swagger:model Identity
type Identity struct {
	IdentityBase
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Identity) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IdentityBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IdentityBase = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Identity) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.IdentityBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this identity
func (m *Identity) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with IdentityBase
	if err := m.IdentityBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Identity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identity) UnmarshalBinary(b []byte) error {
	var res Identity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
