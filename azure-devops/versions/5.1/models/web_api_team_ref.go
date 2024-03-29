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

// WebAPITeamRef web Api team ref
// swagger:model WebApiTeamRef
type WebAPITeamRef struct {

	// Team (Identity) Guid. A Team Foundation ID.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Team name
	Name string `json:"name,omitempty"`

	// Team REST API Url
	URL string `json:"url,omitempty"`
}

// Validate validates this web Api team ref
func (m *WebAPITeamRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebAPITeamRef) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebAPITeamRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebAPITeamRef) UnmarshalBinary(b []byte) error {
	var res WebAPITeamRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
