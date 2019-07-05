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

// ContributionType A contribution type, given by a json schema
// swagger:model ContributionType
type ContributionType struct {
	ContributionBase

	// Controls whether or not contributions of this type have the type indexed for queries. This allows clients to find all extensions that have a contribution of this type.  NOTE: Only TrustedPartners are allowed to specify indexed contribution types.
	Indexed bool `json:"indexed,omitempty"`

	// Friendly name of the contribution/type
	Name string `json:"name,omitempty"`

	// Describes the allowed properties for this contribution type
	Properties map[string]ContributionPropertyDescription `json:"properties,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ContributionType) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ContributionBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ContributionBase = aO0

	// now for regular properties
	var propsContributionType struct {
		Indexed bool `json:"indexed,omitempty"`

		Name string `json:"name,omitempty"`

		Properties map[string]ContributionPropertyDescription `json:"properties,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsContributionType); err != nil {
		return err
	}
	m.Indexed = propsContributionType.Indexed

	m.Name = propsContributionType.Name

	m.Properties = propsContributionType.Properties

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ContributionType) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ContributionBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsContributionType struct {
		Indexed bool `json:"indexed,omitempty"`

		Name string `json:"name,omitempty"`

		Properties map[string]ContributionPropertyDescription `json:"properties,omitempty"`
	}
	propsContributionType.Indexed = m.Indexed

	propsContributionType.Name = m.Name

	propsContributionType.Properties = m.Properties

	jsonDataPropsContributionType, errContributionType := swag.WriteJSON(propsContributionType)
	if errContributionType != nil {
		return nil, errContributionType
	}
	_parts = append(_parts, jsonDataPropsContributionType)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this contribution type
func (m *ContributionType) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ContributionBase
	if err := m.ContributionBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContributionType) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for k := range m.Properties {

		if err := validate.Required("properties"+"."+k, "body", m.Properties[k]); err != nil {
			return err
		}
		if val, ok := m.Properties[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContributionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContributionType) UnmarshalBinary(b []byte) error {
	var res ContributionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
