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

// PublishTaskGroupMetadata publish task group metadata
// swagger:model PublishTaskGroupMetadata
type PublishTaskGroupMetadata struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// parent definition revision
	ParentDefinitionRevision int32 `json:"parentDefinitionRevision,omitempty"`

	// preview
	Preview bool `json:"preview,omitempty"`

	// task group Id
	// Format: uuid
	TaskGroupID strfmt.UUID `json:"taskGroupId,omitempty"`

	// task group revision
	TaskGroupRevision int32 `json:"taskGroupRevision,omitempty"`
}

// Validate validates this publish task group metadata
func (m *PublishTaskGroupMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaskGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublishTaskGroupMetadata) validateTaskGroupID(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskGroupID) { // not required
		return nil
	}

	if err := validate.FormatOf("taskGroupId", "body", "uuid", m.TaskGroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublishTaskGroupMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublishTaskGroupMetadata) UnmarshalBinary(b []byte) error {
	var res PublishTaskGroupMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
