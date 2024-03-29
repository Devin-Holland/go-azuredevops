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

// ResourceFilterOptions resource filter options
// swagger:model ResourceFilterOptions
type ResourceFilterOptions struct {

	// identities
	Identities []*IdentityRef `json:"identities"`

	// resource types
	ResourceTypes []string `json:"resourceTypes"`
}

// Validate validates this resource filter options
func (m *ResourceFilterOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceFilterOptions) validateIdentities(formats strfmt.Registry) error {

	if swag.IsZero(m.Identities) { // not required
		return nil
	}

	for i := 0; i < len(m.Identities); i++ {
		if swag.IsZero(m.Identities[i]) { // not required
			continue
		}

		if m.Identities[i] != nil {
			if err := m.Identities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("identities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceFilterOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceFilterOptions) UnmarshalBinary(b []byte) error {
	var res ResourceFilterOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
