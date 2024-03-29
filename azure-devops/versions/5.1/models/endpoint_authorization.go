// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EndpointAuthorization endpoint authorization
// swagger:model EndpointAuthorization
type EndpointAuthorization struct {

	// Gets or sets the parameters for the selected authorization scheme.
	Parameters map[string]string `json:"parameters,omitempty"`

	// Gets or sets the scheme used for service endpoint authentication.
	Scheme string `json:"scheme,omitempty"`
}

// Validate validates this endpoint authorization
func (m *EndpointAuthorization) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointAuthorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointAuthorization) UnmarshalBinary(b []byte) error {
	var res EndpointAuthorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
