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

// IgnoredGate ignored gate
// swagger:model IgnoredGate
type IgnoredGate struct {

	// Gets the date on which gate is last ignored.
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"lastModifiedOn,omitempty"`

	// Name of gate ignored.
	Name string `json:"name,omitempty"`
}

// Validate validates this ignored gate
func (m *IgnoredGate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgnoredGate) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgnoredGate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgnoredGate) UnmarshalBinary(b []byte) error {
	var res IgnoredGate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
