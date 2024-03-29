// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WorkItemReference Contains reference to a work item.
// swagger:model WorkItemReference
type WorkItemReference struct {

	// Work item ID.
	ID int32 `json:"id,omitempty"`

	// REST API URL of the resource
	URL string `json:"url,omitempty"`
}

// Validate validates this work item reference
func (m *WorkItemReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemReference) UnmarshalBinary(b []byte) error {
	var res WorkItemReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
