// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AgentChangeEvent agent change event
// swagger:model AgentChangeEvent
type AgentChangeEvent struct {

	// agent
	Agent *TaskAgent `json:"agent,omitempty"`

	// event type
	EventType string `json:"eventType,omitempty"`

	// pool
	Pool *TaskAgentPoolReference `json:"pool,omitempty"`
}

// Validate validates this agent change event
func (m *AgentChangeEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentChangeEvent) validateAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *AgentChangeEvent) validatePool(formats strfmt.Registry) error {

	if swag.IsZero(m.Pool) { // not required
		return nil
	}

	if m.Pool != nil {
		if err := m.Pool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentChangeEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentChangeEvent) UnmarshalBinary(b []byte) error {
	var res AgentChangeEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
