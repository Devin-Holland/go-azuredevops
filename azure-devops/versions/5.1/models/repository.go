// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Repository Defines the details of the repository.
// swagger:model Repository
type Repository struct {

	// Id of the repository.
	ID string `json:"id,omitempty"`

	// Name of the repository.
	Name string `json:"name,omitempty"`

	// Version control type of the result file.
	// Enum: [git tfvc custom]
	Type interface{} `json:"type,omitempty"`
}

// Validate validates this repository
func (m *Repository) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Repository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Repository) UnmarshalBinary(b []byte) error {
	var res Repository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}