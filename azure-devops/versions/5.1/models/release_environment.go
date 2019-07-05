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

// ReleaseEnvironment release environment
// swagger:model ReleaseEnvironment
type ReleaseEnvironment struct {

	// Gets list of conditions.
	Conditions []*ReleaseCondition `json:"conditions"`

	// Gets date on which it got created.
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

	// Gets definition environment id.
	DefinitionEnvironmentID int32 `json:"definitionEnvironmentId,omitempty"`

	// Gets list of deploy phases snapshot.
	DeployPhasesSnapshot []*DeployPhase `json:"deployPhasesSnapshot"`

	// Gets deploy steps.
	DeploySteps []*DeploymentAttempt `json:"deploySteps"`

	// Gets environment options.
	EnvironmentOptions *EnvironmentOptions `json:"environmentOptions,omitempty"`

	// Gets the unique identifier of this field.
	ID int32 `json:"id,omitempty"`

	// Gets date on which it got modified.
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modifiedOn,omitempty"`

	// Gets name.
	Name string `json:"name,omitempty"`

	// Gets next scheduled UTC time.
	// Format: date-time
	NextScheduledUtcTime strfmt.DateTime `json:"nextScheduledUtcTime,omitempty"`

	// Gets the identity who is owner for release environment.
	Owner *IdentityRef `json:"owner,omitempty"`

	// Gets list of post deploy approvals snapshot.
	PostApprovalsSnapshot *ReleaseDefinitionApprovals `json:"postApprovalsSnapshot,omitempty"`

	// Gets list of post deploy approvals.
	PostDeployApprovals []*ReleaseApproval `json:"postDeployApprovals"`

	// Post deployment gates snapshot data.
	PostDeploymentGatesSnapshot *ReleaseDefinitionGatesStep `json:"postDeploymentGatesSnapshot,omitempty"`

	// Gets list of pre deploy approvals snapshot.
	PreApprovalsSnapshot *ReleaseDefinitionApprovals `json:"preApprovalsSnapshot,omitempty"`

	// Gets list of pre deploy approvals.
	PreDeployApprovals []*ReleaseApproval `json:"preDeployApprovals"`

	// Pre deployment gates snapshot data.
	PreDeploymentGatesSnapshot *ReleaseDefinitionGatesStep `json:"preDeploymentGatesSnapshot,omitempty"`

	// Gets process parameters.
	ProcessParameters *ProcessParameters `json:"processParameters,omitempty"`

	// Gets rank.
	Rank int32 `json:"rank,omitempty"`

	// Gets release reference which specifies the reference of the release to which this release environment is associated.
	Release *ReleaseShallowReference `json:"release,omitempty"`

	// Gets the identity who created release.
	ReleaseCreatedBy *IdentityRef `json:"releaseCreatedBy,omitempty"`

	// Gets releaseDefinitionReference which specifies the reference of the release definition to which this release environment is associated.
	ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`

	// Gets release id.
	ReleaseID int32 `json:"releaseId,omitempty"`

	// Gets schedule deployment time of release environment.
	// Format: date-time
	ScheduledDeploymentTime strfmt.DateTime `json:"scheduledDeploymentTime,omitempty"`

	// Gets list of schedules.
	Schedules []*ReleaseSchedule `json:"schedules"`

	// Gets environment status.
	// Enum: [undefined notStarted inProgress succeeded canceled rejected queued scheduled partiallySucceeded]
	Status interface{} `json:"status,omitempty"`

	// Gets time to deploy.
	TimeToDeploy float64 `json:"timeToDeploy,omitempty"`

	// Gets trigger reason.
	TriggerReason string `json:"triggerReason,omitempty"`

	// Gets the list of variable groups.
	VariableGroups []*VariableGroup `json:"variableGroups"`

	// Gets the dictionary of variables.
	Variables map[string]ConfigurationVariableValue `json:"variables,omitempty"`
}

// Validate validates this release environment
func (m *ReleaseEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployPhasesSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploySteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextScheduledUtcTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostApprovalsSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostDeployApprovals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostDeploymentGatesSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreApprovalsSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreDeployApprovals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreDeploymentGatesSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledDeploymentTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariableGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseEnvironment) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseEnvironment) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("createdOn", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReleaseEnvironment) validateDeployPhasesSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.DeployPhasesSnapshot) { // not required
		return nil
	}

	for i := 0; i < len(m.DeployPhasesSnapshot); i++ {
		if swag.IsZero(m.DeployPhasesSnapshot[i]) { // not required
			continue
		}

		if m.DeployPhasesSnapshot[i] != nil {
			if err := m.DeployPhasesSnapshot[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployPhasesSnapshot" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseEnvironment) validateDeploySteps(formats strfmt.Registry) error {

	if swag.IsZero(m.DeploySteps) { // not required
		return nil
	}

	for i := 0; i < len(m.DeploySteps); i++ {
		if swag.IsZero(m.DeploySteps[i]) { // not required
			continue
		}

		if m.DeploySteps[i] != nil {
			if err := m.DeploySteps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deploySteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseEnvironment) validateEnvironmentOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentOptions) { // not required
		return nil
	}

	if m.EnvironmentOptions != nil {
		if err := m.EnvironmentOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environmentOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedOn", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReleaseEnvironment) validateNextScheduledUtcTime(formats strfmt.Registry) error {

	if swag.IsZero(m.NextScheduledUtcTime) { // not required
		return nil
	}

	if err := validate.FormatOf("nextScheduledUtcTime", "body", "date-time", m.NextScheduledUtcTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReleaseEnvironment) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validatePostApprovalsSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.PostApprovalsSnapshot) { // not required
		return nil
	}

	if m.PostApprovalsSnapshot != nil {
		if err := m.PostApprovalsSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postApprovalsSnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validatePostDeployApprovals(formats strfmt.Registry) error {

	if swag.IsZero(m.PostDeployApprovals) { // not required
		return nil
	}

	for i := 0; i < len(m.PostDeployApprovals); i++ {
		if swag.IsZero(m.PostDeployApprovals[i]) { // not required
			continue
		}

		if m.PostDeployApprovals[i] != nil {
			if err := m.PostDeployApprovals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postDeployApprovals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseEnvironment) validatePostDeploymentGatesSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.PostDeploymentGatesSnapshot) { // not required
		return nil
	}

	if m.PostDeploymentGatesSnapshot != nil {
		if err := m.PostDeploymentGatesSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDeploymentGatesSnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validatePreApprovalsSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.PreApprovalsSnapshot) { // not required
		return nil
	}

	if m.PreApprovalsSnapshot != nil {
		if err := m.PreApprovalsSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preApprovalsSnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validatePreDeployApprovals(formats strfmt.Registry) error {

	if swag.IsZero(m.PreDeployApprovals) { // not required
		return nil
	}

	for i := 0; i < len(m.PreDeployApprovals); i++ {
		if swag.IsZero(m.PreDeployApprovals[i]) { // not required
			continue
		}

		if m.PreDeployApprovals[i] != nil {
			if err := m.PreDeployApprovals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preDeployApprovals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseEnvironment) validatePreDeploymentGatesSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.PreDeploymentGatesSnapshot) { // not required
		return nil
	}

	if m.PreDeploymentGatesSnapshot != nil {
		if err := m.PreDeploymentGatesSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preDeploymentGatesSnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validateProcessParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessParameters) { // not required
		return nil
	}

	if m.ProcessParameters != nil {
		if err := m.ProcessParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processParameters")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validateRelease(formats strfmt.Registry) error {

	if swag.IsZero(m.Release) { // not required
		return nil
	}

	if m.Release != nil {
		if err := m.Release.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("release")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validateReleaseCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseCreatedBy) { // not required
		return nil
	}

	if m.ReleaseCreatedBy != nil {
		if err := m.ReleaseCreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releaseCreatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validateReleaseDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseDefinition) { // not required
		return nil
	}

	if m.ReleaseDefinition != nil {
		if err := m.ReleaseDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releaseDefinition")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseEnvironment) validateScheduledDeploymentTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledDeploymentTime) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduledDeploymentTime", "body", "date-time", m.ScheduledDeploymentTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReleaseEnvironment) validateSchedules(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedules) { // not required
		return nil
	}

	for i := 0; i < len(m.Schedules); i++ {
		if swag.IsZero(m.Schedules[i]) { // not required
			continue
		}

		if m.Schedules[i] != nil {
			if err := m.Schedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseEnvironment) validateVariableGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.VariableGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.VariableGroups); i++ {
		if swag.IsZero(m.VariableGroups[i]) { // not required
			continue
		}

		if m.VariableGroups[i] != nil {
			if err := m.VariableGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variableGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseEnvironment) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for k := range m.Variables {

		if err := validate.Required("variables"+"."+k, "body", m.Variables[k]); err != nil {
			return err
		}
		if val, ok := m.Variables[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseEnvironment) UnmarshalBinary(b []byte) error {
	var res ReleaseEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}