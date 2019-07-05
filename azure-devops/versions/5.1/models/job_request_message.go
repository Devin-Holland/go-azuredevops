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

// JobRequestMessage job request message
// swagger:model JobRequestMessage
type JobRequestMessage struct {

	// environment
	Environment *JobEnvironment `json:"environment,omitempty"`

	// job Id
	// Format: uuid
	JobID strfmt.UUID `json:"jobId,omitempty"`

	// job name
	JobName string `json:"jobName,omitempty"`

	// job ref name
	JobRefName string `json:"jobRefName,omitempty"`

	// message type
	MessageType string `json:"messageType,omitempty"`

	// plan
	Plan *TaskOrchestrationPlanReference `json:"plan,omitempty"`

	// timeline
	Timeline *TimelineReference `json:"timeline,omitempty"`
}

// Validate validates this job request message
func (m *JobRequestMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobRequestMessage) validateEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *JobRequestMessage) validateJobID(formats strfmt.Registry) error {

	if swag.IsZero(m.JobID) { // not required
		return nil
	}

	if err := validate.FormatOf("jobId", "body", "uuid", m.JobID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobRequestMessage) validatePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

func (m *JobRequestMessage) validateTimeline(formats strfmt.Registry) error {

	if swag.IsZero(m.Timeline) { // not required
		return nil
	}

	if m.Timeline != nil {
		if err := m.Timeline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeline")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobRequestMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobRequestMessage) UnmarshalBinary(b []byte) error {
	var res JobRequestMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}