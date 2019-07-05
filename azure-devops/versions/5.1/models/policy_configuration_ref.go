// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PolicyConfigurationRef Policy configuration reference.
// swagger:model PolicyConfigurationRef
type PolicyConfigurationRef struct {

	// The policy configuration ID.
	ID int32 `json:"id,omitempty"`

	// The policy configuration type.
	Type *PolicyTypeRef `json:"type,omitempty"`

	// The URL where the policy configuration can be retrieved.
	URL string `json:"url,omitempty"`
}

// Validate validates this policy configuration ref
func (m *PolicyConfigurationRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyConfigurationRef) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyConfigurationRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyConfigurationRef) UnmarshalBinary(b []byte) error {
	var res PolicyConfigurationRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}