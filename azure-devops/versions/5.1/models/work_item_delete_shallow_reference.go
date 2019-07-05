// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WorkItemDeleteShallowReference Shallow Reference to a deleted work item.
// swagger:model WorkItemDeleteShallowReference
type WorkItemDeleteShallowReference struct {

	// Work item ID.
	ID int32 `json:"id,omitempty"`

	// REST API URL of the resource
	URL string `json:"url,omitempty"`
}

// Validate validates this work item delete shallow reference
func (m *WorkItemDeleteShallowReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemDeleteShallowReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemDeleteShallowReference) UnmarshalBinary(b []byte) error {
	var res WorkItemDeleteShallowReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
