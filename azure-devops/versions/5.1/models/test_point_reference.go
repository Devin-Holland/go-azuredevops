// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestPointReference test point reference
// swagger:model TestPointReference
type TestPointReference struct {

	// id
	ID int32 `json:"id,omitempty"`

	// state
	// Enum: [none ready completed notReady inProgress maxValue]
	State interface{} `json:"state,omitempty"`
}

// Validate validates this test point reference
func (m *TestPointReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestPointReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestPointReference) UnmarshalBinary(b []byte) error {
	var res TestPointReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
