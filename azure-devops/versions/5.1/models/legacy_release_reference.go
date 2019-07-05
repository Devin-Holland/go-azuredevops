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

// LegacyReleaseReference legacy release reference
// swagger:model LegacyReleaseReference
type LegacyReleaseReference struct {

	// attempt
	Attempt int32 `json:"attempt,omitempty"`

	// environment creation date
	// Format: date-time
	EnvironmentCreationDate strfmt.DateTime `json:"environmentCreationDate,omitempty"`

	// primary artifact build Id
	PrimaryArtifactBuildID int32 `json:"primaryArtifactBuildId,omitempty"`

	// primary artifact project Id
	PrimaryArtifactProjectID string `json:"primaryArtifactProjectId,omitempty"`

	// primary artifact type
	PrimaryArtifactType string `json:"primaryArtifactType,omitempty"`

	// release creation date
	// Format: date-time
	ReleaseCreationDate strfmt.DateTime `json:"releaseCreationDate,omitempty"`

	// release def Id
	ReleaseDefID int32 `json:"releaseDefId,omitempty"`

	// release env def Id
	ReleaseEnvDefID int32 `json:"releaseEnvDefId,omitempty"`

	// release env Id
	ReleaseEnvID int32 `json:"releaseEnvId,omitempty"`

	// release env name
	ReleaseEnvName string `json:"releaseEnvName,omitempty"`

	// release env Uri
	ReleaseEnvURI string `json:"releaseEnvUri,omitempty"`

	// release Id
	ReleaseID int32 `json:"releaseId,omitempty"`

	// release name
	ReleaseName string `json:"releaseName,omitempty"`

	// release ref Id
	ReleaseRefID int32 `json:"releaseRefId,omitempty"`

	// release Uri
	ReleaseURI string `json:"releaseUri,omitempty"`
}

// Validate validates this legacy release reference
func (m *LegacyReleaseReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LegacyReleaseReference) validateEnvironmentCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentCreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("environmentCreationDate", "body", "date-time", m.EnvironmentCreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyReleaseReference) validateReleaseCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseCreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("releaseCreationDate", "body", "date-time", m.ReleaseCreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LegacyReleaseReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LegacyReleaseReference) UnmarshalBinary(b []byte) error {
	var res LegacyReleaseReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}