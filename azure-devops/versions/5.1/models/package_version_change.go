// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PackageVersionChange A change to a single package version.
// swagger:model PackageVersionChange
type PackageVersionChange struct {

	// The type of change that was performed.
	// Enum: [addOrUpdate delete]
	ChangeType interface{} `json:"changeType,omitempty"`

	// Token marker for this change, allowing the caller to send this value back to the service and receive changes beyond this one.
	ContinuationToken int64 `json:"continuationToken,omitempty"`

	// Package version that was changed.
	PackageVersion *PackageVersion `json:"packageVersion,omitempty"`
}

// Validate validates this package version change
func (m *PackageVersionChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackageVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageVersionChange) validatePackageVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.PackageVersion) { // not required
		return nil
	}

	if m.PackageVersion != nil {
		if err := m.PackageVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packageVersion")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackageVersionChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageVersionChange) UnmarshalBinary(b []byte) error {
	var res PackageVersionChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
