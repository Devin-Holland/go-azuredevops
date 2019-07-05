// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PointWorkItemProperty Test point workitem property.
// swagger:model PointWorkItemProperty
type PointWorkItemProperty struct {

	// key value pair of test point work item property.
	WorkItem map[string]interface{} `json:"workItem,omitempty"`
}

// Validate validates this point work item property
func (m *PointWorkItemProperty) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PointWorkItemProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PointWorkItemProperty) UnmarshalBinary(b []byte) error {
	var res PointWorkItemProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
