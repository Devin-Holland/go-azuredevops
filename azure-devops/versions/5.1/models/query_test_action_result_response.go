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

// QueryTestActionResultResponse query test action result response
// swagger:model QueryTestActionResultResponse
type QueryTestActionResultResponse struct {

	// test action results
	TestActionResults []*TestActionResult `json:"testActionResults"`

	// test attachments
	TestAttachments []*TestResultAttachment `json:"testAttachments"`

	// test result parameters
	TestResultParameters []*TestResultParameter `json:"testResultParameters"`
}

// Validate validates this query test action result response
func (m *QueryTestActionResultResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTestActionResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestResultParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryTestActionResultResponse) validateTestActionResults(formats strfmt.Registry) error {

	if swag.IsZero(m.TestActionResults) { // not required
		return nil
	}

	for i := 0; i < len(m.TestActionResults); i++ {
		if swag.IsZero(m.TestActionResults[i]) { // not required
			continue
		}

		if m.TestActionResults[i] != nil {
			if err := m.TestActionResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testActionResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryTestActionResultResponse) validateTestAttachments(formats strfmt.Registry) error {

	if swag.IsZero(m.TestAttachments) { // not required
		return nil
	}

	for i := 0; i < len(m.TestAttachments); i++ {
		if swag.IsZero(m.TestAttachments[i]) { // not required
			continue
		}

		if m.TestAttachments[i] != nil {
			if err := m.TestAttachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testAttachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryTestActionResultResponse) validateTestResultParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.TestResultParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.TestResultParameters); i++ {
		if swag.IsZero(m.TestResultParameters[i]) { // not required
			continue
		}

		if m.TestResultParameters[i] != nil {
			if err := m.TestResultParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testResultParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryTestActionResultResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryTestActionResultResponse) UnmarshalBinary(b []byte) error {
	var res QueryTestActionResultResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
