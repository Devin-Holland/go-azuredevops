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

// BuildReference2 build reference2
// swagger:model BuildReference2
type BuildReference2 struct {

	// branch name
	BranchName string `json:"branchName,omitempty"`

	// build configuration Id
	BuildConfigurationID int32 `json:"buildConfigurationId,omitempty"`

	// build definition Id
	BuildDefinitionID int32 `json:"buildDefinitionId,omitempty"`

	// build deleted
	BuildDeleted bool `json:"buildDeleted,omitempty"`

	// build flavor
	BuildFlavor string `json:"buildFlavor,omitempty"`

	// build Id
	BuildID int32 `json:"buildId,omitempty"`

	// build number
	BuildNumber string `json:"buildNumber,omitempty"`

	// build platform
	BuildPlatform string `json:"buildPlatform,omitempty"`

	// build system
	BuildSystem string `json:"buildSystem,omitempty"`

	// build Uri
	BuildURI string `json:"buildUri,omitempty"`

	// coverage Id
	CoverageID int32 `json:"coverageId,omitempty"`

	// created date
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// project Id
	// Format: uuid
	ProjectID strfmt.UUID `json:"projectId,omitempty"`

	// repo Id
	RepoID string `json:"repoId,omitempty"`

	// repo type
	RepoType string `json:"repoType,omitempty"`

	// source version
	SourceVersion string `json:"sourceVersion,omitempty"`
}

// Validate validates this build reference2
func (m *BuildReference2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildReference2) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildReference2) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectId", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildReference2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildReference2) UnmarshalBinary(b []byte) error {
	var res BuildReference2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}