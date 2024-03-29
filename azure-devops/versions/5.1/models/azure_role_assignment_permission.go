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

// AzureRoleAssignmentPermission azure role assignment permission
// swagger:model AzureRoleAssignmentPermission
type AzureRoleAssignmentPermission struct {
	AzurePermission

	// role assignment Id
	// Format: uuid
	RoleAssignmentID strfmt.UUID `json:"roleAssignmentId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AzureRoleAssignmentPermission) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AzurePermission
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AzurePermission = aO0

	// now for regular properties
	var propsAzureRoleAssignmentPermission struct {
		RoleAssignmentID strfmt.UUID `json:"roleAssignmentId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsAzureRoleAssignmentPermission); err != nil {
		return err
	}
	m.RoleAssignmentID = propsAzureRoleAssignmentPermission.RoleAssignmentID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AzureRoleAssignmentPermission) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.AzurePermission)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsAzureRoleAssignmentPermission struct {
		RoleAssignmentID strfmt.UUID `json:"roleAssignmentId,omitempty"`
	}
	propsAzureRoleAssignmentPermission.RoleAssignmentID = m.RoleAssignmentID

	jsonDataPropsAzureRoleAssignmentPermission, errAzureRoleAssignmentPermission := swag.WriteJSON(propsAzureRoleAssignmentPermission)
	if errAzureRoleAssignmentPermission != nil {
		return nil, errAzureRoleAssignmentPermission
	}
	_parts = append(_parts, jsonDataPropsAzureRoleAssignmentPermission)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this azure role assignment permission
func (m *AzureRoleAssignmentPermission) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AzurePermission
	if err := m.AzurePermission.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleAssignmentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureRoleAssignmentPermission) validateRoleAssignmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleAssignmentID) { // not required
		return nil
	}

	if err := validate.FormatOf("roleAssignmentId", "body", "uuid", m.RoleAssignmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureRoleAssignmentPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureRoleAssignmentPermission) UnmarshalBinary(b []byte) error {
	var res AzureRoleAssignmentPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
