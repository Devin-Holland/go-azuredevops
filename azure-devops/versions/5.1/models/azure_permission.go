// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AzurePermission azure permission
// swagger:model AzurePermission
type AzurePermission struct {

	// provisioned
	Provisioned bool `json:"provisioned,omitempty"`

	// resource provider
	ResourceProvider string `json:"resourceProvider,omitempty"`
}

// Validate validates this azure permission
func (m *AzurePermission) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzurePermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzurePermission) UnmarshalBinary(b []byte) error {
	var res AzurePermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}