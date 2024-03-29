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

// CreateTestMessageLogEntryRequest create test message log entry request
// swagger:model CreateTestMessageLogEntryRequest
type CreateTestMessageLogEntryRequest struct {

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// test message log entry
	TestMessageLogEntry []*TestMessageLogEntry `json:"testMessageLogEntry"`

	// test run Id
	TestRunID int32 `json:"testRunId,omitempty"`
}

// Validate validates this create test message log entry request
func (m *CreateTestMessageLogEntryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTestMessageLogEntry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTestMessageLogEntryRequest) validateTestMessageLogEntry(formats strfmt.Registry) error {

	if swag.IsZero(m.TestMessageLogEntry) { // not required
		return nil
	}

	for i := 0; i < len(m.TestMessageLogEntry); i++ {
		if swag.IsZero(m.TestMessageLogEntry[i]) { // not required
			continue
		}

		if m.TestMessageLogEntry[i] != nil {
			if err := m.TestMessageLogEntry[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testMessageLogEntry" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTestMessageLogEntryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTestMessageLogEntryRequest) UnmarshalBinary(b []byte) error {
	var res CreateTestMessageLogEntryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
