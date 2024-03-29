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

// ExtensionLicensing How an extension should handle including contributions based on licensing
// swagger:model ExtensionLicensing
type ExtensionLicensing struct {

	// A list of contributions which deviate from the default licensing behavior
	Overrides []*LicensingOverride `json:"overrides"`
}

// Validate validates this extension licensing
func (m *ExtensionLicensing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionLicensing) validateOverrides(formats strfmt.Registry) error {

	if swag.IsZero(m.Overrides) { // not required
		return nil
	}

	for i := 0; i < len(m.Overrides); i++ {
		if swag.IsZero(m.Overrides[i]) { // not required
			continue
		}

		if m.Overrides[i] != nil {
			if err := m.Overrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionLicensing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionLicensing) UnmarshalBinary(b []byte) error {
	var res ExtensionLicensing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
