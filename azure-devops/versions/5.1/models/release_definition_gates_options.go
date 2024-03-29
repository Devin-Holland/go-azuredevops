// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ReleaseDefinitionGatesOptions release definition gates options
// swagger:model ReleaseDefinitionGatesOptions
type ReleaseDefinitionGatesOptions struct {

	// Gets or sets as the gates enabled or not.
	IsEnabled bool `json:"isEnabled,omitempty"`

	// Gets or sets the minimum duration for steady results after a successful gates evaluation.
	MinimumSuccessDuration int32 `json:"minimumSuccessDuration,omitempty"`

	// Gets or sets the time between re-evaluation of gates.
	SamplingInterval int32 `json:"samplingInterval,omitempty"`

	// Gets or sets the delay before evaluation.
	StabilizationTime int32 `json:"stabilizationTime,omitempty"`

	// Gets or sets the timeout after which gates fail.
	Timeout int32 `json:"timeout,omitempty"`
}

// Validate validates this release definition gates options
func (m *ReleaseDefinitionGatesOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseDefinitionGatesOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseDefinitionGatesOptions) UnmarshalBinary(b []byte) error {
	var res ReleaseDefinitionGatesOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
