// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MemberEntitlement Deprecated: Use UserEntitlement instead
// swagger:model MemberEntitlement
type MemberEntitlement struct {
	UserEntitlement

	// Member reference
	Member *GraphMember `json:"member,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MemberEntitlement) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 UserEntitlement
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.UserEntitlement = aO0

	// now for regular properties
	var propsMemberEntitlement struct {
		Member *GraphMember `json:"member,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsMemberEntitlement); err != nil {
		return err
	}
	m.Member = propsMemberEntitlement.Member

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MemberEntitlement) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.UserEntitlement)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsMemberEntitlement struct {
		Member *GraphMember `json:"member,omitempty"`
	}
	propsMemberEntitlement.Member = m.Member

	jsonDataPropsMemberEntitlement, errMemberEntitlement := swag.WriteJSON(propsMemberEntitlement)
	if errMemberEntitlement != nil {
		return nil, errMemberEntitlement
	}
	_parts = append(_parts, jsonDataPropsMemberEntitlement)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this member entitlement
func (m *MemberEntitlement) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UserEntitlement
	if err := m.UserEntitlement.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMember(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemberEntitlement) validateMember(formats strfmt.Registry) error {

	if swag.IsZero(m.Member) { // not required
		return nil
	}

	if m.Member != nil {
		if err := m.Member.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemberEntitlement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberEntitlement) UnmarshalBinary(b []byte) error {
	var res MemberEntitlement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
