// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TaskAgentPoolMaintenanceOptions task agent pool maintenance options
// swagger:model TaskAgentPoolMaintenanceOptions
type TaskAgentPoolMaintenanceOptions struct {

	// time to consider a System.DefaultWorkingDirectory is stale
	WorkingDirectoryExpirationInDays int32 `json:"workingDirectoryExpirationInDays,omitempty"`
}

// Validate validates this task agent pool maintenance options
func (m *TaskAgentPoolMaintenanceOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskAgentPoolMaintenanceOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAgentPoolMaintenanceOptions) UnmarshalBinary(b []byte) error {
	var res TaskAgentPoolMaintenanceOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
