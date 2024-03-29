// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountRecentActivityWorkItemModel2 Represents Work Item Recent Activity
// swagger:model AccountRecentActivityWorkItemModel2
type AccountRecentActivityWorkItemModel2 struct {
	AccountRecentActivityWorkItemModelBase

	// Assigned To
	AssignedTo *IdentityRef `json:"assignedTo,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AccountRecentActivityWorkItemModel2) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AccountRecentActivityWorkItemModelBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AccountRecentActivityWorkItemModelBase = aO0

	// now for regular properties
	var propsAccountRecentActivityWorkItemModel2 struct {
		AssignedTo *IdentityRef `json:"assignedTo,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsAccountRecentActivityWorkItemModel2); err != nil {
		return err
	}
	m.AssignedTo = propsAccountRecentActivityWorkItemModel2.AssignedTo

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AccountRecentActivityWorkItemModel2) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.AccountRecentActivityWorkItemModelBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsAccountRecentActivityWorkItemModel2 struct {
		AssignedTo *IdentityRef `json:"assignedTo,omitempty"`
	}
	propsAccountRecentActivityWorkItemModel2.AssignedTo = m.AssignedTo

	jsonDataPropsAccountRecentActivityWorkItemModel2, errAccountRecentActivityWorkItemModel2 := swag.WriteJSON(propsAccountRecentActivityWorkItemModel2)
	if errAccountRecentActivityWorkItemModel2 != nil {
		return nil, errAccountRecentActivityWorkItemModel2
	}
	_parts = append(_parts, jsonDataPropsAccountRecentActivityWorkItemModel2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this account recent activity work item model2
func (m *AccountRecentActivityWorkItemModel2) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AccountRecentActivityWorkItemModelBase
	if err := m.AccountRecentActivityWorkItemModelBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRecentActivityWorkItemModel2) validateAssignedTo(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignedTo) { // not required
		return nil
	}

	if m.AssignedTo != nil {
		if err := m.AssignedTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignedTo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRecentActivityWorkItemModel2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRecentActivityWorkItemModel2) UnmarshalBinary(b []byte) error {
	var res AccountRecentActivityWorkItemModel2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
