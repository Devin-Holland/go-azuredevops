// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestFieldData test field data
// swagger:model TestFieldData
type TestFieldData struct {

	// dimensions
	Dimensions map[string]interface{} `json:"dimensions,omitempty"`

	// measure
	Measure int64 `json:"measure,omitempty"`
}

// Validate validates this test field data
func (m *TestFieldData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestFieldData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestFieldData) UnmarshalBinary(b []byte) error {
	var res TestFieldData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}