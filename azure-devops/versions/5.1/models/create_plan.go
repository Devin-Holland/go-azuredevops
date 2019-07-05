// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreatePlan create plan
// swagger:model CreatePlan
type CreatePlan struct {

	// Description of the plan
	Description string `json:"description,omitempty"`

	// Name of the plan to create.
	Name string `json:"name,omitempty"`

	// Plan properties.
	Properties interface{} `json:"properties,omitempty"`

	// Type of plan to create.
	// Enum: [deliveryTimelineView]
	Type interface{} `json:"type,omitempty"`
}

// Validate validates this create plan
func (m *CreatePlan) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreatePlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePlan) UnmarshalBinary(b []byte) error {
	var res CreatePlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
