// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ExtensionFile extension file
// swagger:model ExtensionFile
type ExtensionFile struct {

	// asset type
	AssetType string `json:"assetType,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// source
	Source string `json:"source,omitempty"`
}

// Validate validates this extension file
func (m *ExtensionFile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionFile) UnmarshalBinary(b []byte) error {
	var res ExtensionFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}