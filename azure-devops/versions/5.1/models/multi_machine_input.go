// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MultiMachineInput multi machine input
// swagger:model MultiMachineInput
type MultiMachineInput struct {
	ParallelExecutionInputBase
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MultiMachineInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ParallelExecutionInputBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ParallelExecutionInputBase = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MultiMachineInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ParallelExecutionInputBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this multi machine input
func (m *MultiMachineInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ParallelExecutionInputBase
	if err := m.ParallelExecutionInputBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MultiMachineInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiMachineInput) UnmarshalBinary(b []byte) error {
	var res MultiMachineInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
