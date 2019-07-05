// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// StageReference Stage in pipeline
// swagger:model StageReference
type StageReference struct {

	// Attempt number of stage
	Attempt int32 `json:"attempt,omitempty"`

	// Name of the stage. Maximum supported length for name is 256 character.
	StageName string `json:"stageName,omitempty"`
}

// Validate validates this stage reference
func (m *StageReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StageReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StageReference) UnmarshalBinary(b []byte) error {
	var res StageReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}