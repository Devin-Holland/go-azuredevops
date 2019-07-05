// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WikiCreateParametersV2 Wiki creation parameters.
// swagger:model WikiCreateParametersV2
type WikiCreateParametersV2 struct {
	WikiCreateBaseParameters

	// Version of the wiki. Not required for ProjectWiki type.
	Version *GitVersionDescriptor `json:"version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WikiCreateParametersV2) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WikiCreateBaseParameters
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WikiCreateBaseParameters = aO0

	// now for regular properties
	var propsWikiCreateParametersV2 struct {
		Version *GitVersionDescriptor `json:"version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsWikiCreateParametersV2); err != nil {
		return err
	}
	m.Version = propsWikiCreateParametersV2.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WikiCreateParametersV2) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WikiCreateBaseParameters)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsWikiCreateParametersV2 struct {
		Version *GitVersionDescriptor `json:"version,omitempty"`
	}
	propsWikiCreateParametersV2.Version = m.Version

	jsonDataPropsWikiCreateParametersV2, errWikiCreateParametersV2 := swag.WriteJSON(propsWikiCreateParametersV2)
	if errWikiCreateParametersV2 != nil {
		return nil, errWikiCreateParametersV2
	}
	_parts = append(_parts, jsonDataPropsWikiCreateParametersV2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this wiki create parameters v2
func (m *WikiCreateParametersV2) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WikiCreateBaseParameters
	if err := m.WikiCreateBaseParameters.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WikiCreateParametersV2) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WikiCreateParametersV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WikiCreateParametersV2) UnmarshalBinary(b []byte) error {
	var res WikiCreateParametersV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}