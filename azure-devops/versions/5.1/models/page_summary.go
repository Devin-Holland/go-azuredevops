// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PageSummary page summary
// swagger:model PageSummary
type PageSummary struct {

	// average page time
	AveragePageTime float64 `json:"averagePageTime,omitempty"`

	// page Url
	PageURL string `json:"pageUrl,omitempty"`

	// percentage pages meeting goal
	PercentagePagesMeetingGoal int32 `json:"percentagePagesMeetingGoal,omitempty"`

	// percentile data
	PercentileData []*SummaryPercentileData `json:"percentileData"`

	// scenario name
	ScenarioName string `json:"scenarioName,omitempty"`

	// test name
	TestName string `json:"testName,omitempty"`

	// total pages
	TotalPages int32 `json:"totalPages,omitempty"`
}

// Validate validates this page summary
func (m *PageSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePercentileData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageSummary) validatePercentileData(formats strfmt.Registry) error {

	if swag.IsZero(m.PercentileData) { // not required
		return nil
	}

	for i := 0; i < len(m.PercentileData); i++ {
		if swag.IsZero(m.PercentileData[i]) { // not required
			continue
		}

		if m.PercentileData[i] != nil {
			if err := m.PercentileData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("percentileData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PageSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageSummary) UnmarshalBinary(b []byte) error {
	var res PageSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
