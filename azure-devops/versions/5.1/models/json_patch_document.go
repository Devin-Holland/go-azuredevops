// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// JSONPatchDocument The JSON model for JSON Patch Operations
// swagger:model JsonPatchDocument
type JSONPatchDocument struct {
	JSONPatchDocumentAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *JSONPatchDocument) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 JSONPatchDocumentAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.JSONPatchDocumentAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m JSONPatchDocument) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.JSONPatchDocumentAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this Json patch document
func (m *JSONPatchDocument) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with JSONPatchDocumentAllOf0
	if err := m.JSONPatchDocumentAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *JSONPatchDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONPatchDocument) UnmarshalBinary(b []byte) error {
	var res JSONPatchDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JSONPatchDocumentAllOf0 JSON patch document all of0
// swagger:model JSONPatchDocumentAllOf0
type JSONPatchDocumentAllOf0 []*JSONPatchOperation

// Validate validates this JSON patch document all of0
func (m JSONPatchDocumentAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
