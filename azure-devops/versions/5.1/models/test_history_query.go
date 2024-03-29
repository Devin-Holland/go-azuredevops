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

// TestHistoryQuery Filter to get TestCase result history.
// swagger:model TestHistoryQuery
type TestHistoryQuery struct {

	// Automated test name of the TestCase.
	AutomatedTestName string `json:"automatedTestName,omitempty"`

	// Results to be get for a particular branches.
	Branch string `json:"branch,omitempty"`

	// Get the results history only for this BuildDefinationId. This to get used in query GroupBy should be Branch. If this is provided, Branch will have no use.
	BuildDefinitionID int32 `json:"buildDefinitionId,omitempty"`

	// It will be filled by server. If not null means there are some results still to be get, and we need to call this REST API with this ContinuousToken. It is not supposed to be created (or altered, if received from server in last batch) by user.
	ContinuationToken string `json:"continuationToken,omitempty"`

	// Group the result on the basis of TestResultGroupBy. This can be Branch, Environment or null(if results are fetched by BuildDefinitionId)
	// Enum: [branch environment]
	GroupBy interface{} `json:"groupBy,omitempty"`

	// History to get between time interval MaxCompleteDate and  (MaxCompleteDate - TrendDays). Default is current date time.
	// Format: date-time
	MaxCompleteDate strfmt.DateTime `json:"maxCompleteDate,omitempty"`

	// Get the results history only for this ReleaseEnvDefinitionId. This to get used in query GroupBy should be Environment.
	ReleaseEnvDefinitionID int32 `json:"releaseEnvDefinitionId,omitempty"`

	// List of TestResultHistoryForGroup which are grouped by GroupBy
	ResultsForGroup []*TestResultHistoryForGroup `json:"resultsForGroup"`

	// Get the results history only for this testCaseId. This to get used in query to filter the result along with automatedtestname
	TestCaseID int32 `json:"testCaseId,omitempty"`

	// Number of days for which history to collect. Maximum supported value is 7 days. Default is 7 days.
	TrendDays int32 `json:"trendDays,omitempty"`
}

// Validate validates this test history query
func (m *TestHistoryQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxCompleteDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultsForGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestHistoryQuery) validateMaxCompleteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxCompleteDate) { // not required
		return nil
	}

	if err := validate.FormatOf("maxCompleteDate", "body", "date-time", m.MaxCompleteDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestHistoryQuery) validateResultsForGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultsForGroup) { // not required
		return nil
	}

	for i := 0; i < len(m.ResultsForGroup); i++ {
		if swag.IsZero(m.ResultsForGroup[i]) { // not required
			continue
		}

		if m.ResultsForGroup[i] != nil {
			if err := m.ResultsForGroup[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resultsForGroup" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestHistoryQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestHistoryQuery) UnmarshalBinary(b []byte) error {
	var res TestHistoryQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
