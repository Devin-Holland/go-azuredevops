// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BuildEvent build event
// swagger:model BuildEvent
type BuildEvent struct {

	// data
	Data []string `json:"data"`

	// identifier
	Identifier string `json:"identifier,omitempty"`
}

// Validate validates this build event
func (m *BuildEvent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuildEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildEvent) UnmarshalBinary(b []byte) error {
	var res BuildEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}