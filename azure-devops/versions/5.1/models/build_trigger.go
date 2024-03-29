// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BuildTrigger Represents a trigger for a buld definition.
// swagger:model BuildTrigger
type BuildTrigger struct {

	// The type of the trigger.
	// Enum: [none continuousIntegration batchedContinuousIntegration schedule gatedCheckIn batchedGatedCheckIn pullRequest buildCompletion all]
	TriggerType interface{} `json:"triggerType,omitempty"`
}

// Validate validates this build trigger
func (m *BuildTrigger) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuildTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildTrigger) UnmarshalBinary(b []byte) error {
	var res BuildTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
