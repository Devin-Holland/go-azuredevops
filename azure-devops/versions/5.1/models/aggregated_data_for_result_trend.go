// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AggregatedDataForResultTrend aggregated data for result trend
// swagger:model AggregatedDataForResultTrend
type AggregatedDataForResultTrend struct {

	// This is tests execution duration.
	Duration string `json:"duration,omitempty"`

	// results by outcome
	ResultsByOutcome map[string]AggregatedResultsByOutcome `json:"resultsByOutcome,omitempty"`

	// run summary by state
	RunSummaryByState map[string]AggregatedRunsByState `json:"runSummaryByState,omitempty"`

	// test results context
	TestResultsContext *TestResultsContext `json:"testResultsContext,omitempty"`

	// total tests
	TotalTests int32 `json:"totalTests,omitempty"`
}

// Validate validates this aggregated data for result trend
func (m *AggregatedDataForResultTrend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResultsByOutcome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunSummaryByState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestResultsContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregatedDataForResultTrend) validateResultsByOutcome(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultsByOutcome) { // not required
		return nil
	}

	for k := range m.ResultsByOutcome {

		if err := validate.Required("resultsByOutcome"+"."+k, "body", m.ResultsByOutcome[k]); err != nil {
			return err
		}
		if val, ok := m.ResultsByOutcome[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *AggregatedDataForResultTrend) validateRunSummaryByState(formats strfmt.Registry) error {

	if swag.IsZero(m.RunSummaryByState) { // not required
		return nil
	}

	for k := range m.RunSummaryByState {

		if err := validate.Required("runSummaryByState"+"."+k, "body", m.RunSummaryByState[k]); err != nil {
			return err
		}
		if val, ok := m.RunSummaryByState[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *AggregatedDataForResultTrend) validateTestResultsContext(formats strfmt.Registry) error {

	if swag.IsZero(m.TestResultsContext) { // not required
		return nil
	}

	if m.TestResultsContext != nil {
		if err := m.TestResultsContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testResultsContext")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AggregatedDataForResultTrend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregatedDataForResultTrend) UnmarshalBinary(b []byte) error {
	var res AggregatedDataForResultTrend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
