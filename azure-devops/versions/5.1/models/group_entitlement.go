// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GroupEntitlement A group entity with additional properties including its license, extensions, and project membership
// swagger:model GroupEntitlement
type GroupEntitlement struct {

	// Extension Rules.
	ExtensionRules []*Extension `json:"extensionRules"`

	// Member reference.
	Group *GraphGroup `json:"group,omitempty"`

	// The unique identifier which matches the Id of the GraphMember.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// [Readonly] The last time the group licensing rule was executed (regardless of whether any changes were made).
	// Format: date-time
	LastExecuted strfmt.DateTime `json:"lastExecuted,omitempty"`

	// License Rule.
	LicenseRule *AccessLevel `json:"licenseRule,omitempty"`

	// Group members. Only used when creating a new group.
	Members []*UserEntitlement `json:"members"`

	// Relation between a project and the member's effective permissions in that project.
	ProjectEntitlements []*ProjectEntitlement `json:"projectEntitlements"`

	// The status of the group rule.
	// Enum: [applyPending applied incompatible unableToApply]
	Status interface{} `json:"status,omitempty"`
}

// Validate validates this group entitlement
func (m *GroupEntitlement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensionRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastExecuted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectEntitlements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupEntitlement) validateExtensionRules(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtensionRules) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtensionRules); i++ {
		if swag.IsZero(m.ExtensionRules[i]) { // not required
			continue
		}

		if m.ExtensionRules[i] != nil {
			if err := m.ExtensionRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extensionRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GroupEntitlement) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *GroupEntitlement) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GroupEntitlement) validateLastExecuted(formats strfmt.Registry) error {

	if swag.IsZero(m.LastExecuted) { // not required
		return nil
	}

	if err := validate.FormatOf("lastExecuted", "body", "date-time", m.LastExecuted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GroupEntitlement) validateLicenseRule(formats strfmt.Registry) error {

	if swag.IsZero(m.LicenseRule) { // not required
		return nil
	}

	if m.LicenseRule != nil {
		if err := m.LicenseRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licenseRule")
			}
			return err
		}
	}

	return nil
}

func (m *GroupEntitlement) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GroupEntitlement) validateProjectEntitlements(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectEntitlements) { // not required
		return nil
	}

	for i := 0; i < len(m.ProjectEntitlements); i++ {
		if swag.IsZero(m.ProjectEntitlements[i]) { // not required
			continue
		}

		if m.ProjectEntitlements[i] != nil {
			if err := m.ProjectEntitlements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectEntitlements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupEntitlement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupEntitlement) UnmarshalBinary(b []byte) error {
	var res GroupEntitlement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
