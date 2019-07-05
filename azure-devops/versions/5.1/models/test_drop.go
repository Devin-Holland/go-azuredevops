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

// TestDrop test drop
// swagger:model TestDrop
type TestDrop struct {

	// access data
	AccessData *DropAccessData `json:"accessData,omitempty"`

	// created date
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// drop type
	DropType string `json:"dropType,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// load test definition
	LoadTestDefinition *LoadTestDefinition `json:"loadTestDefinition,omitempty"`

	// test run Id
	TestRunID string `json:"testRunId,omitempty"`
}

// Validate validates this test drop
func (m *TestDrop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadTestDefinition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestDrop) validateAccessData(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessData) { // not required
		return nil
	}

	if m.AccessData != nil {
		if err := m.AccessData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessData")
			}
			return err
		}
	}

	return nil
}

func (m *TestDrop) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestDrop) validateLoadTestDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadTestDefinition) { // not required
		return nil
	}

	if m.LoadTestDefinition != nil {
		if err := m.LoadTestDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loadTestDefinition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestDrop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestDrop) UnmarshalBinary(b []byte) error {
	var res TestDrop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
