// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PhaseTarget Represents the target of a phase.
// swagger:model PhaseTarget
type PhaseTarget struct {

	// The type of the target.
	Type int32 `json:"type,omitempty"`
}

// Validate validates this phase target
func (m *PhaseTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PhaseTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhaseTarget) UnmarshalBinary(b []byte) error {
	var res PhaseTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
