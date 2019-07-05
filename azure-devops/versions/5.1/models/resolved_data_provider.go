// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ResolvedDataProvider Entry for a specific data provider's resulting data
// swagger:model ResolvedDataProvider
type ResolvedDataProvider struct {

	// The total time the data provider took to resolve its data (in milliseconds)
	Duration string `json:"duration,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this resolved data provider
func (m *ResolvedDataProvider) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResolvedDataProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResolvedDataProvider) UnmarshalBinary(b []byte) error {
	var res ResolvedDataProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}