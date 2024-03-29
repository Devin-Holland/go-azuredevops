// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InputValuesError Error information related to a subscription input value.
// swagger:model InputValuesError
type InputValuesError struct {

	// The error message.
	Message string `json:"message,omitempty"`
}

// Validate validates this input values error
func (m *InputValuesError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InputValuesError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InputValuesError) UnmarshalBinary(b []byte) error {
	var res InputValuesError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
