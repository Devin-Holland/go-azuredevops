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

// PermissionEvaluation Represents an evaluated permission.
// swagger:model PermissionEvaluation
type PermissionEvaluation struct {

	// Permission bit for this evaluated permission.
	Permissions int32 `json:"permissions,omitempty"`

	// Security namespace identifier for this evaluated permission.
	// Format: uuid
	SecurityNamespaceID strfmt.UUID `json:"securityNamespaceId,omitempty"`

	// Security namespace-specific token for this evaluated permission.
	Token string `json:"token,omitempty"`

	// Permission evaluation value.
	Value bool `json:"value,omitempty"`
}

// Validate validates this permission evaluation
func (m *PermissionEvaluation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecurityNamespaceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PermissionEvaluation) validateSecurityNamespaceID(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityNamespaceID) { // not required
		return nil
	}

	if err := validate.FormatOf("securityNamespaceId", "body", "uuid", m.SecurityNamespaceID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PermissionEvaluation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionEvaluation) UnmarshalBinary(b []byte) error {
	var res PermissionEvaluation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
