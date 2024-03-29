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

// TestPoint Test Point Class
// swagger:model TestPoint
type TestPoint struct {

	// Comment associated to the Test Point
	Comment string `json:"comment,omitempty"`

	// Configuration associated with the Test Point
	Configuration *TestConfigurationReference `json:"configuration,omitempty"`

	// Id of the Test Point
	ID int32 `json:"id,omitempty"`

	// Variable to decide whether the test case is Active or not
	IsActive bool `json:"isActive,omitempty"`

	// Is the Test Point for Automated Test Case or Manual
	IsAutomated bool `json:"isAutomated,omitempty"`

	// Last Reset to Active Time Stamp for the Test Point
	// Format: date-time
	LastResetToActive strfmt.DateTime `json:"lastResetToActive,omitempty"`

	// Last Updated details for the Test Point
	LastUpdatedBy *IdentityRef `json:"lastUpdatedBy,omitempty"`

	// Last Update Time Stamp for the Test Point
	// Format: date-time
	LastUpdatedDate strfmt.DateTime `json:"lastUpdatedDate,omitempty"`

	// Reference links
	Links *ReferenceLinks `json:"links,omitempty"`

	// Project under which the Test Point is
	Project *TeamProjectReference `json:"project,omitempty"`

	// Results associated to the Test Point
	Results *TestPointResults `json:"results,omitempty"`

	// Test Case Reference
	TestCaseReference *TestCaseReference `json:"testCaseReference,omitempty"`

	// Test Plan under which the Test Point is
	TestPlan *TestPlanReference `json:"testPlan,omitempty"`

	// Test Suite under which the Test Point is
	TestSuite *TestSuiteReference `json:"testSuite,omitempty"`

	// Tester associated with the Test Point
	Tester *IdentityRef `json:"tester,omitempty"`
}

// Validate validates this test point
func (m *TestPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastResetToActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestCaseReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestSuite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTester(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestPoint) validateConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *TestPoint) validateLastResetToActive(formats strfmt.Registry) error {

	if swag.IsZero(m.LastResetToActive) { // not required
		return nil
	}

	if err := validate.FormatOf("lastResetToActive", "body", "date-time", m.LastResetToActive.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestPoint) validateLastUpdatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedBy) { // not required
		return nil
	}

	if m.LastUpdatedBy != nil {
		if err := m.LastUpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUpdatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TestPoint) validateLastUpdatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedDate", "body", "date-time", m.LastUpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestPoint) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *TestPoint) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *TestPoint) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	if m.Results != nil {
		if err := m.Results.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

func (m *TestPoint) validateTestCaseReference(formats strfmt.Registry) error {

	if swag.IsZero(m.TestCaseReference) { // not required
		return nil
	}

	if m.TestCaseReference != nil {
		if err := m.TestCaseReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testCaseReference")
			}
			return err
		}
	}

	return nil
}

func (m *TestPoint) validateTestPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.TestPlan) { // not required
		return nil
	}

	if m.TestPlan != nil {
		if err := m.TestPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testPlan")
			}
			return err
		}
	}

	return nil
}

func (m *TestPoint) validateTestSuite(formats strfmt.Registry) error {

	if swag.IsZero(m.TestSuite) { // not required
		return nil
	}

	if m.TestSuite != nil {
		if err := m.TestSuite.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testSuite")
			}
			return err
		}
	}

	return nil
}

func (m *TestPoint) validateTester(formats strfmt.Registry) error {

	if swag.IsZero(m.Tester) { // not required
		return nil
	}

	if m.Tester != nil {
		if err := m.Tester.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tester")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestPoint) UnmarshalBinary(b []byte) error {
	var res TestPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
