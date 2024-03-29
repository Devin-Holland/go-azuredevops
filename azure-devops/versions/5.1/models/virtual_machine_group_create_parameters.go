// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VirtualMachineGroupCreateParameters virtual machine group create parameters
// swagger:model VirtualMachineGroupCreateParameters
type VirtualMachineGroupCreateParameters struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this virtual machine group create parameters
func (m *VirtualMachineGroupCreateParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineGroupCreateParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineGroupCreateParameters) UnmarshalBinary(b []byte) error {
	var res VirtualMachineGroupCreateParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
