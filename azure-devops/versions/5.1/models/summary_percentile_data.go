// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SummaryPercentileData summary percentile data
// swagger:model SummaryPercentileData
type SummaryPercentileData struct {

	// percentile
	Percentile int32 `json:"percentile,omitempty"`

	// percentile value
	PercentileValue float64 `json:"percentileValue,omitempty"`
}

// Validate validates this summary percentile data
func (m *SummaryPercentileData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SummaryPercentileData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryPercentileData) UnmarshalBinary(b []byte) error {
	var res SummaryPercentileData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}