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

// EmailRecipients email recipients
// swagger:model EmailRecipients
type EmailRecipients struct {

	// List of email addresses.
	EmailAddresses []string `json:"emailAddresses"`

	// List of TFS IDs guids.
	TfsIds []strfmt.UUID `json:"tfsIds"`
}

// Validate validates this email recipients
func (m *EmailRecipients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTfsIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailRecipients) validateTfsIds(formats strfmt.Registry) error {

	if swag.IsZero(m.TfsIds) { // not required
		return nil
	}

	for i := 0; i < len(m.TfsIds); i++ {

		if err := validate.FormatOf("tfsIds"+"."+strconv.Itoa(i), "body", "uuid", m.TfsIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmailRecipients) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailRecipients) UnmarshalBinary(b []byte) error {
	var res EmailRecipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
