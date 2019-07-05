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

// CloneOperationCommonResponse Common Response for clone operation
// swagger:model CloneOperationCommonResponse
type CloneOperationCommonResponse struct {

	// Various statistics related to the clone operation
	CloneStatistics *CloneStatistics `json:"cloneStatistics,omitempty"`

	// Completion data of the operation
	// Format: date-time
	CompletionDate strfmt.DateTime `json:"completionDate,omitempty"`

	// Creation data of the operation
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// Reference links
	Links *ReferenceLinks `json:"links,omitempty"`

	// Message related to the job
	Message string `json:"message,omitempty"`

	// Clone operation Id
	OpID int32 `json:"opId,omitempty"`

	// Clone operation state
	// Enum: [failed inProgress queued succeeded]
	State interface{} `json:"state,omitempty"`
}

// Validate validates this clone operation common response
func (m *CloneOperationCommonResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloneStatistics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloneOperationCommonResponse) validateCloneStatistics(formats strfmt.Registry) error {

	if swag.IsZero(m.CloneStatistics) { // not required
		return nil
	}

	if m.CloneStatistics != nil {
		if err := m.CloneStatistics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneStatistics")
			}
			return err
		}
	}

	return nil
}

func (m *CloneOperationCommonResponse) validateCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("completionDate", "body", "date-time", m.CompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloneOperationCommonResponse) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloneOperationCommonResponse) validateLinks(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CloneOperationCommonResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloneOperationCommonResponse) UnmarshalBinary(b []byte) error {
	var res CloneOperationCommonResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}