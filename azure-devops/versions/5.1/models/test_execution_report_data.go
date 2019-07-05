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

// TestExecutionReportData test execution report data
// swagger:model TestExecutionReportData
type TestExecutionReportData struct {

	// report data
	ReportData []*DatedTestFieldData `json:"reportData"`
}

// Validate validates this test execution report data
func (m *TestExecutionReportData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestExecutionReportData) validateReportData(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportData) { // not required
		return nil
	}

	for i := 0; i < len(m.ReportData); i++ {
		if swag.IsZero(m.ReportData[i]) { // not required
			continue
		}

		if m.ReportData[i] != nil {
			if err := m.ReportData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reportData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestExecutionReportData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestExecutionReportData) UnmarshalBinary(b []byte) error {
	var res TestExecutionReportData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}