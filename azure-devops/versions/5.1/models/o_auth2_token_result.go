// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OAuth2TokenResult o auth2 token result
// swagger:model OAuth2TokenResult
type OAuth2TokenResult struct {

	// access token
	AccessToken string `json:"accessToken,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// error description
	ErrorDescription string `json:"errorDescription,omitempty"`

	// expires in
	ExpiresIn string `json:"expiresIn,omitempty"`

	// issued at
	IssuedAt string `json:"issuedAt,omitempty"`

	// refresh token
	RefreshToken string `json:"refreshToken,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`
}

// Validate validates this o auth2 token result
func (m *OAuth2TokenResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OAuth2TokenResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuth2TokenResult) UnmarshalBinary(b []byte) error {
	var res OAuth2TokenResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
