// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BacklogFields backlog fields
// swagger:model BacklogFields
type BacklogFields struct {

	// Field Type (e.g. Order, Activity) to Field Reference Name map
	TypeFields map[string]string `json:"typeFields,omitempty"`
}

// Validate validates this backlog fields
func (m *BacklogFields) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BacklogFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BacklogFields) UnmarshalBinary(b []byte) error {
	var res BacklogFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
