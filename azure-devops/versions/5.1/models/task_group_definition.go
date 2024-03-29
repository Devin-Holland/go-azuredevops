// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TaskGroupDefinition task group definition
// swagger:model TaskGroupDefinition
type TaskGroupDefinition struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// is expanded
	IsExpanded bool `json:"isExpanded,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// visible rule
	VisibleRule string `json:"visibleRule,omitempty"`
}

// Validate validates this task group definition
func (m *TaskGroupDefinition) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskGroupDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskGroupDefinition) UnmarshalBinary(b []byte) error {
	var res TaskGroupDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
