// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TaskVersion task version
// swagger:model TaskVersion
type TaskVersion struct {

	// is test
	IsTest bool `json:"isTest,omitempty"`

	// major
	Major int32 `json:"major,omitempty"`

	// minor
	Minor int32 `json:"minor,omitempty"`

	// patch
	Patch int32 `json:"patch,omitempty"`
}

// Validate validates this task version
func (m *TaskVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskVersion) UnmarshalBinary(b []byte) error {
	var res TaskVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
