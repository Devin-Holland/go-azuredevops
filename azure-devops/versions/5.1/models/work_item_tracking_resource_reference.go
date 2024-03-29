// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WorkItemTrackingResourceReference Base class for work item tracking resource references.
// swagger:model WorkItemTrackingResourceReference
type WorkItemTrackingResourceReference struct {

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this work item tracking resource reference
func (m *WorkItemTrackingResourceReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemTrackingResourceReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemTrackingResourceReference) UnmarshalBinary(b []byte) error {
	var res WorkItemTrackingResourceReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
