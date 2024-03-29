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

// ContinuousIntegrationTrigger Represents a continuous integration (CI) trigger.
// swagger:model ContinuousIntegrationTrigger
type ContinuousIntegrationTrigger struct {
	BuildTrigger

	// Indicates whether changes should be batched while another CI build is running.
	BatchChanges bool `json:"batchChanges,omitempty"`

	// branch filters
	BranchFilters []string `json:"branchFilters"`

	// The maximum number of simultaneous CI builds that will run per branch.
	MaxConcurrentBuildsPerBranch int32 `json:"maxConcurrentBuildsPerBranch,omitempty"`

	// path filters
	PathFilters []string `json:"pathFilters"`

	// The polling interval, in seconds.
	PollingInterval int32 `json:"pollingInterval,omitempty"`

	// The ID of the job used to poll an external repository.
	// Format: uuid
	PollingJobID strfmt.UUID `json:"pollingJobId,omitempty"`

	// settings source type
	SettingsSourceType int32 `json:"settingsSourceType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ContinuousIntegrationTrigger) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuildTrigger
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuildTrigger = aO0

	// now for regular properties
	var propsContinuousIntegrationTrigger struct {
		BatchChanges bool `json:"batchChanges,omitempty"`

		BranchFilters []string `json:"branchFilters"`

		MaxConcurrentBuildsPerBranch int32 `json:"maxConcurrentBuildsPerBranch,omitempty"`

		PathFilters []string `json:"pathFilters"`

		PollingInterval int32 `json:"pollingInterval,omitempty"`

		PollingJobID strfmt.UUID `json:"pollingJobId,omitempty"`

		SettingsSourceType int32 `json:"settingsSourceType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsContinuousIntegrationTrigger); err != nil {
		return err
	}
	m.BatchChanges = propsContinuousIntegrationTrigger.BatchChanges

	m.BranchFilters = propsContinuousIntegrationTrigger.BranchFilters

	m.MaxConcurrentBuildsPerBranch = propsContinuousIntegrationTrigger.MaxConcurrentBuildsPerBranch

	m.PathFilters = propsContinuousIntegrationTrigger.PathFilters

	m.PollingInterval = propsContinuousIntegrationTrigger.PollingInterval

	m.PollingJobID = propsContinuousIntegrationTrigger.PollingJobID

	m.SettingsSourceType = propsContinuousIntegrationTrigger.SettingsSourceType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ContinuousIntegrationTrigger) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BuildTrigger)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsContinuousIntegrationTrigger struct {
		BatchChanges bool `json:"batchChanges,omitempty"`

		BranchFilters []string `json:"branchFilters"`

		MaxConcurrentBuildsPerBranch int32 `json:"maxConcurrentBuildsPerBranch,omitempty"`

		PathFilters []string `json:"pathFilters"`

		PollingInterval int32 `json:"pollingInterval,omitempty"`

		PollingJobID strfmt.UUID `json:"pollingJobId,omitempty"`

		SettingsSourceType int32 `json:"settingsSourceType,omitempty"`
	}
	propsContinuousIntegrationTrigger.BatchChanges = m.BatchChanges

	propsContinuousIntegrationTrigger.BranchFilters = m.BranchFilters

	propsContinuousIntegrationTrigger.MaxConcurrentBuildsPerBranch = m.MaxConcurrentBuildsPerBranch

	propsContinuousIntegrationTrigger.PathFilters = m.PathFilters

	propsContinuousIntegrationTrigger.PollingInterval = m.PollingInterval

	propsContinuousIntegrationTrigger.PollingJobID = m.PollingJobID

	propsContinuousIntegrationTrigger.SettingsSourceType = m.SettingsSourceType

	jsonDataPropsContinuousIntegrationTrigger, errContinuousIntegrationTrigger := swag.WriteJSON(propsContinuousIntegrationTrigger)
	if errContinuousIntegrationTrigger != nil {
		return nil, errContinuousIntegrationTrigger
	}
	_parts = append(_parts, jsonDataPropsContinuousIntegrationTrigger)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this continuous integration trigger
func (m *ContinuousIntegrationTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuildTrigger
	if err := m.BuildTrigger.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePollingJobID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContinuousIntegrationTrigger) validatePollingJobID(formats strfmt.Registry) error {

	if swag.IsZero(m.PollingJobID) { // not required
		return nil
	}

	if err := validate.FormatOf("pollingJobId", "body", "uuid", m.PollingJobID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContinuousIntegrationTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContinuousIntegrationTrigger) UnmarshalBinary(b []byte) error {
	var res ContinuousIntegrationTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
