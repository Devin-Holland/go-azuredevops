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

// ReleaseDefinitionEnvironmentTemplate release definition environment template
// swagger:model ReleaseDefinitionEnvironmentTemplate
type ReleaseDefinitionEnvironmentTemplate struct {

	// Indicates whether template can be deleted or not.
	CanDelete bool `json:"canDelete,omitempty"`

	// Category of the ReleaseDefinition environment template.
	Category string `json:"category,omitempty"`

	// Description of the ReleaseDefinition environment template.
	Description string `json:"description,omitempty"`

	// ReleaseDefinition environment data which used to create this template.
	Environment *ReleaseDefinitionEnvironment `json:"environment,omitempty"`

	// ID of the task which used to display icon used for this template.
	// Format: uuid
	IconTaskID strfmt.UUID `json:"iconTaskId,omitempty"`

	// Icon uri of the template.
	IconURI string `json:"iconUri,omitempty"`

	// ID of the ReleaseDefinition environment template.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Indicates whether template deleted or not.
	IsDeleted bool `json:"isDeleted,omitempty"`

	// Name of the ReleaseDefinition environment template.
	Name string `json:"name,omitempty"`
}

// Validate validates this release definition environment template
func (m *ReleaseDefinitionEnvironmentTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIconTaskID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseDefinitionEnvironmentTemplate) validateEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseDefinitionEnvironmentTemplate) validateIconTaskID(formats strfmt.Registry) error {

	if swag.IsZero(m.IconTaskID) { // not required
		return nil
	}

	if err := validate.FormatOf("iconTaskId", "body", "uuid", m.IconTaskID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReleaseDefinitionEnvironmentTemplate) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseDefinitionEnvironmentTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseDefinitionEnvironmentTemplate) UnmarshalBinary(b []byte) error {
	var res ReleaseDefinitionEnvironmentTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}