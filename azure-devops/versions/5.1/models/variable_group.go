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

// VariableGroup variable group
// swagger:model VariableGroup
type VariableGroup struct {

	// Gets or sets the identity who created.
	CreatedBy *IdentityRef `json:"createdBy,omitempty"`

	// Gets date on which it got created.
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

	// Gets or sets description.
	Description string `json:"description,omitempty"`

	// Gets the unique identifier of this field.
	ID int32 `json:"id,omitempty"`

	// Denotes if a variable group is shared with other project or not.
	IsShared bool `json:"isShared,omitempty"`

	// Gets or sets the identity who modified.
	ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`

	// Gets date on which it got modified.
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modifiedOn,omitempty"`

	// Gets or sets name.
	Name string `json:"name,omitempty"`

	// Gets or sets provider data.
	ProviderData VariableGroupProviderData `json:"providerData,omitempty"`

	// Gets or sets type.
	Type string `json:"type,omitempty"`

	// Gets and sets the dictionary of variables.
	Variables map[string]VariableValue `json:"variables,omitempty"`
}

// Validate validates this variable group
func (m *VariableGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariableGroup) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *VariableGroup) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("createdOn", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VariableGroup) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *VariableGroup) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedOn", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VariableGroup) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for k := range m.Variables {

		if err := validate.Required("variables"+"."+k, "body", m.Variables[k]); err != nil {
			return err
		}
		if val, ok := m.Variables[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VariableGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariableGroup) UnmarshalBinary(b []byte) error {
	var res VariableGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
