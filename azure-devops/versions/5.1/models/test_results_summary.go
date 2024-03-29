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

// TestResultsSummary test results summary
// swagger:model TestResultsSummary
type TestResultsSummary struct {

	// overall page summary
	OverallPageSummary *PageSummary `json:"overallPageSummary,omitempty"`

	// overall request summary
	OverallRequestSummary *RequestSummary `json:"overallRequestSummary,omitempty"`

	// overall scenario summary
	OverallScenarioSummary *ScenarioSummary `json:"overallScenarioSummary,omitempty"`

	// overall test summary
	OverallTestSummary *TestSummary `json:"overallTestSummary,omitempty"`

	// overall transaction summary
	OverallTransactionSummary *TransactionSummary `json:"overallTransactionSummary,omitempty"`

	// top slow pages
	TopSlowPages []*PageSummary `json:"topSlowPages"`

	// top slow requests
	TopSlowRequests []*RequestSummary `json:"topSlowRequests"`

	// top slow tests
	TopSlowTests []*TestSummary `json:"topSlowTests"`

	// top slow transactions
	TopSlowTransactions []*TransactionSummary `json:"topSlowTransactions"`
}

// Validate validates this test results summary
func (m *TestResultsSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverallPageSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverallRequestSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverallScenarioSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverallTestSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverallTransactionSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopSlowPages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopSlowRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopSlowTests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopSlowTransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestResultsSummary) validateOverallPageSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.OverallPageSummary) { // not required
		return nil
	}

	if m.OverallPageSummary != nil {
		if err := m.OverallPageSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overallPageSummary")
			}
			return err
		}
	}

	return nil
}

func (m *TestResultsSummary) validateOverallRequestSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.OverallRequestSummary) { // not required
		return nil
	}

	if m.OverallRequestSummary != nil {
		if err := m.OverallRequestSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overallRequestSummary")
			}
			return err
		}
	}

	return nil
}

func (m *TestResultsSummary) validateOverallScenarioSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.OverallScenarioSummary) { // not required
		return nil
	}

	if m.OverallScenarioSummary != nil {
		if err := m.OverallScenarioSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overallScenarioSummary")
			}
			return err
		}
	}

	return nil
}

func (m *TestResultsSummary) validateOverallTestSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.OverallTestSummary) { // not required
		return nil
	}

	if m.OverallTestSummary != nil {
		if err := m.OverallTestSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overallTestSummary")
			}
			return err
		}
	}

	return nil
}

func (m *TestResultsSummary) validateOverallTransactionSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.OverallTransactionSummary) { // not required
		return nil
	}

	if m.OverallTransactionSummary != nil {
		if err := m.OverallTransactionSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overallTransactionSummary")
			}
			return err
		}
	}

	return nil
}

func (m *TestResultsSummary) validateTopSlowPages(formats strfmt.Registry) error {

	if swag.IsZero(m.TopSlowPages) { // not required
		return nil
	}

	for i := 0; i < len(m.TopSlowPages); i++ {
		if swag.IsZero(m.TopSlowPages[i]) { // not required
			continue
		}

		if m.TopSlowPages[i] != nil {
			if err := m.TopSlowPages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topSlowPages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestResultsSummary) validateTopSlowRequests(formats strfmt.Registry) error {

	if swag.IsZero(m.TopSlowRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.TopSlowRequests); i++ {
		if swag.IsZero(m.TopSlowRequests[i]) { // not required
			continue
		}

		if m.TopSlowRequests[i] != nil {
			if err := m.TopSlowRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topSlowRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestResultsSummary) validateTopSlowTests(formats strfmt.Registry) error {

	if swag.IsZero(m.TopSlowTests) { // not required
		return nil
	}

	for i := 0; i < len(m.TopSlowTests); i++ {
		if swag.IsZero(m.TopSlowTests[i]) { // not required
			continue
		}

		if m.TopSlowTests[i] != nil {
			if err := m.TopSlowTests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topSlowTests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestResultsSummary) validateTopSlowTransactions(formats strfmt.Registry) error {

	if swag.IsZero(m.TopSlowTransactions) { // not required
		return nil
	}

	for i := 0; i < len(m.TopSlowTransactions); i++ {
		if swag.IsZero(m.TopSlowTransactions[i]) { // not required
			continue
		}

		if m.TopSlowTransactions[i] != nil {
			if err := m.TopSlowTransactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topSlowTransactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestResultsSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultsSummary) UnmarshalBinary(b []byte) error {
	var res TestResultsSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
