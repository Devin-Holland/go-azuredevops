// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ScenarioSummary scenario summary
// swagger:model ScenarioSummary
type ScenarioSummary struct {

	// max user load
	MaxUserLoad int32 `json:"maxUserLoad,omitempty"`

	// min user load
	MinUserLoad int32 `json:"minUserLoad,omitempty"`

	// scenario name
	ScenarioName string `json:"scenarioName,omitempty"`
}

// Validate validates this scenario summary
func (m *ScenarioSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScenarioSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScenarioSummary) UnmarshalBinary(b []byte) error {
	var res ScenarioSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
