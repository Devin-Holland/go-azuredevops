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

// ReleaseDefinitionGatesStep release definition gates step
// swagger:model ReleaseDefinitionGatesStep
type ReleaseDefinitionGatesStep struct {

	// Gets or sets the gates.
	Gates []*ReleaseDefinitionGate `json:"gates"`

	// Gets or sets the gate options.
	GatesOptions *ReleaseDefinitionGatesOptions `json:"gatesOptions,omitempty"`

	// ID of the ReleaseDefinitionGateStep.
	ID int32 `json:"id,omitempty"`
}

// Validate validates this release definition gates step
func (m *ReleaseDefinitionGatesStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatesOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseDefinitionGatesStep) validateGates(formats strfmt.Registry) error {

	if swag.IsZero(m.Gates) { // not required
		return nil
	}

	for i := 0; i < len(m.Gates); i++ {
		if swag.IsZero(m.Gates[i]) { // not required
			continue
		}

		if m.Gates[i] != nil {
			if err := m.Gates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseDefinitionGatesStep) validateGatesOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.GatesOptions) { // not required
		return nil
	}

	if m.GatesOptions != nil {
		if err := m.GatesOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gatesOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseDefinitionGatesStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseDefinitionGatesStep) UnmarshalBinary(b []byte) error {
	var res ReleaseDefinitionGatesStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
