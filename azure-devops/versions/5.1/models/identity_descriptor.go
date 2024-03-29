// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// IdentityDescriptor An Identity descriptor is a wrapper for the identity type (Windows SID, Passport) along with a unique identifier such as the SID or PUID.
// swagger:model IdentityDescriptor
type IdentityDescriptor struct {

	// The unique identifier for this identity, not exceeding 256 chars, which will be persisted.
	Identifier string `json:"identifier,omitempty"`

	// Type of descriptor (for example, Windows, Passport, etc.).
	IdentityType string `json:"identityType,omitempty"`
}

// Validate validates this identity descriptor
func (m *IdentityDescriptor) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IdentityDescriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityDescriptor) UnmarshalBinary(b []byte) error {
	var res IdentityDescriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
