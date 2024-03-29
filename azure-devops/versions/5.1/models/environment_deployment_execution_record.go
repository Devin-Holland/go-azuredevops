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

// EnvironmentDeploymentExecutionRecord EnvironmentDeploymentExecutionRecord.
// swagger:model EnvironmentDeploymentExecutionRecord
type EnvironmentDeploymentExecutionRecord struct {

	// Definition of the environment deployment execution owner
	Definition *TaskOrchestrationOwner `json:"definition,omitempty"`

	// Id of the Environment
	EnvironmentID int32 `json:"environmentId,omitempty"`

	// Finish time of the environment deployment execution
	// Format: date-time
	FinishTime strfmt.DateTime `json:"finishTime,omitempty"`

	// Id of the Environment deployment execution history record
	ID int64 `json:"id,omitempty"`

	// Job Attempt
	JobAttempt int32 `json:"jobAttempt,omitempty"`

	// Job name
	JobName string `json:"jobName,omitempty"`

	// Owner of the environment deployment execution record
	Owner *TaskOrchestrationOwner `json:"owner,omitempty"`

	// Plan Id
	// Format: uuid
	PlanID strfmt.UUID `json:"planId,omitempty"`

	// Plan type of the environment deployment execution record
	PlanType string `json:"planType,omitempty"`

	// Queue time of the environment deployment execution
	// Format: date-time
	QueueTime strfmt.DateTime `json:"queueTime,omitempty"`

	// Request identifier of the Environment deployment execution history record
	RequestIdentifier string `json:"requestIdentifier,omitempty"`

	// Resource Id
	ResourceID int32 `json:"resourceId,omitempty"`

	// Result of the environment deployment execution
	// Enum: [succeeded succeededWithIssues failed canceled skipped abandoned]
	Result interface{} `json:"result,omitempty"`

	// Project Id
	// Format: uuid
	ScopeID strfmt.UUID `json:"scopeId,omitempty"`

	// Service owner Id
	// Format: uuid
	ServiceOwner strfmt.UUID `json:"serviceOwner,omitempty"`

	// Stage Attempt
	StageAttempt int32 `json:"stageAttempt,omitempty"`

	// Stage name
	StageName string `json:"stageName,omitempty"`

	// Start time of the environment deployment execution
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`
}

// Validate validates this environment deployment execution record
func (m *EnvironmentDeploymentExecutionRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentDeploymentExecutionRecord) validateDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.Definition) { // not required
		return nil
	}

	if m.Definition != nil {
		if err := m.Definition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("definition")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentDeploymentExecutionRecord) validateFinishTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("finishTime", "body", "date-time", m.FinishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentDeploymentExecutionRecord) validateOwner(formats strfmt.Registry) error {

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

func (m *EnvironmentDeploymentExecutionRecord) validatePlanID(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanID) { // not required
		return nil
	}

	if err := validate.FormatOf("planId", "body", "uuid", m.PlanID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentDeploymentExecutionRecord) validateQueueTime(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueTime) { // not required
		return nil
	}

	if err := validate.FormatOf("queueTime", "body", "date-time", m.QueueTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentDeploymentExecutionRecord) validateScopeID(formats strfmt.Registry) error {

	if swag.IsZero(m.ScopeID) { // not required
		return nil
	}

	if err := validate.FormatOf("scopeId", "body", "uuid", m.ScopeID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentDeploymentExecutionRecord) validateServiceOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceOwner) { // not required
		return nil
	}

	if err := validate.FormatOf("serviceOwner", "body", "uuid", m.ServiceOwner.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentDeploymentExecutionRecord) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentDeploymentExecutionRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentDeploymentExecutionRecord) UnmarshalBinary(b []byte) error {
	var res EnvironmentDeploymentExecutionRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
