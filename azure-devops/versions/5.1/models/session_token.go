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

// SessionToken session token
// swagger:model SessionToken
type SessionToken struct {

	// access Id
	// Format: uuid
	AccessID strfmt.UUID `json:"accessId,omitempty"`

	// This is populated when user requests a compact token. The alternate token value is self describing token.
	AlternateToken string `json:"alternateToken,omitempty"`

	// authorization Id
	// Format: uuid
	AuthorizationID strfmt.UUID `json:"authorizationId,omitempty"`

	// claims
	Claims map[string]string `json:"claims,omitempty"`

	// client Id
	// Format: uuid
	ClientID strfmt.UUID `json:"clientId,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// host authorization Id
	// Format: uuid
	HostAuthorizationID strfmt.UUID `json:"hostAuthorizationId,omitempty"`

	// is public
	IsPublic bool `json:"isPublic,omitempty"`

	// is valid
	IsValid bool `json:"isValid,omitempty"`

	// public data
	PublicData string `json:"publicData,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// target accounts
	TargetAccounts []strfmt.UUID `json:"targetAccounts"`

	// This is computed and not returned in Get queries
	Token string `json:"token,omitempty"`

	// user Id
	// Format: uuid
	UserID strfmt.UUID `json:"userId,omitempty"`

	// valid from
	// Format: date-time
	ValidFrom strfmt.DateTime `json:"validFrom,omitempty"`

	// valid to
	// Format: date-time
	ValidTo strfmt.DateTime `json:"validTo,omitempty"`
}

// Validate validates this session token
func (m *SessionToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostAuthorizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SessionToken) validateAccessID(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessID) { // not required
		return nil
	}

	if err := validate.FormatOf("accessId", "body", "uuid", m.AccessID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SessionToken) validateAuthorizationID(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthorizationID) { // not required
		return nil
	}

	if err := validate.FormatOf("authorizationId", "body", "uuid", m.AuthorizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SessionToken) validateClientID(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientID) { // not required
		return nil
	}

	if err := validate.FormatOf("clientId", "body", "uuid", m.ClientID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SessionToken) validateHostAuthorizationID(formats strfmt.Registry) error {

	if swag.IsZero(m.HostAuthorizationID) { // not required
		return nil
	}

	if err := validate.FormatOf("hostAuthorizationId", "body", "uuid", m.HostAuthorizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SessionToken) validateTargetAccounts(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetAccounts) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetAccounts); i++ {

		if err := validate.FormatOf("targetAccounts"+"."+strconv.Itoa(i), "body", "uuid", m.TargetAccounts[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *SessionToken) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	if err := validate.FormatOf("userId", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SessionToken) validateValidFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("validFrom", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SessionToken) validateValidTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidTo) { // not required
		return nil
	}

	if err := validate.FormatOf("validTo", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SessionToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SessionToken) UnmarshalBinary(b []byte) error {
	var res SessionToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
