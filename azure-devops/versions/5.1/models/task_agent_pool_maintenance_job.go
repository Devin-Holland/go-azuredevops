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

// TaskAgentPoolMaintenanceJob task agent pool maintenance job
// swagger:model TaskAgentPoolMaintenanceJob
type TaskAgentPoolMaintenanceJob struct {

	// The maintenance definition for the maintenance job
	DefinitionID int32 `json:"definitionId,omitempty"`

	// The total error counts during the maintenance job
	ErrorCount int32 `json:"errorCount,omitempty"`

	// Time that the maintenance job was completed
	// Format: date-time
	FinishTime strfmt.DateTime `json:"finishTime,omitempty"`

	// Id of the maintenance job
	JobID int32 `json:"jobId,omitempty"`

	// The log download url for the maintenance job
	LogsDownloadURL string `json:"logsDownloadUrl,omitempty"`

	// Orchestration/Plan Id for the maintenance job
	// Format: uuid
	OrchestrationID strfmt.UUID `json:"orchestrationId,omitempty"`

	// Pool reference for the maintenance job
	Pool *TaskAgentPoolReference `json:"pool,omitempty"`

	// Time that the maintenance job was queued
	// Format: date-time
	QueueTime strfmt.DateTime `json:"queueTime,omitempty"`

	// The identity that queued the maintenance job
	RequestedBy *IdentityRef `json:"requestedBy,omitempty"`

	// The maintenance job result
	// Enum: [succeeded failed canceled]
	Result interface{} `json:"result,omitempty"`

	// Time that the maintenance job was started
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// Status of the maintenance job
	// Enum: [inProgress completed cancelling queued]
	Status interface{} `json:"status,omitempty"`

	// target agents
	TargetAgents []*TaskAgentPoolMaintenanceJobTargetAgent `json:"targetAgents"`

	// The total warning counts during the maintenance job
	WarningCount int32 `json:"warningCount,omitempty"`
}

// Validate validates this task agent pool maintenance job
func (m *TaskAgentPoolMaintenanceJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrchestrationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetAgents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAgentPoolMaintenanceJob) validateFinishTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("finishTime", "body", "date-time", m.FinishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentPoolMaintenanceJob) validateOrchestrationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrchestrationID) { // not required
		return nil
	}

	if err := validate.FormatOf("orchestrationId", "body", "uuid", m.OrchestrationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentPoolMaintenanceJob) validatePool(formats strfmt.Registry) error {

	if swag.IsZero(m.Pool) { // not required
		return nil
	}

	if m.Pool != nil {
		if err := m.Pool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pool")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgentPoolMaintenanceJob) validateQueueTime(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueTime) { // not required
		return nil
	}

	if err := validate.FormatOf("queueTime", "body", "date-time", m.QueueTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentPoolMaintenanceJob) validateRequestedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedBy) { // not required
		return nil
	}

	if m.RequestedBy != nil {
		if err := m.RequestedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgentPoolMaintenanceJob) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentPoolMaintenanceJob) validateTargetAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetAgents) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetAgents); i++ {
		if swag.IsZero(m.TargetAgents[i]) { // not required
			continue
		}

		if m.TargetAgents[i] != nil {
			if err := m.TargetAgents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetAgents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAgentPoolMaintenanceJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAgentPoolMaintenanceJob) UnmarshalBinary(b []byte) error {
	var res TaskAgentPoolMaintenanceJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
