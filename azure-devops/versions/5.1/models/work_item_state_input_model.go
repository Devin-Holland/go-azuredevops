// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WorkItemStateInputModel Class That represents a work item state input.
// swagger:model WorkItemStateInputModel
type WorkItemStateInputModel struct {

	// Color of the state
	Color string `json:"color,omitempty"`

	// Name of the state
	Name string `json:"name,omitempty"`

	// Order in which state should appear
	Order int32 `json:"order,omitempty"`

	// Category of the state
	StateCategory string `json:"stateCategory,omitempty"`
}

// Validate validates this work item state input model
func (m *WorkItemStateInputModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemStateInputModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemStateInputModel) UnmarshalBinary(b []byte) error {
	var res WorkItemStateInputModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}