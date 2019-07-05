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

// ExtensionState The state of an extension
// swagger:model ExtensionState
type ExtensionState struct {
	InstalledExtensionState

	// extension name
	ExtensionName string `json:"extensionName,omitempty"`

	// The time at which the version was last checked
	// Format: date-time
	LastVersionCheck strfmt.DateTime `json:"lastVersionCheck,omitempty"`

	// publisher name
	PublisherName string `json:"publisherName,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ExtensionState) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InstalledExtensionState
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InstalledExtensionState = aO0

	// now for regular properties
	var propsExtensionState struct {
		ExtensionName string `json:"extensionName,omitempty"`

		LastVersionCheck strfmt.DateTime `json:"lastVersionCheck,omitempty"`

		PublisherName string `json:"publisherName,omitempty"`

		Version string `json:"version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsExtensionState); err != nil {
		return err
	}
	m.ExtensionName = propsExtensionState.ExtensionName

	m.LastVersionCheck = propsExtensionState.LastVersionCheck

	m.PublisherName = propsExtensionState.PublisherName

	m.Version = propsExtensionState.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ExtensionState) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.InstalledExtensionState)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsExtensionState struct {
		ExtensionName string `json:"extensionName,omitempty"`

		LastVersionCheck strfmt.DateTime `json:"lastVersionCheck,omitempty"`

		PublisherName string `json:"publisherName,omitempty"`

		Version string `json:"version,omitempty"`
	}
	propsExtensionState.ExtensionName = m.ExtensionName

	propsExtensionState.LastVersionCheck = m.LastVersionCheck

	propsExtensionState.PublisherName = m.PublisherName

	propsExtensionState.Version = m.Version

	jsonDataPropsExtensionState, errExtensionState := swag.WriteJSON(propsExtensionState)
	if errExtensionState != nil {
		return nil, errExtensionState
	}
	_parts = append(_parts, jsonDataPropsExtensionState)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this extension state
func (m *ExtensionState) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InstalledExtensionState
	if err := m.InstalledExtensionState.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastVersionCheck(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionState) validateLastVersionCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.LastVersionCheck) { // not required
		return nil
	}

	if err := validate.FormatOf("lastVersionCheck", "body", "date-time", m.LastVersionCheck.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionState) UnmarshalBinary(b []byte) error {
	var res ExtensionState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
