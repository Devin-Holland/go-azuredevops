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

// TestDefinitionBasic test definition basic
// swagger:model TestDefinitionBasic
type TestDefinitionBasic struct {

	// access data
	AccessData *DropAccessData `json:"accessData,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// created date
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// last modified date
	// Format: date-time
	LastModifiedDate strfmt.DateTime `json:"lastModifiedDate,omitempty"`

	// load test type
	// Enum: [visualStudioLoadTest jMeter oldLoadTestFile]
	LoadTestType interface{} `json:"loadTestType,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this test definition basic
func (m *TestDefinitionBasic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestDefinitionBasic) validateAccessData(formats strfmt.Registry) error {

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

func (m *TestDefinitionBasic) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestDefinitionBasic) validateLastModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModifiedDate", "body", "date-time", m.LastModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestDefinitionBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestDefinitionBasic) UnmarshalBinary(b []byte) error {
	var res TestDefinitionBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
