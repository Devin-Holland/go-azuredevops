// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AzureKeyVaultPermission azure key vault permission
// swagger:model AzureKeyVaultPermission
type AzureKeyVaultPermission struct {
	AzureResourcePermission

	// vault
	Vault string `json:"vault,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AzureKeyVaultPermission) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AzureResourcePermission
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AzureResourcePermission = aO0

	// now for regular properties
	var propsAzureKeyVaultPermission struct {
		Vault string `json:"vault,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsAzureKeyVaultPermission); err != nil {
		return err
	}
	m.Vault = propsAzureKeyVaultPermission.Vault

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AzureKeyVaultPermission) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.AzureResourcePermission)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsAzureKeyVaultPermission struct {
		Vault string `json:"vault,omitempty"`
	}
	propsAzureKeyVaultPermission.Vault = m.Vault

	jsonDataPropsAzureKeyVaultPermission, errAzureKeyVaultPermission := swag.WriteJSON(propsAzureKeyVaultPermission)
	if errAzureKeyVaultPermission != nil {
		return nil, errAzureKeyVaultPermission
	}
	_parts = append(_parts, jsonDataPropsAzureKeyVaultPermission)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this azure key vault permission
func (m *AzureKeyVaultPermission) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AzureResourcePermission
	if err := m.AzureResourcePermission.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AzureKeyVaultPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureKeyVaultPermission) UnmarshalBinary(b []byte) error {
	var res AzureKeyVaultPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
