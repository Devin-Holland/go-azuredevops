// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Issue issue
// swagger:model Issue
type Issue struct {

	// Issue data.
	Data map[string]string `json:"data,omitempty"`

	// Issue type, for example error, warning or info.
	IssueType string `json:"issueType,omitempty"`

	// Issue message.
	Message string `json:"message,omitempty"`
}

// Validate validates this issue
func (m *Issue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Issue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Issue) UnmarshalBinary(b []byte) error {
	var res Issue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
