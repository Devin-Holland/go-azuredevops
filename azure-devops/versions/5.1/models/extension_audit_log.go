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

// ExtensionAuditLog Audit log for an extension
// swagger:model ExtensionAuditLog
type ExtensionAuditLog struct {

	// Collection of audit log entries
	Entries []*ExtensionAuditLogEntry `json:"entries"`

	// Extension that the change was made for
	ExtensionName string `json:"extensionName,omitempty"`

	// Publisher that the extension is part of
	PublisherName string `json:"publisherName,omitempty"`
}

// Validate validates this extension audit log
func (m *ExtensionAuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionAuditLog) validateEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.Entries) { // not required
		return nil
	}

	for i := 0; i < len(m.Entries); i++ {
		if swag.IsZero(m.Entries[i]) { // not required
			continue
		}

		if m.Entries[i] != nil {
			if err := m.Entries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionAuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionAuditLog) UnmarshalBinary(b []byte) error {
	var res ExtensionAuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
