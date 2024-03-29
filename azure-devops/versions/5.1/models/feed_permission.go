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

// FeedPermission Permissions for a feed.
// swagger:model FeedPermission
type FeedPermission struct {

	// Display name for the identity.
	DisplayName string `json:"displayName,omitempty"`

	// Identity associated with this role.
	IdentityDescriptor *IdentityDescriptor `json:"identityDescriptor,omitempty"`

	// Id of the identity associated with this role.
	// Format: uuid
	IdentityID strfmt.UUID `json:"identityId,omitempty"`

	// The role for this identity on a feed.
	// Enum: [custom none reader contributor administrator collaborator]
	Role interface{} `json:"role,omitempty"`
}

// Validate validates this feed permission
func (m *FeedPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentityDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeedPermission) validateIdentityDescriptor(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentityDescriptor) { // not required
		return nil
	}

	if m.IdentityDescriptor != nil {
		if err := m.IdentityDescriptor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identityDescriptor")
			}
			return err
		}
	}

	return nil
}

func (m *FeedPermission) validateIdentityID(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentityID) { // not required
		return nil
	}

	if err := validate.FormatOf("identityId", "body", "uuid", m.IdentityID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeedPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedPermission) UnmarshalBinary(b []byte) error {
	var res FeedPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
