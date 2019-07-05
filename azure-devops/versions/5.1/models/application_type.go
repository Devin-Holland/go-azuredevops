// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ApplicationType application type
// swagger:model ApplicationType
type ApplicationType struct {

	// action Uri link
	ActionURILink string `json:"actionUriLink,omitempty"`

	// aut portal link
	AutPortalLink string `json:"autPortalLink,omitempty"`

	// is enabled
	IsEnabled bool `json:"isEnabled,omitempty"`

	// max components allowed for collection
	MaxComponentsAllowedForCollection int32 `json:"maxComponentsAllowedForCollection,omitempty"`

	// max counters allowed
	MaxCountersAllowed int32 `json:"maxCountersAllowed,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this application type
func (m *ApplicationType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationType) UnmarshalBinary(b []byte) error {
	var res ApplicationType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}