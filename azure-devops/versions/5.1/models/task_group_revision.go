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

// TaskGroupRevision task group revision
// swagger:model TaskGroupRevision
type TaskGroupRevision struct {

	// change type
	// Enum: [add update delete undelete]
	ChangeType interface{} `json:"changeType,omitempty"`

	// changed by
	ChangedBy *IdentityRef `json:"changedBy,omitempty"`

	// changed date
	// Format: date-time
	ChangedDate strfmt.DateTime `json:"changedDate,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// file Id
	FileID int32 `json:"fileId,omitempty"`

	// major version
	MajorVersion int32 `json:"majorVersion,omitempty"`

	// revision
	Revision int32 `json:"revision,omitempty"`

	// task group Id
	// Format: uuid
	TaskGroupID strfmt.UUID `json:"taskGroupId,omitempty"`
}

// Validate validates this task group revision
func (m *TaskGroupRevision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskGroupRevision) validateChangedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangedBy) { // not required
		return nil
	}

	if m.ChangedBy != nil {
		if err := m.ChangedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TaskGroupRevision) validateChangedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("changedDate", "body", "date-time", m.ChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskGroupRevision) validateTaskGroupID(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskGroupID) { // not required
		return nil
	}

	if err := validate.FormatOf("taskGroupId", "body", "uuid", m.TaskGroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskGroupRevision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskGroupRevision) UnmarshalBinary(b []byte) error {
	var res TaskGroupRevision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}