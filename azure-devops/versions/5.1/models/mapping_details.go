// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MappingDetails mapping details
// swagger:model MappingDetails
type MappingDetails struct {

	// mappings
	Mappings map[string]InputValue `json:"mappings,omitempty"`
}

// Validate validates this mapping details
func (m *MappingDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMappings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MappingDetails) validateMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.Mappings) { // not required
		return nil
	}

	for k := range m.Mappings {

		if err := validate.Required("mappings"+"."+k, "body", m.Mappings[k]); err != nil {
			return err
		}
		if val, ok := m.Mappings[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MappingDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MappingDetails) UnmarshalBinary(b []byte) error {
	var res MappingDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
