// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AggregatedResultsDifference aggregated results difference
// swagger:model AggregatedResultsDifference
type AggregatedResultsDifference struct {

	// increase in duration
	IncreaseInDuration string `json:"increaseInDuration,omitempty"`

	// increase in failures
	IncreaseInFailures int32 `json:"increaseInFailures,omitempty"`

	// increase in other tests
	IncreaseInOtherTests int32 `json:"increaseInOtherTests,omitempty"`

	// increase in passed tests
	IncreaseInPassedTests int32 `json:"increaseInPassedTests,omitempty"`

	// increase in total tests
	IncreaseInTotalTests int32 `json:"increaseInTotalTests,omitempty"`
}

// Validate validates this aggregated results difference
func (m *AggregatedResultsDifference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AggregatedResultsDifference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregatedResultsDifference) UnmarshalBinary(b []byte) error {
	var res AggregatedResultsDifference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
