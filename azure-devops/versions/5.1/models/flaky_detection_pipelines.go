// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FlakyDetectionPipelines flaky detection pipelines
// swagger:model FlakyDetectionPipelines
type FlakyDetectionPipelines struct {

	// AllowedPipelines - List All Pipelines allowed for detection.
	AllowedPipelines []int32 `json:"allowedPipelines"`

	// IsAllPipelinesAllowed if users configure all system's pipelines.
	IsAllPipelinesAllowed bool `json:"isAllPipelinesAllowed,omitempty"`
}

// Validate validates this flaky detection pipelines
func (m *FlakyDetectionPipelines) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlakyDetectionPipelines) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlakyDetectionPipelines) UnmarshalBinary(b []byte) error {
	var res FlakyDetectionPipelines
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
