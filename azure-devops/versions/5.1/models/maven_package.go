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

// MavenPackage maven package
// swagger:model MavenPackage
type MavenPackage struct {

	// artifact Id
	ArtifactID string `json:"artifactId,omitempty"`

	// artifact index
	ArtifactIndex *ReferenceLink `json:"artifactIndex,omitempty"`

	// artifact metadata
	ArtifactMetadata *ReferenceLink `json:"artifactMetadata,omitempty"`

	// deleted date
	// Format: date-time
	DeletedDate strfmt.DateTime `json:"deletedDate,omitempty"`

	// files
	Files *ReferenceLinks `json:"files,omitempty"`

	// group Id
	GroupID string `json:"groupId,omitempty"`

	// pom
	Pom *MavenPomMetadata `json:"pom,omitempty"`

	// requested file
	RequestedFile *ReferenceLink `json:"requestedFile,omitempty"`

	// snapshot metadata
	SnapshotMetadata *ReferenceLink `json:"snapshotMetadata,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// versions
	Versions *ReferenceLinks `json:"versions,omitempty"`

	// versions index
	VersionsIndex *ReferenceLink `json:"versionsIndex,omitempty"`
}

// Validate validates this maven package
func (m *MavenPackage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifactIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtifactMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionsIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MavenPackage) validateArtifactIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.ArtifactIndex) { // not required
		return nil
	}

	if m.ArtifactIndex != nil {
		if err := m.ArtifactIndex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactIndex")
			}
			return err
		}
	}

	return nil
}

func (m *MavenPackage) validateArtifactMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.ArtifactMetadata) { // not required
		return nil
	}

	if m.ArtifactMetadata != nil {
		if err := m.ArtifactMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *MavenPackage) validateDeletedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("deletedDate", "body", "date-time", m.DeletedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MavenPackage) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Files) { // not required
		return nil
	}

	if m.Files != nil {
		if err := m.Files.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("files")
			}
			return err
		}
	}

	return nil
}

func (m *MavenPackage) validatePom(formats strfmt.Registry) error {

	if swag.IsZero(m.Pom) { // not required
		return nil
	}

	if m.Pom != nil {
		if err := m.Pom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pom")
			}
			return err
		}
	}

	return nil
}

func (m *MavenPackage) validateRequestedFile(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedFile) { // not required
		return nil
	}

	if m.RequestedFile != nil {
		if err := m.RequestedFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedFile")
			}
			return err
		}
	}

	return nil
}

func (m *MavenPackage) validateSnapshotMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.SnapshotMetadata) { // not required
		return nil
	}

	if m.SnapshotMetadata != nil {
		if err := m.SnapshotMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *MavenPackage) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	if m.Versions != nil {
		if err := m.Versions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versions")
			}
			return err
		}
	}

	return nil
}

func (m *MavenPackage) validateVersionsIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionsIndex) { // not required
		return nil
	}

	if m.VersionsIndex != nil {
		if err := m.VersionsIndex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionsIndex")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MavenPackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MavenPackage) UnmarshalBinary(b []byte) error {
	var res MavenPackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}