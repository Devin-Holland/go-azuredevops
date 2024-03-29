// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DiagnosticLogMetadata diagnostic log metadata
// swagger:model DiagnosticLogMetadata
type DiagnosticLogMetadata struct {

	// agent Id
	AgentID int32 `json:"agentId,omitempty"`

	// agent name
	AgentName string `json:"agentName,omitempty"`

	// file name
	FileName string `json:"fileName,omitempty"`

	// phase name
	PhaseName string `json:"phaseName,omitempty"`

	// phase result
	PhaseResult string `json:"phaseResult,omitempty"`

	// pool Id
	PoolID int32 `json:"poolId,omitempty"`
}

// Validate validates this diagnostic log metadata
func (m *DiagnosticLogMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DiagnosticLogMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiagnosticLogMetadata) UnmarshalBinary(b []byte) error {
	var res DiagnosticLogMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
