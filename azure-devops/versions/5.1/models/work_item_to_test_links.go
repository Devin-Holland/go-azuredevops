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

// WorkItemToTestLinks work item to test links
// swagger:model WorkItemToTestLinks
type WorkItemToTestLinks struct {

	// executed in
	// Enum: [any tcm tfs]
	ExecutedIn interface{} `json:"executedIn,omitempty"`

	// tests
	Tests []*TestMethod `json:"tests"`

	// work item
	WorkItem *WorkItemReference `json:"workItem,omitempty"`
}

// Validate validates this work item to test links
func (m *WorkItemToTestLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTests(formats); err != nil {
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

func (m *WorkItemToTestLinks) validateTests(formats strfmt.Registry) error {

	if swag.IsZero(m.Tests) { // not required
		return nil
	}

	for i := 0; i < len(m.Tests); i++ {
		if swag.IsZero(m.Tests[i]) { // not required
			continue
		}

		if m.Tests[i] != nil {
			if err := m.Tests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkItemToTestLinks) validateWorkItem(formats strfmt.Registry) error {

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
func (m *WorkItemToTestLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemToTestLinks) UnmarshalBinary(b []byte) error {
	var res WorkItemToTestLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
