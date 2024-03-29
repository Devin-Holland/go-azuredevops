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

// TaskDefinitionReference task definition reference
// swagger:model TaskDefinitionReference
type TaskDefinitionReference struct {

	// Gets or sets the definition type. Values can be 'task' or 'metaTask'.
	DefinitionType string `json:"definitionType,omitempty"`

	// Gets or sets the unique identifier of task.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Gets or sets the version specification of task.
	VersionSpec string `json:"versionSpec,omitempty"`
}

// Validate validates this task definition reference
func (m *TaskDefinitionReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskDefinitionReference) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskDefinitionReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskDefinitionReference) UnmarshalBinary(b []byte) error {
	var res TaskDefinitionReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
