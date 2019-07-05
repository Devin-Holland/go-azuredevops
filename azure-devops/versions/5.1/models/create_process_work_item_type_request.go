// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreateProcessWorkItemTypeRequest Class for create work item type request
// swagger:model CreateProcessWorkItemTypeRequest
type CreateProcessWorkItemTypeRequest struct {

	// Color hexadecimal code to represent the work item type
	Color string `json:"color,omitempty"`

	// Description of the work item type
	Description string `json:"description,omitempty"`

	// Icon to represent the work item type
	Icon string `json:"icon,omitempty"`

	// Parent work item type for work item type
	InheritsFrom string `json:"inheritsFrom,omitempty"`

	// True if the work item type need to be disabled
	IsDisabled bool `json:"isDisabled,omitempty"`

	// Name of work item type
	Name string `json:"name,omitempty"`
}

// Validate validates this create process work item type request
func (m *CreateProcessWorkItemTypeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateProcessWorkItemTypeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateProcessWorkItemTypeRequest) UnmarshalBinary(b []byte) error {
	var res CreateProcessWorkItemTypeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}