// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FieldReference An abstracted reference to a field
// swagger:model FieldReference
type FieldReference struct {

	// fieldRefName for the field
	ReferenceName string `json:"referenceName,omitempty"`

	// Full http link to more information about the field
	URL string `json:"url,omitempty"`
}

// Validate validates this field reference
func (m *FieldReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FieldReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FieldReference) UnmarshalBinary(b []byte) error {
	var res FieldReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
