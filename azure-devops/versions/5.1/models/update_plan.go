// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UpdatePlan update plan
// swagger:model UpdatePlan
type UpdatePlan struct {

	// Description of the plan
	Description string `json:"description,omitempty"`

	// Name of the plan to create.
	Name string `json:"name,omitempty"`

	// Plan properties.
	Properties interface{} `json:"properties,omitempty"`

	// Revision of the plan that was updated - the value used here should match the one the server gave the client in the Plan.
	Revision int32 `json:"revision,omitempty"`

	// Type of the plan
	// Enum: [deliveryTimelineView]
	Type interface{} `json:"type,omitempty"`
}

// Validate validates this update plan
func (m *UpdatePlan) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePlan) UnmarshalBinary(b []byte) error {
	var res UpdatePlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
