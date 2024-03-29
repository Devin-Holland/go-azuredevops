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

// ActionDefinition action definition
// swagger:model ActionDefinition
type ActionDefinition struct {

	// The bit mask integer for this action. Must be a power of 2.
	Bit int32 `json:"bit,omitempty"`

	// The localized display name for this action.
	DisplayName string `json:"displayName,omitempty"`

	// The non-localized name for this action.
	Name string `json:"name,omitempty"`

	// The namespace that this action belongs to.  This will only be used for reading from the database.
	// Format: uuid
	NamespaceID strfmt.UUID `json:"namespaceId,omitempty"`
}

// Validate validates this action definition
func (m *ActionDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespaceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionDefinition) validateNamespaceID(formats strfmt.Registry) error {

	if swag.IsZero(m.NamespaceID) { // not required
		return nil
	}

	if err := validate.FormatOf("namespaceId", "body", "uuid", m.NamespaceID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionDefinition) UnmarshalBinary(b []byte) error {
	var res ActionDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
