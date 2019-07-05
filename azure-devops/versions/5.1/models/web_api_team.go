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

// WebAPITeam web Api team
// swagger:model WebApiTeam
type WebAPITeam struct {
	WebAPITeamRef

	// Team description
	Description string `json:"description,omitempty"`

	// Team identity.
	Identity *Identity `json:"identity,omitempty"`

	// Identity REST API Url to this team
	IdentityURL string `json:"identityUrl,omitempty"`

	// project Id
	// Format: uuid
	ProjectID strfmt.UUID `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WebAPITeam) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WebAPITeamRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WebAPITeamRef = aO0

	// now for regular properties
	var propsWebAPITeam struct {
		Description string `json:"description,omitempty"`

		Identity *Identity `json:"identity,omitempty"`

		IdentityURL string `json:"identityUrl,omitempty"`

		ProjectID strfmt.UUID `json:"projectId,omitempty"`

		ProjectName string `json:"projectName,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsWebAPITeam); err != nil {
		return err
	}
	m.Description = propsWebAPITeam.Description

	m.Identity = propsWebAPITeam.Identity

	m.IdentityURL = propsWebAPITeam.IdentityURL

	m.ProjectID = propsWebAPITeam.ProjectID

	m.ProjectName = propsWebAPITeam.ProjectName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WebAPITeam) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WebAPITeamRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsWebAPITeam struct {
		Description string `json:"description,omitempty"`

		Identity *Identity `json:"identity,omitempty"`

		IdentityURL string `json:"identityUrl,omitempty"`

		ProjectID strfmt.UUID `json:"projectId,omitempty"`

		ProjectName string `json:"projectName,omitempty"`
	}
	propsWebAPITeam.Description = m.Description

	propsWebAPITeam.Identity = m.Identity

	propsWebAPITeam.IdentityURL = m.IdentityURL

	propsWebAPITeam.ProjectID = m.ProjectID

	propsWebAPITeam.ProjectName = m.ProjectName

	jsonDataPropsWebAPITeam, errWebAPITeam := swag.WriteJSON(propsWebAPITeam)
	if errWebAPITeam != nil {
		return nil, errWebAPITeam
	}
	_parts = append(_parts, jsonDataPropsWebAPITeam)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this web Api team
func (m *WebAPITeam) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WebAPITeamRef
	if err := m.WebAPITeamRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebAPITeam) validateIdentity(formats strfmt.Registry) error {

	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {
		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *WebAPITeam) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectId", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebAPITeam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebAPITeam) UnmarshalBinary(b []byte) error {
	var res WebAPITeam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
