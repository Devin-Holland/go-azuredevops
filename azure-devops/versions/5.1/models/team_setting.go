// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TeamSetting Data contract for TeamSettings
// swagger:model TeamSetting
type TeamSetting struct {
	TeamSettingsDataContractBase

	// Backlog Iteration
	BacklogIteration *TeamSettingsIteration `json:"backlogIteration,omitempty"`

	// Information about categories that are visible on the backlog.
	BacklogVisibilities map[string]bool `json:"backlogVisibilities,omitempty"`

	// BugsBehavior (Off, AsTasks, AsRequirements, ...)
	// Enum: [off asRequirements asTasks]
	BugsBehavior interface{} `json:"bugsBehavior,omitempty"`

	// Default Iteration, the iteration used when creating a new work item on the queries page.
	DefaultIteration *TeamSettingsIteration `json:"defaultIteration,omitempty"`

	// Default Iteration macro (if any)
	DefaultIterationMacro string `json:"defaultIterationMacro,omitempty"`

	// Days that the team is working
	WorkingDays []interface{} `json:"workingDays"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TeamSetting) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TeamSettingsDataContractBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TeamSettingsDataContractBase = aO0

	// now for regular properties
	var propsTeamSetting struct {
		BacklogIteration *TeamSettingsIteration `json:"backlogIteration,omitempty"`

		BacklogVisibilities map[string]bool `json:"backlogVisibilities,omitempty"`

		BugsBehavior interface{} `json:"bugsBehavior,omitempty"`

		DefaultIteration *TeamSettingsIteration `json:"defaultIteration,omitempty"`

		DefaultIterationMacro string `json:"defaultIterationMacro,omitempty"`

		WorkingDays []interface{} `json:"workingDays"`
	}
	if err := swag.ReadJSON(raw, &propsTeamSetting); err != nil {
		return err
	}
	m.BacklogIteration = propsTeamSetting.BacklogIteration

	m.BacklogVisibilities = propsTeamSetting.BacklogVisibilities

	m.BugsBehavior = propsTeamSetting.BugsBehavior

	m.DefaultIteration = propsTeamSetting.DefaultIteration

	m.DefaultIterationMacro = propsTeamSetting.DefaultIterationMacro

	m.WorkingDays = propsTeamSetting.WorkingDays

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TeamSetting) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TeamSettingsDataContractBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTeamSetting struct {
		BacklogIteration *TeamSettingsIteration `json:"backlogIteration,omitempty"`

		BacklogVisibilities map[string]bool `json:"backlogVisibilities,omitempty"`

		BugsBehavior interface{} `json:"bugsBehavior,omitempty"`

		DefaultIteration *TeamSettingsIteration `json:"defaultIteration,omitempty"`

		DefaultIterationMacro string `json:"defaultIterationMacro,omitempty"`

		WorkingDays []interface{} `json:"workingDays"`
	}
	propsTeamSetting.BacklogIteration = m.BacklogIteration

	propsTeamSetting.BacklogVisibilities = m.BacklogVisibilities

	propsTeamSetting.BugsBehavior = m.BugsBehavior

	propsTeamSetting.DefaultIteration = m.DefaultIteration

	propsTeamSetting.DefaultIterationMacro = m.DefaultIterationMacro

	propsTeamSetting.WorkingDays = m.WorkingDays

	jsonDataPropsTeamSetting, errTeamSetting := swag.WriteJSON(propsTeamSetting)
	if errTeamSetting != nil {
		return nil, errTeamSetting
	}
	_parts = append(_parts, jsonDataPropsTeamSetting)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this team setting
func (m *TeamSetting) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TeamSettingsDataContractBase
	if err := m.TeamSettingsDataContractBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBacklogIteration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultIteration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkingDays(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamSetting) validateBacklogIteration(formats strfmt.Registry) error {

	if swag.IsZero(m.BacklogIteration) { // not required
		return nil
	}

	if m.BacklogIteration != nil {
		if err := m.BacklogIteration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backlogIteration")
			}
			return err
		}
	}

	return nil
}

func (m *TeamSetting) validateDefaultIteration(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultIteration) { // not required
		return nil
	}

	if m.DefaultIteration != nil {
		if err := m.DefaultIteration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultIteration")
			}
			return err
		}
	}

	return nil
}

var teamSettingWorkingDaysItemsEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["sunday","monday","tuesday","wednesday","thursday","friday","saturday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		teamSettingWorkingDaysItemsEnum = append(teamSettingWorkingDaysItemsEnum, v)
	}
}

func (m *TeamSetting) validateWorkingDaysItemsEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, teamSettingWorkingDaysItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *TeamSetting) validateWorkingDays(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkingDays) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamSetting) UnmarshalBinary(b []byte) error {
	var res TeamSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}