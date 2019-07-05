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

// TaskOrchestrationQueuedPlan task orchestration queued plan
// swagger:model TaskOrchestrationQueuedPlan
type TaskOrchestrationQueuedPlan struct {

	// assign time
	// Format: date-time
	AssignTime strfmt.DateTime `json:"assignTime,omitempty"`

	// definition
	Definition *TaskOrchestrationOwner `json:"definition,omitempty"`

	// owner
	Owner *TaskOrchestrationOwner `json:"owner,omitempty"`

	// plan group
	PlanGroup string `json:"planGroup,omitempty"`

	// plan Id
	// Format: uuid
	PlanID strfmt.UUID `json:"planId,omitempty"`

	// pool Id
	PoolID int32 `json:"poolId,omitempty"`

	// queue position
	QueuePosition int32 `json:"queuePosition,omitempty"`

	// queue time
	// Format: date-time
	QueueTime strfmt.DateTime `json:"queueTime,omitempty"`

	// scope identifier
	// Format: uuid
	ScopeIdentifier strfmt.UUID `json:"scopeIdentifier,omitempty"`
}

// Validate validates this task orchestration queued plan
func (m *TaskOrchestrationQueuedPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefinition(formats); err != nil {
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

	if err := m.validateScopeIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskOrchestrationQueuedPlan) validateAssignTime(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignTime) { // not required
		return nil
	}

	if err := validate.FormatOf("assignTime", "body", "date-time", m.AssignTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskOrchestrationQueuedPlan) validateDefinition(formats strfmt.Registry) error {

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

func (m *TaskOrchestrationQueuedPlan) validateOwner(formats strfmt.Registry) error {

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

func (m *TaskOrchestrationQueuedPlan) validatePlanID(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanID) { // not required
		return nil
	}

	if err := validate.FormatOf("planId", "body", "uuid", m.PlanID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskOrchestrationQueuedPlan) validateQueueTime(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueTime) { // not required
		return nil
	}

	if err := validate.FormatOf("queueTime", "body", "date-time", m.QueueTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskOrchestrationQueuedPlan) validateScopeIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.ScopeIdentifier) { // not required
		return nil
	}

	if err := validate.FormatOf("scopeIdentifier", "body", "uuid", m.ScopeIdentifier.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskOrchestrationQueuedPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskOrchestrationQueuedPlan) UnmarshalBinary(b []byte) error {
	var res TaskOrchestrationQueuedPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
