// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExtensionRequestUrls extension request urls
// swagger:model ExtensionRequestUrls
type ExtensionRequestUrls struct {
	ExtensionUrls

	// Link to view the extension request
	RequestPage string `json:"requestPage,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ExtensionRequestUrls) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ExtensionUrls
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ExtensionUrls = aO0

	// now for regular properties
	var propsExtensionRequestUrls struct {
		RequestPage string `json:"requestPage,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsExtensionRequestUrls); err != nil {
		return err
	}
	m.RequestPage = propsExtensionRequestUrls.RequestPage

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ExtensionRequestUrls) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ExtensionUrls)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsExtensionRequestUrls struct {
		RequestPage string `json:"requestPage,omitempty"`
	}
	propsExtensionRequestUrls.RequestPage = m.RequestPage

	jsonDataPropsExtensionRequestUrls, errExtensionRequestUrls := swag.WriteJSON(propsExtensionRequestUrls)
	if errExtensionRequestUrls != nil {
		return nil, errExtensionRequestUrls
	}
	_parts = append(_parts, jsonDataPropsExtensionRequestUrls)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this extension request urls
func (m *ExtensionRequestUrls) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ExtensionUrls
	if err := m.ExtensionUrls.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionRequestUrls) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionRequestUrls) UnmarshalBinary(b []byte) error {
	var res ExtensionRequestUrls
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}