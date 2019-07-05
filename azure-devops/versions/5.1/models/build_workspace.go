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

// BuildWorkspace Represents a workspace mapping.
// swagger:model BuildWorkspace
type BuildWorkspace struct {

	// mappings
	Mappings []*MappingDetails `json:"mappings"`
}

// Validate validates this build workspace
func (m *BuildWorkspace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMappings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildWorkspace) validateMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.Mappings) { // not required
		return nil
	}

	for i := 0; i < len(m.Mappings); i++ {
		if swag.IsZero(m.Mappings[i]) { // not required
			continue
		}

		if m.Mappings[i] != nil {
			if err := m.Mappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildWorkspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildWorkspace) UnmarshalBinary(b []byte) error {
	var res BuildWorkspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
