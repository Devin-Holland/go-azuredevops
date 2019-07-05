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

// TestRunEx2 test run ex2
// swagger:model TestRunEx2
type TestRunEx2 struct {

	// bit value
	BitValue bool `json:"bitValue,omitempty"`

	// created date
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// date time value
	// Format: date-time
	DateTimeValue strfmt.DateTime `json:"dateTimeValue,omitempty"`

	// field Id
	FieldID int32 `json:"fieldId,omitempty"`

	// field name
	FieldName string `json:"fieldName,omitempty"`

	// float value
	FloatValue float64 `json:"floatValue,omitempty"`

	// guid value
	// Format: uuid
	GUIDValue strfmt.UUID `json:"guidValue,omitempty"`

	// int value
	IntValue int32 `json:"intValue,omitempty"`

	// project Id
	// Format: uuid
	ProjectID strfmt.UUID `json:"projectId,omitempty"`

	// string value
	StringValue string `json:"stringValue,omitempty"`

	// test run Id
	TestRunID int32 `json:"testRunId,omitempty"`
}

// Validate validates this test run ex2
func (m *TestRunEx2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTimeValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGUIDValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestRunEx2) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRunEx2) validateDateTimeValue(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTimeValue) { // not required
		return nil
	}

	if err := validate.FormatOf("dateTimeValue", "body", "date-time", m.DateTimeValue.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRunEx2) validateGUIDValue(formats strfmt.Registry) error {

	if swag.IsZero(m.GUIDValue) { // not required
		return nil
	}

	if err := validate.FormatOf("guidValue", "body", "uuid", m.GUIDValue.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRunEx2) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectId", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestRunEx2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestRunEx2) UnmarshalBinary(b []byte) error {
	var res TestRunEx2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}