// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SourceIDInput source Id input
// swagger:model SourceIdInput
type SourceIDInput struct {

	// ID of source.
	ID string `json:"id,omitempty"`

	// Name of the source.
	Name string `json:"name,omitempty"`
}

// Validate validates this source Id input
func (m *SourceIDInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SourceIDInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceIDInput) UnmarshalBinary(b []byte) error {
	var res SourceIDInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
