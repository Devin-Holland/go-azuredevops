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

// LegacyBuildConfiguration legacy build configuration
// swagger:model LegacyBuildConfiguration
type LegacyBuildConfiguration struct {

	// branch name
	BranchName string `json:"branchName,omitempty"`

	// build configuration Id
	BuildConfigurationID int32 `json:"buildConfigurationId,omitempty"`

	// build definition Id
	BuildDefinitionID int32 `json:"buildDefinitionId,omitempty"`

	// build definition name
	BuildDefinitionName string `json:"buildDefinitionName,omitempty"`

	// build flavor
	BuildFlavor string `json:"buildFlavor,omitempty"`

	// build Id
	BuildID int32 `json:"buildId,omitempty"`

	// build number
	BuildNumber string `json:"buildNumber,omitempty"`

	// build platform
	BuildPlatform string `json:"buildPlatform,omitempty"`

	// build quality
	BuildQuality string `json:"buildQuality,omitempty"`

	// build system
	BuildSystem string `json:"buildSystem,omitempty"`

	// build Uri
	BuildURI string `json:"buildUri,omitempty"`

	// completed date
	// Format: date-time
	CompletedDate strfmt.DateTime `json:"completedDate,omitempty"`

	// created date
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// old build configuration Id
	OldBuildConfigurationID int32 `json:"oldBuildConfigurationId,omitempty"`

	// repository Id
	RepositoryID string `json:"repositoryId,omitempty"`

	// repository type
	RepositoryType string `json:"repositoryType,omitempty"`

	// source version
	SourceVersion string `json:"sourceVersion,omitempty"`

	// team project name
	TeamProjectName string `json:"teamProjectName,omitempty"`
}

// Validate validates this legacy build configuration
func (m *LegacyBuildConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LegacyBuildConfiguration) validateCompletedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("completedDate", "body", "date-time", m.CompletedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyBuildConfiguration) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LegacyBuildConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LegacyBuildConfiguration) UnmarshalBinary(b []byte) error {
	var res LegacyBuildConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}