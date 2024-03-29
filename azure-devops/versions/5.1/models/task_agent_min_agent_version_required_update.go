// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TaskAgentMinAgentVersionRequiredUpdate task agent min agent version required update
// swagger:model TaskAgentMinAgentVersionRequiredUpdate
type TaskAgentMinAgentVersionRequiredUpdate struct {
	TaskAgentUpdateReason

	// job definition
	JobDefinition *TaskOrchestrationOwner `json:"jobDefinition,omitempty"`

	// job owner
	JobOwner *TaskOrchestrationOwner `json:"jobOwner,omitempty"`

	// min agent version
	MinAgentVersion *Demand `json:"minAgentVersion,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TaskAgentMinAgentVersionRequiredUpdate) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TaskAgentUpdateReason
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TaskAgentUpdateReason = aO0

	// now for regular properties
	var propsTaskAgentMinAgentVersionRequiredUpdate struct {
		JobDefinition *TaskOrchestrationOwner `json:"jobDefinition,omitempty"`

		JobOwner *TaskOrchestrationOwner `json:"jobOwner,omitempty"`

		MinAgentVersion *Demand `json:"minAgentVersion,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsTaskAgentMinAgentVersionRequiredUpdate); err != nil {
		return err
	}
	m.JobDefinition = propsTaskAgentMinAgentVersionRequiredUpdate.JobDefinition

	m.JobOwner = propsTaskAgentMinAgentVersionRequiredUpdate.JobOwner

	m.MinAgentVersion = propsTaskAgentMinAgentVersionRequiredUpdate.MinAgentVersion

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TaskAgentMinAgentVersionRequiredUpdate) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TaskAgentUpdateReason)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTaskAgentMinAgentVersionRequiredUpdate struct {
		JobDefinition *TaskOrchestrationOwner `json:"jobDefinition,omitempty"`

		JobOwner *TaskOrchestrationOwner `json:"jobOwner,omitempty"`

		MinAgentVersion *Demand `json:"minAgentVersion,omitempty"`
	}
	propsTaskAgentMinAgentVersionRequiredUpdate.JobDefinition = m.JobDefinition

	propsTaskAgentMinAgentVersionRequiredUpdate.JobOwner = m.JobOwner

	propsTaskAgentMinAgentVersionRequiredUpdate.MinAgentVersion = m.MinAgentVersion

	jsonDataPropsTaskAgentMinAgentVersionRequiredUpdate, errTaskAgentMinAgentVersionRequiredUpdate := swag.WriteJSON(propsTaskAgentMinAgentVersionRequiredUpdate)
	if errTaskAgentMinAgentVersionRequiredUpdate != nil {
		return nil, errTaskAgentMinAgentVersionRequiredUpdate
	}
	_parts = append(_parts, jsonDataPropsTaskAgentMinAgentVersionRequiredUpdate)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this task agent min agent version required update
func (m *TaskAgentMinAgentVersionRequiredUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TaskAgentUpdateReason
	if err := m.TaskAgentUpdateReason.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinAgentVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAgentMinAgentVersionRequiredUpdate) validateJobDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.JobDefinition) { // not required
		return nil
	}

	if m.JobDefinition != nil {
		if err := m.JobDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobDefinition")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgentMinAgentVersionRequiredUpdate) validateJobOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.JobOwner) { // not required
		return nil
	}

	if m.JobOwner != nil {
		if err := m.JobOwner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobOwner")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgentMinAgentVersionRequiredUpdate) validateMinAgentVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.MinAgentVersion) { // not required
		return nil
	}

	if m.MinAgentVersion != nil {
		if err := m.MinAgentVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("minAgentVersion")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAgentMinAgentVersionRequiredUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAgentMinAgentVersionRequiredUpdate) UnmarshalBinary(b []byte) error {
	var res TaskAgentMinAgentVersionRequiredUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
