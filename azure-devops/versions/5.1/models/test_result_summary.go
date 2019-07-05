// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestResultSummary test result summary
// swagger:model TestResultSummary
type TestResultSummary struct {

	// aggregated results analysis
	AggregatedResultsAnalysis *AggregatedResultsAnalysis `json:"aggregatedResultsAnalysis,omitempty"`

	// no config runs count
	NoConfigRunsCount int32 `json:"noConfigRunsCount,omitempty"`

	// team project
	TeamProject *TeamProjectReference `json:"teamProject,omitempty"`

	// test failures
	TestFailures *TestFailuresAnalysis `json:"testFailures,omitempty"`

	// test results context
	TestResultsContext *TestResultsContext `json:"testResultsContext,omitempty"`

	// total runs count
	TotalRunsCount int32 `json:"totalRunsCount,omitempty"`
}

// Validate validates this test result summary
func (m *TestResultSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregatedResultsAnalysis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestFailures(formats); err != nil {
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

func (m *TestResultSummary) validateAggregatedResultsAnalysis(formats strfmt.Registry) error {

	if swag.IsZero(m.AggregatedResultsAnalysis) { // not required
		return nil
	}

	if m.AggregatedResultsAnalysis != nil {
		if err := m.AggregatedResultsAnalysis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregatedResultsAnalysis")
			}
			return err
		}
	}

	return nil
}

func (m *TestResultSummary) validateTeamProject(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamProject) { // not required
		return nil
	}

	if m.TeamProject != nil {
		if err := m.TeamProject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamProject")
			}
			return err
		}
	}

	return nil
}

func (m *TestResultSummary) validateTestFailures(formats strfmt.Registry) error {

	if swag.IsZero(m.TestFailures) { // not required
		return nil
	}

	if m.TestFailures != nil {
		if err := m.TestFailures.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testFailures")
			}
			return err
		}
	}

	return nil
}

func (m *TestResultSummary) validateTestResultsContext(formats strfmt.Registry) error {

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
func (m *TestResultSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultSummary) UnmarshalBinary(b []byte) error {
	var res TestResultSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
