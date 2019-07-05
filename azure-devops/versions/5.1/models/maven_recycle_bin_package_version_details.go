// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MavenRecycleBinPackageVersionDetails maven recycle bin package version details
// swagger:model MavenRecycleBinPackageVersionDetails
type MavenRecycleBinPackageVersionDetails struct {

	// Setting to false will undo earlier deletion and restore the package to feed.
	Deleted bool `json:"deleted,omitempty"`
}

// Validate validates this maven recycle bin package version details
func (m *MavenRecycleBinPackageVersionDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MavenRecycleBinPackageVersionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MavenRecycleBinPackageVersionDetails) UnmarshalBinary(b []byte) error {
	var res MavenRecycleBinPackageVersionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
