// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TeamFieldValue Represents a single TeamFieldValue
// swagger:model TeamFieldValue
type TeamFieldValue struct {

	// include children
	IncludeChildren bool `json:"includeChildren,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this team field value
func (m *TeamFieldValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TeamFieldValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamFieldValue) UnmarshalBinary(b []byte) error {
	var res TeamFieldValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
