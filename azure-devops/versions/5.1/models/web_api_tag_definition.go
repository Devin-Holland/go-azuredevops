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

// WebAPITagDefinition The representation of a tag definition which is sent across the wire.
// swagger:model WebApiTagDefinition
type WebAPITagDefinition struct {

	// Whether or not the tag definition is active.
	Active bool `json:"active,omitempty"`

	// ID of the tag definition.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The name of the tag definition.
	Name string `json:"name,omitempty"`

	// Resource URL for the Tag Definition.
	URL string `json:"url,omitempty"`
}

// Validate validates this web Api tag definition
func (m *WebAPITagDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebAPITagDefinition) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebAPITagDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebAPITagDefinition) UnmarshalBinary(b []byte) error {
	var res WebAPITagDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}