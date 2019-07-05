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

// InstalledExtensionQuery installed extension query
// swagger:model InstalledExtensionQuery
type InstalledExtensionQuery struct {

	// asset types
	AssetTypes []string `json:"assetTypes"`

	// monikers
	Monikers []*ExtensionIdentifier `json:"monikers"`
}

// Validate validates this installed extension query
func (m *InstalledExtensionQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMonikers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstalledExtensionQuery) validateMonikers(formats strfmt.Registry) error {

	if swag.IsZero(m.Monikers) { // not required
		return nil
	}

	for i := 0; i < len(m.Monikers); i++ {
		if swag.IsZero(m.Monikers[i]) { // not required
			continue
		}

		if m.Monikers[i] != nil {
			if err := m.Monikers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("monikers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstalledExtensionQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstalledExtensionQuery) UnmarshalBinary(b []byte) error {
	var res InstalledExtensionQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
