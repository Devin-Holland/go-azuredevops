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

// NuGetPackageVersionDeletionState Deletion state of a NuGet package.
// swagger:model NuGetPackageVersionDeletionState
type NuGetPackageVersionDeletionState struct {

	// Utc date the package was deleted.
	// Format: date-time
	DeletedDate strfmt.DateTime `json:"deletedDate,omitempty"`

	// Name of the package.
	Name string `json:"name,omitempty"`

	// Version of the package.
	Version string `json:"version,omitempty"`
}

// Validate validates this nu get package version deletion state
func (m *NuGetPackageVersionDeletionState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeletedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NuGetPackageVersionDeletionState) validateDeletedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("deletedDate", "body", "date-time", m.DeletedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NuGetPackageVersionDeletionState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NuGetPackageVersionDeletionState) UnmarshalBinary(b []byte) error {
	var res NuGetPackageVersionDeletionState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
