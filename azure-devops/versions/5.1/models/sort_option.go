// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SortOption Defines how to sort the result.
// swagger:model SortOption
type SortOption struct {

	// Field name on which sorting should be done.
	Field string `json:"field,omitempty"`

	// Order (ASC/DESC) in which the results should be sorted.
	SortOrder string `json:"sortOrder,omitempty"`
}

// Validate validates this sort option
func (m *SortOption) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SortOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SortOption) UnmarshalBinary(b []byte) error {
	var res SortOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
