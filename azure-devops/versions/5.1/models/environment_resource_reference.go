// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EnvironmentResourceReference EnvironmentResourceReference.
// swagger:model EnvironmentResourceReference
type EnvironmentResourceReference struct {

	// Id of the resource.
	ID int32 `json:"id,omitempty"`

	// Name of the resource.
	Name string `json:"name,omitempty"`

	// Type of the resource.
	// Enum: [undefined generic virtualMachine kubernetes]
	Type interface{} `json:"type,omitempty"`
}

// Validate validates this environment resource reference
func (m *EnvironmentResourceReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentResourceReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentResourceReference) UnmarshalBinary(b []byte) error {
	var res EnvironmentResourceReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
