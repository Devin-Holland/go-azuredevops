// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProcessBehaviorUpdateRequest Process Behavior Replace Payload.
// swagger:model ProcessBehaviorUpdateRequest
type ProcessBehaviorUpdateRequest struct {

	// Color.
	Color string `json:"color,omitempty"`

	// Behavior Name.
	Name string `json:"name,omitempty"`
}

// Validate validates this process behavior update request
func (m *ProcessBehaviorUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProcessBehaviorUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessBehaviorUpdateRequest) UnmarshalBinary(b []byte) error {
	var res ProcessBehaviorUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
