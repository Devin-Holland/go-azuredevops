// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MinimalPackageVersion Core data about any package, including its id and version information and basic state.
// swagger:model MinimalPackageVersion
type MinimalPackageVersion struct {

	// Upstream source this package was ingested from.
	// Format: uuid
	DirectUpstreamSourceID strfmt.UUID `json:"directUpstreamSourceId,omitempty"`

	// Id for the package.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// [Obsolete] Used for legacy scenarios and may be removed in future versions.
	IsCachedVersion bool `json:"isCachedVersion,omitempty"`

	// True if this package has been deleted.
	IsDeleted bool `json:"isDeleted,omitempty"`

	// True if this is the latest version of the package by package type sort order.
	IsLatest bool `json:"isLatest,omitempty"`

	// (NuGet Only) True if this package is listed.
	IsListed bool `json:"isListed,omitempty"`

	// Normalized version using normalization rules specific to a package type.
	NormalizedVersion string `json:"normalizedVersion,omitempty"`

	// Package description.
	PackageDescription string `json:"packageDescription,omitempty"`

	// UTC Date the package was published to the service.
	// Format: date-time
	PublishDate strfmt.DateTime `json:"publishDate,omitempty"`

	// Internal storage id.
	StorageID string `json:"storageId,omitempty"`

	// Display version.
	Version string `json:"version,omitempty"`

	// List of views containing this package version.
	Views []*FeedView `json:"views"`
}

// Validate validates this minimal package version
func (m *MinimalPackageVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectUpstreamSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViews(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MinimalPackageVersion) validateDirectUpstreamSourceID(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectUpstreamSourceID) { // not required
		return nil
	}

	if err := validate.FormatOf("directUpstreamSourceId", "body", "uuid", m.DirectUpstreamSourceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MinimalPackageVersion) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MinimalPackageVersion) validatePublishDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishDate) { // not required
		return nil
	}

	if err := validate.FormatOf("publishDate", "body", "date-time", m.PublishDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MinimalPackageVersion) validateViews(formats strfmt.Registry) error {

	if swag.IsZero(m.Views) { // not required
		return nil
	}

	for i := 0; i < len(m.Views); i++ {
		if swag.IsZero(m.Views[i]) { // not required
			continue
		}

		if m.Views[i] != nil {
			if err := m.Views[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("views" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MinimalPackageVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MinimalPackageVersion) UnmarshalBinary(b []byte) error {
	var res MinimalPackageVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
