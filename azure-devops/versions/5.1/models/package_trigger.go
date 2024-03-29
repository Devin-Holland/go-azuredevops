// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PackageTrigger package trigger
// swagger:model PackageTrigger
type PackageTrigger struct {
	ReleaseTriggerBase

	// Package trigger alias.
	Alias string `json:"alias,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PackageTrigger) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ReleaseTriggerBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ReleaseTriggerBase = aO0

	// now for regular properties
	var propsPackageTrigger struct {
		Alias string `json:"alias,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsPackageTrigger); err != nil {
		return err
	}
	m.Alias = propsPackageTrigger.Alias

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PackageTrigger) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ReleaseTriggerBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsPackageTrigger struct {
		Alias string `json:"alias,omitempty"`
	}
	propsPackageTrigger.Alias = m.Alias

	jsonDataPropsPackageTrigger, errPackageTrigger := swag.WriteJSON(propsPackageTrigger)
	if errPackageTrigger != nil {
		return nil, errPackageTrigger
	}
	_parts = append(_parts, jsonDataPropsPackageTrigger)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this package trigger
func (m *PackageTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ReleaseTriggerBase
	if err := m.ReleaseTriggerBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PackageTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageTrigger) UnmarshalBinary(b []byte) error {
	var res PackageTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
