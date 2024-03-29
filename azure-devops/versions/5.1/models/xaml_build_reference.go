// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// XamlBuildReference xaml build reference
// swagger:model XamlBuildReference
type XamlBuildReference struct {

	// Id of the resource
	ID int32 `json:"id,omitempty"`

	// Name of the linked resource (definition name, controller name, etc.)
	Name string `json:"name,omitempty"`

	// Full http link to the resource
	URL string `json:"url,omitempty"`
}

// Validate validates this xaml build reference
func (m *XamlBuildReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *XamlBuildReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *XamlBuildReference) UnmarshalBinary(b []byte) error {
	var res XamlBuildReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
