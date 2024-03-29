// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AgentTargetExecutionOptions Additional options for running phases against an agent queue.
// swagger:model AgentTargetExecutionOptions
type AgentTargetExecutionOptions struct {

	// Indicates the type of execution options.
	Type int32 `json:"type,omitempty"`
}

// Validate validates this agent target execution options
func (m *AgentTargetExecutionOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AgentTargetExecutionOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentTargetExecutionOptions) UnmarshalBinary(b []byte) error {
	var res AgentTargetExecutionOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
