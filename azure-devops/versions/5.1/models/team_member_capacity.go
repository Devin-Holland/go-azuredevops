// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TeamMemberCapacity Represents capacity for a specific team member
// swagger:model TeamMemberCapacity
type TeamMemberCapacity struct {
	CapacityContractBase

	// Shallow Ref to the associated team member
	TeamMember *Member `json:"teamMember,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TeamMemberCapacity) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CapacityContractBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CapacityContractBase = aO0

	// now for regular properties
	var propsTeamMemberCapacity struct {
		TeamMember *Member `json:"teamMember,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsTeamMemberCapacity); err != nil {
		return err
	}
	m.TeamMember = propsTeamMemberCapacity.TeamMember

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TeamMemberCapacity) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.CapacityContractBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTeamMemberCapacity struct {
		TeamMember *Member `json:"teamMember,omitempty"`
	}
	propsTeamMemberCapacity.TeamMember = m.TeamMember

	jsonDataPropsTeamMemberCapacity, errTeamMemberCapacity := swag.WriteJSON(propsTeamMemberCapacity)
	if errTeamMemberCapacity != nil {
		return nil, errTeamMemberCapacity
	}
	_parts = append(_parts, jsonDataPropsTeamMemberCapacity)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this team member capacity
func (m *TeamMemberCapacity) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CapacityContractBase
	if err := m.CapacityContractBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamMember(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamMemberCapacity) validateTeamMember(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamMember) { // not required
		return nil
	}

	if m.TeamMember != nil {
		if err := m.TeamMember.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamMember")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamMemberCapacity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamMemberCapacity) UnmarshalBinary(b []byte) error {
	var res TeamMemberCapacity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
