// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WebAPIUserLoadTestMachineInput web Api user load test machine input
// swagger:model WebApiUserLoadTestMachineInput
type WebAPIUserLoadTestMachineInput struct {
	WebAPILoadTestMachineInput

	// agent group name
	AgentGroupName string `json:"agentGroupName,omitempty"`

	// tenant Id
	TenantID string `json:"tenantId,omitempty"`

	// user load agent resources Uri
	UserLoadAgentResourcesURI string `json:"userLoadAgentResourcesUri,omitempty"`

	// vsts account Uri
	VstsAccountURI string `json:"vstsAccountUri,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WebAPIUserLoadTestMachineInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WebAPILoadTestMachineInput
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WebAPILoadTestMachineInput = aO0

	// now for regular properties
	var propsWebAPIUserLoadTestMachineInput struct {
		AgentGroupName string `json:"agentGroupName,omitempty"`

		TenantID string `json:"tenantId,omitempty"`

		UserLoadAgentResourcesURI string `json:"userLoadAgentResourcesUri,omitempty"`

		VstsAccountURI string `json:"vstsAccountUri,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsWebAPIUserLoadTestMachineInput); err != nil {
		return err
	}
	m.AgentGroupName = propsWebAPIUserLoadTestMachineInput.AgentGroupName

	m.TenantID = propsWebAPIUserLoadTestMachineInput.TenantID

	m.UserLoadAgentResourcesURI = propsWebAPIUserLoadTestMachineInput.UserLoadAgentResourcesURI

	m.VstsAccountURI = propsWebAPIUserLoadTestMachineInput.VstsAccountURI

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WebAPIUserLoadTestMachineInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WebAPILoadTestMachineInput)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsWebAPIUserLoadTestMachineInput struct {
		AgentGroupName string `json:"agentGroupName,omitempty"`

		TenantID string `json:"tenantId,omitempty"`

		UserLoadAgentResourcesURI string `json:"userLoadAgentResourcesUri,omitempty"`

		VstsAccountURI string `json:"vstsAccountUri,omitempty"`
	}
	propsWebAPIUserLoadTestMachineInput.AgentGroupName = m.AgentGroupName

	propsWebAPIUserLoadTestMachineInput.TenantID = m.TenantID

	propsWebAPIUserLoadTestMachineInput.UserLoadAgentResourcesURI = m.UserLoadAgentResourcesURI

	propsWebAPIUserLoadTestMachineInput.VstsAccountURI = m.VstsAccountURI

	jsonDataPropsWebAPIUserLoadTestMachineInput, errWebAPIUserLoadTestMachineInput := swag.WriteJSON(propsWebAPIUserLoadTestMachineInput)
	if errWebAPIUserLoadTestMachineInput != nil {
		return nil, errWebAPIUserLoadTestMachineInput
	}
	_parts = append(_parts, jsonDataPropsWebAPIUserLoadTestMachineInput)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this web Api user load test machine input
func (m *WebAPIUserLoadTestMachineInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WebAPILoadTestMachineInput
	if err := m.WebAPILoadTestMachineInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WebAPIUserLoadTestMachineInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebAPIUserLoadTestMachineInput) UnmarshalBinary(b []byte) error {
	var res WebAPIUserLoadTestMachineInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}