// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestMessageLog2 test message log2
// swagger:model TestMessageLog2
type TestMessageLog2 struct {

	// test message log Id
	TestMessageLogID int32 `json:"testMessageLogId,omitempty"`
}

// Validate validates this test message log2
func (m *TestMessageLog2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestMessageLog2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestMessageLog2) UnmarshalBinary(b []byte) error {
	var res TestMessageLog2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
