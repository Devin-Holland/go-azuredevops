// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TaskAgentReference A reference to an agent.
// swagger:model TaskAgentReference
type TaskAgentReference struct {

	// links
	Links *ReferenceLinks `json:"_links,omitempty"`

	// This agent's access point.
	AccessPoint string `json:"accessPoint,omitempty"`

	// Whether or not this agent should run jobs.
	Enabled bool `json:"enabled,omitempty"`

	// Identifier of the agent.
	ID int32 `json:"id,omitempty"`

	// Name of the agent.
	Name string `json:"name,omitempty"`

	// Agent OS.
	OsDescription string `json:"osDescription,omitempty"`

	// Provisioning state of this agent.
	ProvisioningState string `json:"provisioningState,omitempty"`

	// Whether or not the agent is online.
	// Enum: [offline online]
	Status interface{} `json:"status,omitempty"`

	// Agent version.
	Version string `json:"version,omitempty"`
}

// Validate validates this task agent reference
func (m *TaskAgentReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAgentReference) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAgentReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAgentReference) UnmarshalBinary(b []byte) error {
	var res TaskAgentReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
