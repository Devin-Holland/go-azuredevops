// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VirtualMachine virtual machine
// swagger:model VirtualMachine
type VirtualMachine struct {

	// agent
	Agent *TaskAgent `json:"agent,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this virtual machine
func (m *VirtualMachine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachine) validateAgent(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *VirtualMachine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachine) UnmarshalBinary(b []byte) error {
	var res VirtualMachine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}