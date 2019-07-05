// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DesignerProcessTarget Represents the target for the build process.
// swagger:model DesignerProcessTarget
type DesignerProcessTarget struct {

	// Agent specification for the build process.
	AgentSpecification *AgentSpecification `json:"agentSpecification,omitempty"`
}

// Validate validates this designer process target
func (m *DesignerProcessTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentSpecification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DesignerProcessTarget) validateAgentSpecification(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentSpecification) { // not required
		return nil
	}

	if m.AgentSpecification != nil {
		if err := m.AgentSpecification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agentSpecification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DesignerProcessTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DesignerProcessTarget) UnmarshalBinary(b []byte) error {
	var res DesignerProcessTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}