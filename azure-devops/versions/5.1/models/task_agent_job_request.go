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

// TaskAgentJobRequest A job request for an agent.
// swagger:model TaskAgentJobRequest
type TaskAgentJobRequest struct {

	// agent delays
	AgentDelays []*TaskAgentDelaySource `json:"agentDelays"`

	// agent specification
	AgentSpecification *JObject `json:"agentSpecification,omitempty"`

	// The date/time this request was assigned.
	// Format: date-time
	AssignTime strfmt.DateTime `json:"assignTime,omitempty"`

	// Additional data about the request.
	Data map[string]string `json:"data,omitempty"`

	// The pipeline definition associated with this request
	Definition *TaskOrchestrationOwner `json:"definition,omitempty"`

	// A list of demands required to fulfill this request.
	Demands []*Demand `json:"demands"`

	// expected duration
	ExpectedDuration string `json:"expectedDuration,omitempty"`

	// The date/time this request was finished.
	// Format: date-time
	FinishTime strfmt.DateTime `json:"finishTime,omitempty"`

	// The host which triggered this request.
	// Format: uuid
	HostID strfmt.UUID `json:"hostId,omitempty"`

	// ID of the job resulting from this request.
	// Format: uuid
	JobID strfmt.UUID `json:"jobId,omitempty"`

	// Name of the job resulting from this request.
	JobName string `json:"jobName,omitempty"`

	// The deadline for the agent to renew the lock.
	// Format: date-time
	LockedUntil strfmt.DateTime `json:"lockedUntil,omitempty"`

	// matched agents
	MatchedAgents []*TaskAgentReference `json:"matchedAgents"`

	// matches all agents in pool
	MatchesAllAgentsInPool bool `json:"matchesAllAgentsInPool,omitempty"`

	// orchestration Id
	OrchestrationID string `json:"orchestrationId,omitempty"`

	// The pipeline associated with this request
	Owner *TaskOrchestrationOwner `json:"owner,omitempty"`

	// plan group
	PlanGroup string `json:"planGroup,omitempty"`

	// Internal ID for the orchestration plan connected with this request.
	// Format: uuid
	PlanID strfmt.UUID `json:"planId,omitempty"`

	// Internal detail representing the type of orchestration plan.
	PlanType string `json:"planType,omitempty"`

	// The ID of the pool this request targets
	PoolID int32 `json:"poolId,omitempty"`

	// The ID of the queue this request targets
	QueueID int32 `json:"queueId,omitempty"`

	// The date/time this request was queued.
	// Format: date-time
	QueueTime strfmt.DateTime `json:"queueTime,omitempty"`

	// The date/time this request was receieved by an agent.
	// Format: date-time
	ReceiveTime strfmt.DateTime `json:"receiveTime,omitempty"`

	// ID of the request.
	RequestID int64 `json:"requestId,omitempty"`

	// The agent allocated for this request.
	ReservedAgent *TaskAgentReference `json:"reservedAgent,omitempty"`

	// The result of this request.
	// Enum: [succeeded succeededWithIssues failed canceled skipped abandoned]
	Result interface{} `json:"result,omitempty"`

	// Scope of the pipeline; matches the project ID.
	// Format: uuid
	ScopeID strfmt.UUID `json:"scopeId,omitempty"`

	// The service which owns this request.
	// Format: uuid
	ServiceOwner strfmt.UUID `json:"serviceOwner,omitempty"`

	// status message
	StatusMessage string `json:"statusMessage,omitempty"`

	// user delayed
	UserDelayed bool `json:"userDelayed,omitempty"`
}

// Validate validates this task agent job request
func (m *TaskAgentJobRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentDelays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentSpecification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDemands(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockedUntil(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchedAgents(formats); err != nil {
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

	if err := m.validateReceiveTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReservedAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAgentJobRequest) validateAgentDelays(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentDelays) { // not required
		return nil
	}

	for i := 0; i < len(m.AgentDelays); i++ {
		if swag.IsZero(m.AgentDelays[i]) { // not required
			continue
		}

		if m.AgentDelays[i] != nil {
			if err := m.AgentDelays[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agentDelays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TaskAgentJobRequest) validateAgentSpecification(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentSpecification) { // not required
		return nil
	}

	if m.AgentSpecification != nil {
		if err := m.AgentSpecification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agentSpecification")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgentJobRequest) validateAssignTime(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignTime) { // not required
		return nil
	}

	if err := validate.FormatOf("assignTime", "body", "date-time", m.AssignTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentJobRequest) validateDefinition(formats strfmt.Registry) error {

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

func (m *TaskAgentJobRequest) validateDemands(formats strfmt.Registry) error {

	if swag.IsZero(m.Demands) { // not required
		return nil
	}

	for i := 0; i < len(m.Demands); i++ {
		if swag.IsZero(m.Demands[i]) { // not required
			continue
		}

		if m.Demands[i] != nil {
			if err := m.Demands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("demands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TaskAgentJobRequest) validateFinishTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("finishTime", "body", "date-time", m.FinishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentJobRequest) validateHostID(formats strfmt.Registry) error {

	if swag.IsZero(m.HostID) { // not required
		return nil
	}

	if err := validate.FormatOf("hostId", "body", "uuid", m.HostID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentJobRequest) validateJobID(formats strfmt.Registry) error {

	if swag.IsZero(m.JobID) { // not required
		return nil
	}

	if err := validate.FormatOf("jobId", "body", "uuid", m.JobID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentJobRequest) validateLockedUntil(formats strfmt.Registry) error {

	if swag.IsZero(m.LockedUntil) { // not required
		return nil
	}

	if err := validate.FormatOf("lockedUntil", "body", "date-time", m.LockedUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentJobRequest) validateMatchedAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.MatchedAgents) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchedAgents); i++ {
		if swag.IsZero(m.MatchedAgents[i]) { // not required
			continue
		}

		if m.MatchedAgents[i] != nil {
			if err := m.MatchedAgents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchedAgents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TaskAgentJobRequest) validateOwner(formats strfmt.Registry) error {

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

func (m *TaskAgentJobRequest) validatePlanID(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanID) { // not required
		return nil
	}

	if err := validate.FormatOf("planId", "body", "uuid", m.PlanID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentJobRequest) validateQueueTime(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueTime) { // not required
		return nil
	}

	if err := validate.FormatOf("queueTime", "body", "date-time", m.QueueTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentJobRequest) validateReceiveTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ReceiveTime) { // not required
		return nil
	}

	if err := validate.FormatOf("receiveTime", "body", "date-time", m.ReceiveTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentJobRequest) validateReservedAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.ReservedAgent) { // not required
		return nil
	}

	if m.ReservedAgent != nil {
		if err := m.ReservedAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reservedAgent")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgentJobRequest) validateScopeID(formats strfmt.Registry) error {

	if swag.IsZero(m.ScopeID) { // not required
		return nil
	}

	if err := validate.FormatOf("scopeId", "body", "uuid", m.ScopeID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgentJobRequest) validateServiceOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceOwner) { // not required
		return nil
	}

	if err := validate.FormatOf("serviceOwner", "body", "uuid", m.ServiceOwner.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAgentJobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAgentJobRequest) UnmarshalBinary(b []byte) error {
	var res TaskAgentJobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}