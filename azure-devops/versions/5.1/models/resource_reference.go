// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ResourceReference Represents a reference to a resource.
// swagger:model ResourceReference
type ResourceReference struct {

	// An alias to be used when referencing the resource.
	Alias string `json:"alias,omitempty"`
}

// Validate validates this resource reference
func (m *ResourceReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceReference) UnmarshalBinary(b []byte) error {
	var res ResourceReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
