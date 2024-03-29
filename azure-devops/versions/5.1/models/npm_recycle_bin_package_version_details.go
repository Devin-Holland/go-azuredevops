// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NpmRecycleBinPackageVersionDetails npm recycle bin package version details
// swagger:model NpmRecycleBinPackageVersionDetails
type NpmRecycleBinPackageVersionDetails struct {

	// Setting to false will undo earlier deletion and restore the package to feed.
	Deleted bool `json:"deleted,omitempty"`
}

// Validate validates this npm recycle bin package version details
func (m *NpmRecycleBinPackageVersionDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NpmRecycleBinPackageVersionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NpmRecycleBinPackageVersionDetails) UnmarshalBinary(b []byte) error {
	var res NpmRecycleBinPackageVersionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
