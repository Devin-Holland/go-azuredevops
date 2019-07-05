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

// ReleaseTaskLogUpdatedEvent release task log updated event
// swagger:model ReleaseTaskLogUpdatedEvent
type ReleaseTaskLogUpdatedEvent struct {
	RealtimeReleaseEvent

	// lines
	Lines []string `json:"lines"`

	// step record Id
	// Format: uuid
	StepRecordID strfmt.UUID `json:"stepRecordId,omitempty"`

	// timeline record Id
	// Format: uuid
	TimelineRecordID strfmt.UUID `json:"timelineRecordId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ReleaseTaskLogUpdatedEvent) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RealtimeReleaseEvent
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RealtimeReleaseEvent = aO0

	// now for regular properties
	var propsReleaseTaskLogUpdatedEvent struct {
		Lines []string `json:"lines"`

		StepRecordID strfmt.UUID `json:"stepRecordId,omitempty"`

		TimelineRecordID strfmt.UUID `json:"timelineRecordId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsReleaseTaskLogUpdatedEvent); err != nil {
		return err
	}
	m.Lines = propsReleaseTaskLogUpdatedEvent.Lines

	m.StepRecordID = propsReleaseTaskLogUpdatedEvent.StepRecordID

	m.TimelineRecordID = propsReleaseTaskLogUpdatedEvent.TimelineRecordID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ReleaseTaskLogUpdatedEvent) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.RealtimeReleaseEvent)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsReleaseTaskLogUpdatedEvent struct {
		Lines []string `json:"lines"`

		StepRecordID strfmt.UUID `json:"stepRecordId,omitempty"`

		TimelineRecordID strfmt.UUID `json:"timelineRecordId,omitempty"`
	}
	propsReleaseTaskLogUpdatedEvent.Lines = m.Lines

	propsReleaseTaskLogUpdatedEvent.StepRecordID = m.StepRecordID

	propsReleaseTaskLogUpdatedEvent.TimelineRecordID = m.TimelineRecordID

	jsonDataPropsReleaseTaskLogUpdatedEvent, errReleaseTaskLogUpdatedEvent := swag.WriteJSON(propsReleaseTaskLogUpdatedEvent)
	if errReleaseTaskLogUpdatedEvent != nil {
		return nil, errReleaseTaskLogUpdatedEvent
	}
	_parts = append(_parts, jsonDataPropsReleaseTaskLogUpdatedEvent)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this release task log updated event
func (m *ReleaseTaskLogUpdatedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RealtimeReleaseEvent
	if err := m.RealtimeReleaseEvent.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStepRecordID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimelineRecordID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseTaskLogUpdatedEvent) validateStepRecordID(formats strfmt.Registry) error {

	if swag.IsZero(m.StepRecordID) { // not required
		return nil
	}

	if err := validate.FormatOf("stepRecordId", "body", "uuid", m.StepRecordID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReleaseTaskLogUpdatedEvent) validateTimelineRecordID(formats strfmt.Registry) error {

	if swag.IsZero(m.TimelineRecordID) { // not required
		return nil
	}

	if err := validate.FormatOf("timelineRecordId", "body", "uuid", m.TimelineRecordID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseTaskLogUpdatedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseTaskLogUpdatedEvent) UnmarshalBinary(b []byte) error {
	var res ReleaseTaskLogUpdatedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}