// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GraphMembership Relationship between a container and a member
// swagger:model GraphMembership
type GraphMembership struct {

	// This field contains zero or more interesting links about the graph membership. These links may be invoked to obtain additional relationships or more detailed information about this graph membership.
	Links *ReferenceLinks `json:"_links,omitempty"`

	// container descriptor
	ContainerDescriptor string `json:"containerDescriptor,omitempty"`

	// member descriptor
	MemberDescriptor string `json:"memberDescriptor,omitempty"`
}

// Validate validates this graph membership
func (m *GraphMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphMembership) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GraphMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphMembership) UnmarshalBinary(b []byte) error {
	var res GraphMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}