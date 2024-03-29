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

// PackageVersionMetrics All metrics for a certain package version id
// swagger:model PackageVersionMetrics
type PackageVersionMetrics struct {

	// Total count of downloads per package version id.
	DownloadCount float64 `json:"downloadCount,omitempty"`

	// Number of downloads per unique user per package version id.
	DownloadUniqueUsers float64 `json:"downloadUniqueUsers,omitempty"`

	// UTC date and time when package version was last downloaded.
	// Format: date-time
	LastDownloaded strfmt.DateTime `json:"lastDownloaded,omitempty"`

	// Package id.
	// Format: uuid
	PackageID strfmt.UUID `json:"packageId,omitempty"`

	// Package version id.
	// Format: uuid
	PackageVersionID strfmt.UUID `json:"packageVersionId,omitempty"`
}

// Validate validates this package version metrics
func (m *PackageVersionMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastDownloaded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageVersionMetrics) validateLastDownloaded(formats strfmt.Registry) error {

	if swag.IsZero(m.LastDownloaded) { // not required
		return nil
	}

	if err := validate.FormatOf("lastDownloaded", "body", "date-time", m.LastDownloaded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PackageVersionMetrics) validatePackageID(formats strfmt.Registry) error {

	if swag.IsZero(m.PackageID) { // not required
		return nil
	}

	if err := validate.FormatOf("packageId", "body", "uuid", m.PackageID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PackageVersionMetrics) validatePackageVersionID(formats strfmt.Registry) error {

	if swag.IsZero(m.PackageVersionID) { // not required
		return nil
	}

	if err := validate.FormatOf("packageVersionId", "body", "uuid", m.PackageVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackageVersionMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageVersionMetrics) UnmarshalBinary(b []byte) error {
	var res PackageVersionMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
