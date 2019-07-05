// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DeploymentGroupCreateParameter Properties to create Deployment group.
// swagger:model DeploymentGroupCreateParameter
type DeploymentGroupCreateParameter struct {

	// Description of the deployment group.
	Description string `json:"description,omitempty"`

	// Name of the deployment group.
	Name string `json:"name,omitempty"`

	// Identifier of the deployment pool in which deployment agents are registered.
	PoolID int32 `json:"poolId,omitempty"`
}

// Validate validates this deployment group create parameter
func (m *DeploymentGroupCreateParameter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentGroupCreateParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentGroupCreateParameter) UnmarshalBinary(b []byte) error {
	var res DeploymentGroupCreateParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
