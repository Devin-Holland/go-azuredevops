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

// TestCase Test Case Class
// swagger:model TestCase
type TestCase struct {

	// Reference links
	Links *ReferenceLinks `json:"links,omitempty"`

	// Order of the TestCase in the Suite
	Order int32 `json:"order,omitempty"`

	// List of Points associated with the Test Case
	PointAssignments []*PointAssignment `json:"pointAssignments"`

	// Project under which the Test Case is
	Project *TeamProjectReference `json:"project,omitempty"`

	// Test Plan under which the Test Case is
	TestPlan *TestPlanReference `json:"testPlan,omitempty"`

	// Test Suite under which the Test Case is
	TestSuite *TestSuiteReference `json:"testSuite,omitempty"`

	// Work Item details of the TestCase
	WorkItem *WorkItemDetails `json:"workItem,omitempty"`
}

// Validate validates this test case
func (m *TestCase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePointAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestSuite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCase) validateLinks(formats strfmt.Registry) error {

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

func (m *TestCase) validatePointAssignments(formats strfmt.Registry) error {

	if swag.IsZero(m.PointAssignments) { // not required
		return nil
	}

	for i := 0; i < len(m.PointAssignments); i++ {
		if swag.IsZero(m.PointAssignments[i]) { // not required
			continue
		}

		if m.PointAssignments[i] != nil {
			if err := m.PointAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pointAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestCase) validateProject(formats strfmt.Registry) error {

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

func (m *TestCase) validateTestPlan(formats strfmt.Registry) error {

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

func (m *TestCase) validateTestSuite(formats strfmt.Registry) error {

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

func (m *TestCase) validateWorkItem(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkItem) { // not required
		return nil
	}

	if m.WorkItem != nil {
		if err := m.WorkItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workItem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCase) UnmarshalBinary(b []byte) error {
	var res TestCase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}