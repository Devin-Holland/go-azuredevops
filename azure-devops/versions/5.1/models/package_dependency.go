// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PackageDependency A dependency on another package version.
// swagger:model PackageDependency
type PackageDependency struct {

	// Dependency package group (an optional classification within some package types).
	Group string `json:"group,omitempty"`

	// Dependency package name.
	PackageName string `json:"packageName,omitempty"`

	// Dependency package version range.
	VersionRange string `json:"versionRange,omitempty"`
}

// Validate validates this package dependency
func (m *PackageDependency) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackageDependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageDependency) UnmarshalBinary(b []byte) error {
	var res PackageDependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
