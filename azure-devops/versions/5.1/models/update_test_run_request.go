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

// UpdateTestRunRequest update test run request
// swagger:model UpdateTestRunRequest
type UpdateTestRunRequest struct {

	// attachments to add
	AttachmentsToAdd []*TestResultAttachment `json:"attachmentsToAdd"`

	// attachments to delete
	AttachmentsToDelete []*TestResultAttachmentIdentity `json:"attachmentsToDelete"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// should hyderate
	ShouldHyderate bool `json:"shouldHyderate,omitempty"`

	// test run
	TestRun *LegacyTestRun `json:"testRun,omitempty"`
}

// Validate validates this update test run request
func (m *UpdateTestRunRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentsToAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttachmentsToDelete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTestRunRequest) validateAttachmentsToAdd(formats strfmt.Registry) error {

	if swag.IsZero(m.AttachmentsToAdd) { // not required
		return nil
	}

	for i := 0; i < len(m.AttachmentsToAdd); i++ {
		if swag.IsZero(m.AttachmentsToAdd[i]) { // not required
			continue
		}

		if m.AttachmentsToAdd[i] != nil {
			if err := m.AttachmentsToAdd[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachmentsToAdd" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateTestRunRequest) validateAttachmentsToDelete(formats strfmt.Registry) error {

	if swag.IsZero(m.AttachmentsToDelete) { // not required
		return nil
	}

	for i := 0; i < len(m.AttachmentsToDelete); i++ {
		if swag.IsZero(m.AttachmentsToDelete[i]) { // not required
			continue
		}

		if m.AttachmentsToDelete[i] != nil {
			if err := m.AttachmentsToDelete[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachmentsToDelete" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateTestRunRequest) validateTestRun(formats strfmt.Registry) error {

	if swag.IsZero(m.TestRun) { // not required
		return nil
	}

	if m.TestRun != nil {
		if err := m.TestRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testRun")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTestRunRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTestRunRequest) UnmarshalBinary(b []byte) error {
	var res UpdateTestRunRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
