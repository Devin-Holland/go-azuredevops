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

// TestResultTrendFilter test result trend filter
// swagger:model TestResultTrendFilter
type TestResultTrendFilter struct {

	// branch names
	BranchNames []string `json:"branchNames"`

	// build count
	BuildCount int32 `json:"buildCount,omitempty"`

	// definition ids
	DefinitionIds []int32 `json:"definitionIds"`

	// env definition ids
	EnvDefinitionIds []int32 `json:"envDefinitionIds"`

	// max complete date
	// Format: date-time
	MaxCompleteDate strfmt.DateTime `json:"maxCompleteDate,omitempty"`

	// publish context
	PublishContext string `json:"publishContext,omitempty"`

	// test run titles
	TestRunTitles []string `json:"testRunTitles"`

	// trend days
	TrendDays int32 `json:"trendDays,omitempty"`
}

// Validate validates this test result trend filter
func (m *TestResultTrendFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxCompleteDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestResultTrendFilter) validateMaxCompleteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxCompleteDate) { // not required
		return nil
	}

	if err := validate.FormatOf("maxCompleteDate", "body", "date-time", m.MaxCompleteDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestResultTrendFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultTrendFilter) UnmarshalBinary(b []byte) error {
	var res TestResultTrendFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
