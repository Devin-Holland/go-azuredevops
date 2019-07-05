// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkItemQuerySortColumn A sort column.
// swagger:model WorkItemQuerySortColumn
type WorkItemQuerySortColumn struct {

	// The direction to sort by.
	Descending bool `json:"descending,omitempty"`

	// A work item field.
	Field *WorkItemFieldReference `json:"field,omitempty"`
}

// Validate validates this work item query sort column
func (m *WorkItemQuerySortColumn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkItemQuerySortColumn) validateField(formats strfmt.Registry) error {

	if swag.IsZero(m.Field) { // not required
		return nil
	}

	if m.Field != nil {
		if err := m.Field.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("field")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemQuerySortColumn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemQuerySortColumn) UnmarshalBinary(b []byte) error {
	var res WorkItemQuerySortColumn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
