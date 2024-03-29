// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TenantDetails tenant details
// swagger:model TenantDetails
type TenantDetails struct {

	// access details
	AccessDetails []*AgentGroupAccessData `json:"accessDetails"`

	// id
	ID string `json:"id,omitempty"`

	// static machines
	StaticMachines []*WebAPITestMachine `json:"staticMachines"`

	// user load agent input
	UserLoadAgentInput *WebAPIUserLoadTestMachineInput `json:"userLoadAgentInput,omitempty"`

	// user load agent resources Uri
	UserLoadAgentResourcesURI string `json:"userLoadAgentResourcesUri,omitempty"`

	// valid geo locations
	ValidGeoLocations []string `json:"validGeoLocations"`
}

// Validate validates this tenant details
func (m *TenantDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticMachines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserLoadAgentInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantDetails) validateAccessDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessDetails); i++ {
		if swag.IsZero(m.AccessDetails[i]) { // not required
			continue
		}

		if m.AccessDetails[i] != nil {
			if err := m.AccessDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TenantDetails) validateStaticMachines(formats strfmt.Registry) error {

	if swag.IsZero(m.StaticMachines) { // not required
		return nil
	}

	for i := 0; i < len(m.StaticMachines); i++ {
		if swag.IsZero(m.StaticMachines[i]) { // not required
			continue
		}

		if m.StaticMachines[i] != nil {
			if err := m.StaticMachines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("staticMachines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TenantDetails) validateUserLoadAgentInput(formats strfmt.Registry) error {

	if swag.IsZero(m.UserLoadAgentInput) { // not required
		return nil
	}

	if m.UserLoadAgentInput != nil {
		if err := m.UserLoadAgentInput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userLoadAgentInput")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TenantDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantDetails) UnmarshalBinary(b []byte) error {
	var res TenantDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
