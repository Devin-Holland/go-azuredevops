// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IdentityReference Describes a reference to an identity.
// swagger:model IdentityReference
type IdentityReference struct {
	IdentityRef

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Legacy back-compat property. This has been the WIT specific value from Constants. Will be hidden (but exists) on the client unless they are targeting the newest version
	Name string `json:"name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IdentityReference) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IdentityRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IdentityRef = aO0

	// now for regular properties
	var propsIdentityReference struct {
		ID strfmt.UUID `json:"id,omitempty"`

		Name string `json:"name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsIdentityReference); err != nil {
		return err
	}
	m.ID = propsIdentityReference.ID

	m.Name = propsIdentityReference.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IdentityReference) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.IdentityRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsIdentityReference struct {
		ID strfmt.UUID `json:"id,omitempty"`

		Name string `json:"name,omitempty"`
	}
	propsIdentityReference.ID = m.ID

	propsIdentityReference.Name = m.Name

	jsonDataPropsIdentityReference, errIdentityReference := swag.WriteJSON(propsIdentityReference)
	if errIdentityReference != nil {
		return nil, errIdentityReference
	}
	_parts = append(_parts, jsonDataPropsIdentityReference)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this identity reference
func (m *IdentityReference) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with IdentityRef
	if err := m.IdentityRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityReference) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentityReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityReference) UnmarshalBinary(b []byte) error {
	var res IdentityReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
