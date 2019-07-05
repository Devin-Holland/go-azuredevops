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

// ExtensionAuditLogEntry An audit log entry for an extension
// swagger:model ExtensionAuditLogEntry
type ExtensionAuditLogEntry struct {

	// Change that was made to extension
	AuditAction string `json:"auditAction,omitempty"`

	// Date at which the change was made
	// Format: date-time
	AuditDate strfmt.DateTime `json:"auditDate,omitempty"`

	// Extra information about the change
	Comment string `json:"comment,omitempty"`

	// Represents the user who made the change
	UpdatedBy *IdentityRef `json:"updatedBy,omitempty"`
}

// Validate validates this extension audit log entry
func (m *ExtensionAuditLogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionAuditLogEntry) validateAuditDate(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditDate) { // not required
		return nil
	}

	if err := validate.FormatOf("auditDate", "body", "date-time", m.AuditDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExtensionAuditLogEntry) validateUpdatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedBy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionAuditLogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionAuditLogEntry) UnmarshalBinary(b []byte) error {
	var res ExtensionAuditLogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
