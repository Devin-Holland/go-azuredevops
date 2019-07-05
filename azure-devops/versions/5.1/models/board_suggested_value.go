// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BoardSuggestedValue board suggested value
// swagger:model BoardSuggestedValue
type BoardSuggestedValue struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this board suggested value
func (m *BoardSuggestedValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BoardSuggestedValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BoardSuggestedValue) UnmarshalBinary(b []byte) error {
	var res BoardSuggestedValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
