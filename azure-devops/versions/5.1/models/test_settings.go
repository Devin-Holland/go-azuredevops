// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestSettings Represents the test settings of the run. Used to create test settings and fetch test settings
// swagger:model TestSettings
type TestSettings struct {

	// Area path required to create test settings
	AreaPath string `json:"areaPath,omitempty"`

	// Description of the test settings. Used in create test settings.
	Description string `json:"description,omitempty"`

	// Indicates if the tests settings is public or private.Used in create test settings.
	IsPublic bool `json:"isPublic,omitempty"`

	// Xml string of machine roles. Used in create test settings.
	MachineRoles string `json:"machineRoles,omitempty"`

	// Test settings content.
	TestSettingsContent string `json:"testSettingsContent,omitempty"`

	// Test settings id.
	TestSettingsID int32 `json:"testSettingsId,omitempty"`

	// Test settings name.
	TestSettingsName string `json:"testSettingsName,omitempty"`
}

// Validate validates this test settings
func (m *TestSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestSettings) UnmarshalBinary(b []byte) error {
	var res TestSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
