// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AddProcessWorkItemTypeFieldRequest Class that describes a request to add a field in a work item type.
// swagger:model AddProcessWorkItemTypeFieldRequest
type AddProcessWorkItemTypeFieldRequest struct {

	// Allow setting field value to a group identity. Only applies to identity fields.
	AllowGroups bool `json:"allowGroups,omitempty"`

	// The default value of the field.
	DefaultValue interface{} `json:"defaultValue,omitempty"`

	// If true the field cannot be edited.
	ReadOnly bool `json:"readOnly,omitempty"`

	// Reference name of the field.
	ReferenceName string `json:"referenceName,omitempty"`

	// If true the field cannot be empty.
	Required bool `json:"required,omitempty"`
}

// Validate validates this add process work item type field request
func (m *AddProcessWorkItemTypeFieldRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddProcessWorkItemTypeFieldRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddProcessWorkItemTypeFieldRequest) UnmarshalBinary(b []byte) error {
	var res AddProcessWorkItemTypeFieldRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}