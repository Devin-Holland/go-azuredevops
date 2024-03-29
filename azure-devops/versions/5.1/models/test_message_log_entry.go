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

// TestMessageLogEntry test message log entry
// swagger:model TestMessageLogEntry
type TestMessageLogEntry struct {

	// date created
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// entry Id
	EntryID int32 `json:"entryId,omitempty"`

	// log level
	// Format: byte
	LogLevel strfmt.Base64 `json:"logLevel,omitempty"`

	// log user
	// Format: uuid
	LogUser strfmt.UUID `json:"logUser,omitempty"`

	// log user name
	LogUserName string `json:"logUserName,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// test message log Id
	TestMessageLogID int32 `json:"testMessageLogId,omitempty"`
}

// Validate validates this test message log entry
func (m *TestMessageLogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestMessageLogEntry) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestMessageLogEntry) validateLogLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.LogLevel) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *TestMessageLogEntry) validateLogUser(formats strfmt.Registry) error {

	if swag.IsZero(m.LogUser) { // not required
		return nil
	}

	if err := validate.FormatOf("logUser", "body", "uuid", m.LogUser.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestMessageLogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestMessageLogEntry) UnmarshalBinary(b []byte) error {
	var res TestMessageLogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
