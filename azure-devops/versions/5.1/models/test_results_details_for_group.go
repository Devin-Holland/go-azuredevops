// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestResultsDetailsForGroup test results details for group
// swagger:model TestResultsDetailsForGroup
type TestResultsDetailsForGroup struct {

	// group by value
	GroupByValue interface{} `json:"groupByValue,omitempty"`

	// results
	Results []*TestCaseResult `json:"results"`

	// results count by outcome
	ResultsCountByOutcome map[string]AggregatedResultsByOutcome `json:"resultsCountByOutcome,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this test results details for group
func (m *TestResultsDetailsForGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultsCountByOutcome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestResultsDetailsForGroup) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestResultsDetailsForGroup) validateResultsCountByOutcome(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultsCountByOutcome) { // not required
		return nil
	}

	for k := range m.ResultsCountByOutcome {

		if err := validate.Required("resultsCountByOutcome"+"."+k, "body", m.ResultsCountByOutcome[k]); err != nil {
			return err
		}
		if val, ok := m.ResultsCountByOutcome[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestResultsDetailsForGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultsDetailsForGroup) UnmarshalBinary(b []byte) error {
	var res TestResultsDetailsForGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}