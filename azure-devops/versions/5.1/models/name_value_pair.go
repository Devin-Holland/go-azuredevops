// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NameValuePair Name value pair
// swagger:model NameValuePair
type NameValuePair struct {

	// Name
	Name string `json:"name,omitempty"`

	// Value
	Value string `json:"value,omitempty"`
}

// Validate validates this name value pair
func (m *NameValuePair) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NameValuePair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameValuePair) UnmarshalBinary(b []byte) error {
	var res NameValuePair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
