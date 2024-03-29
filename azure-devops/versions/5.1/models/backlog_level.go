// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BacklogLevel Contract representing a backlog level
// swagger:model BacklogLevel
type BacklogLevel struct {

	// Reference name of the corresponding WIT category
	CategoryReferenceName string `json:"categoryReferenceName,omitempty"`

	// Plural name for the backlog level
	PluralName string `json:"pluralName,omitempty"`

	// Collection of work item states that are included in the plan. The server will filter to only these work item types.
	WorkItemStates []string `json:"workItemStates"`

	// Collection of valid workitem type names for the given backlog level
	WorkItemTypes []string `json:"workItemTypes"`
}

// Validate validates this backlog level
func (m *BacklogLevel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BacklogLevel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BacklogLevel) UnmarshalBinary(b []byte) error {
	var res BacklogLevel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
