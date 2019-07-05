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

// TaskAgent A task agent.
// swagger:model TaskAgent
type TaskAgent struct {
	TaskAgentReference

	// The agent cloud request that's currently associated with this agent.
	AssignedAgentCloudRequest *TaskAgentCloudRequest `json:"assignedAgentCloudRequest,omitempty"`

	// The request which is currently assigned to this agent.
	AssignedRequest *TaskAgentJobRequest `json:"assignedRequest,omitempty"`

	// Authorization information for this agent.
	Authorization *TaskAgentAuthorization `json:"authorization,omitempty"`

	// Date on which this agent was created.
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

	// The last request which was completed by this agent.
	LastCompletedRequest *TaskAgentJobRequest `json:"lastCompletedRequest,omitempty"`

	// Maximum job parallelism allowed for this agent.
	MaxParallelism int32 `json:"maxParallelism,omitempty"`

	// Pending update for this agent.
	PendingUpdate *TaskAgentUpdate `json:"pendingUpdate,omitempty"`

	// properties
	Properties *PropertiesCollection `json:"properties,omitempty"`

	// Date on which the last connectivity status change occurred.
	// Format: date-time
	StatusChangedOn strfmt.DateTime `json:"statusChangedOn,omitempty"`

	// system capabilities
	SystemCapabilities map[string]string `json:"systemCapabilities,omitempty"`

	// user capabilities
	UserCapabilities map[string]string `json:"userCapabilities,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TaskAgent) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TaskAgentReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TaskAgentReference = aO0

	// now for regular properties
	var propsTaskAgent struct {
		AssignedAgentCloudRequest *TaskAgentCloudRequest `json:"assignedAgentCloudRequest,omitempty"`

		AssignedRequest *TaskAgentJobRequest `json:"assignedRequest,omitempty"`

		Authorization *TaskAgentAuthorization `json:"authorization,omitempty"`

		CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

		LastCompletedRequest *TaskAgentJobRequest `json:"lastCompletedRequest,omitempty"`

		MaxParallelism int32 `json:"maxParallelism,omitempty"`

		PendingUpdate *TaskAgentUpdate `json:"pendingUpdate,omitempty"`

		Properties *PropertiesCollection `json:"properties,omitempty"`

		StatusChangedOn strfmt.DateTime `json:"statusChangedOn,omitempty"`

		SystemCapabilities map[string]string `json:"systemCapabilities,omitempty"`

		UserCapabilities map[string]string `json:"userCapabilities,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsTaskAgent); err != nil {
		return err
	}
	m.AssignedAgentCloudRequest = propsTaskAgent.AssignedAgentCloudRequest

	m.AssignedRequest = propsTaskAgent.AssignedRequest

	m.Authorization = propsTaskAgent.Authorization

	m.CreatedOn = propsTaskAgent.CreatedOn

	m.LastCompletedRequest = propsTaskAgent.LastCompletedRequest

	m.MaxParallelism = propsTaskAgent.MaxParallelism

	m.PendingUpdate = propsTaskAgent.PendingUpdate

	m.Properties = propsTaskAgent.Properties

	m.StatusChangedOn = propsTaskAgent.StatusChangedOn

	m.SystemCapabilities = propsTaskAgent.SystemCapabilities

	m.UserCapabilities = propsTaskAgent.UserCapabilities

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TaskAgent) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TaskAgentReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTaskAgent struct {
		AssignedAgentCloudRequest *TaskAgentCloudRequest `json:"assignedAgentCloudRequest,omitempty"`

		AssignedRequest *TaskAgentJobRequest `json:"assignedRequest,omitempty"`

		Authorization *TaskAgentAuthorization `json:"authorization,omitempty"`

		CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

		LastCompletedRequest *TaskAgentJobRequest `json:"lastCompletedRequest,omitempty"`

		MaxParallelism int32 `json:"maxParallelism,omitempty"`

		PendingUpdate *TaskAgentUpdate `json:"pendingUpdate,omitempty"`

		Properties *PropertiesCollection `json:"properties,omitempty"`

		StatusChangedOn strfmt.DateTime `json:"statusChangedOn,omitempty"`

		SystemCapabilities map[string]string `json:"systemCapabilities,omitempty"`

		UserCapabilities map[string]string `json:"userCapabilities,omitempty"`
	}
	propsTaskAgent.AssignedAgentCloudRequest = m.AssignedAgentCloudRequest

	propsTaskAgent.AssignedRequest = m.AssignedRequest

	propsTaskAgent.Authorization = m.Authorization

	propsTaskAgent.CreatedOn = m.CreatedOn

	propsTaskAgent.LastCompletedRequest = m.LastCompletedRequest

	propsTaskAgent.MaxParallelism = m.MaxParallelism

	propsTaskAgent.PendingUpdate = m.PendingUpdate

	propsTaskAgent.Properties = m.Properties

	propsTaskAgent.StatusChangedOn = m.StatusChangedOn

	propsTaskAgent.SystemCapabilities = m.SystemCapabilities

	propsTaskAgent.UserCapabilities = m.UserCapabilities

	jsonDataPropsTaskAgent, errTaskAgent := swag.WriteJSON(propsTaskAgent)
	if errTaskAgent != nil {
		return nil, errTaskAgent
	}
	_parts = append(_parts, jsonDataPropsTaskAgent)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this task agent
func (m *TaskAgent) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TaskAgentReference
	if err := m.TaskAgentReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedAgentCloudRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastCompletedRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePendingUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusChangedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAgent) validateAssignedAgentCloudRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignedAgentCloudRequest) { // not required
		return nil
	}

	if m.AssignedAgentCloudRequest != nil {
		if err := m.AssignedAgentCloudRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignedAgentCloudRequest")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgent) validateAssignedRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignedRequest) { // not required
		return nil
	}

	if m.AssignedRequest != nil {
		if err := m.AssignedRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignedRequest")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgent) validateAuthorization(formats strfmt.Registry) error {

	if swag.IsZero(m.Authorization) { // not required
		return nil
	}

	if m.Authorization != nil {
		if err := m.Authorization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorization")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgent) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("createdOn", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskAgent) validateLastCompletedRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.LastCompletedRequest) { // not required
		return nil
	}

	if m.LastCompletedRequest != nil {
		if err := m.LastCompletedRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastCompletedRequest")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgent) validatePendingUpdate(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingUpdate) { // not required
		return nil
	}

	if m.PendingUpdate != nil {
		if err := m.PendingUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pendingUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgent) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgent) validateStatusChangedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusChangedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("statusChangedOn", "body", "date-time", m.StatusChangedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAgent) UnmarshalBinary(b []byte) error {
	var res TaskAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}