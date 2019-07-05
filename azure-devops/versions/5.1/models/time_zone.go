// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TimeZone time zone
// swagger:model TimeZone
type TimeZone struct {

	// Display name of the time zone.
	DisplayName string `json:"displayName,omitempty"`

	// Id of the time zone.
	ID string `json:"id,omitempty"`
}

// Validate validates this time zone
func (m *TimeZone) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeZone) UnmarshalBinary(b []byte) error {
	var res TimeZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}