// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WorkItemBehaviorReference Reference to the behavior of a work item type.
// swagger:model WorkItemBehaviorReference
type WorkItemBehaviorReference struct {

	// The ID of the reference behavior.
	ID string `json:"id,omitempty"`

	// The url of the reference behavior.
	URL string `json:"url,omitempty"`
}

// Validate validates this work item behavior reference
func (m *WorkItemBehaviorReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemBehaviorReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemBehaviorReference) UnmarshalBinary(b []byte) error {
	var res WorkItemBehaviorReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
