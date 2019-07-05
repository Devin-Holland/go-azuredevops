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

// BuildDefinitionRevision Represents a revision of a build definition.
// swagger:model BuildDefinitionRevision
type BuildDefinitionRevision struct {

	// The change type (add, edit, delete).
	// Enum: [add update delete]
	ChangeType interface{} `json:"changeType,omitempty"`

	// The identity of the person or process that changed the definition.
	ChangedBy *IdentityRef `json:"changedBy,omitempty"`

	// The date and time that the definition was changed.
	// Format: date-time
	ChangedDate strfmt.DateTime `json:"changedDate,omitempty"`

	// The comment associated with the change.
	Comment string `json:"comment,omitempty"`

	// A link to the definition at this revision.
	DefinitionURL string `json:"definitionUrl,omitempty"`

	// The name of the definition.
	Name string `json:"name,omitempty"`

	// The revision number.
	Revision int32 `json:"revision,omitempty"`
}

// Validate validates this build definition revision
func (m *BuildDefinitionRevision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildDefinitionRevision) validateChangedBy(formats strfmt.Registry) error {

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

func (m *BuildDefinitionRevision) validateChangedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("changedDate", "body", "date-time", m.ChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildDefinitionRevision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildDefinitionRevision) UnmarshalBinary(b []byte) error {
	var res BuildDefinitionRevision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}