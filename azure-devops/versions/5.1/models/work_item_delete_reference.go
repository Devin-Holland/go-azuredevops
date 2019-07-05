// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WorkItemDeleteReference Reference to a deleted work item.
// swagger:model WorkItemDeleteReference
type WorkItemDeleteReference struct {

	// The HTTP status code for work item operation in a batch request.
	Code int32 `json:"code,omitempty"`

	// The user who deleted the work item type.
	DeletedBy string `json:"deletedBy,omitempty"`

	// The work item deletion date.
	DeletedDate string `json:"deletedDate,omitempty"`

	// Work item ID.
	ID int32 `json:"id,omitempty"`

	// The exception message for work item operation in a batch request.
	Message string `json:"message,omitempty"`

	// Name or title of the work item.
	Name string `json:"name,omitempty"`

	// Parent project of the deleted work item.
	Project string `json:"project,omitempty"`

	// Type of work item.
	Type string `json:"type,omitempty"`

	// REST API URL of the resource
	URL string `json:"url,omitempty"`
}

// Validate validates this work item delete reference
func (m *WorkItemDeleteReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemDeleteReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemDeleteReference) UnmarshalBinary(b []byte) error {
	var res WorkItemDeleteReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
