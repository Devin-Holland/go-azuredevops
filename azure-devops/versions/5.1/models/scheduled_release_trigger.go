// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ScheduledReleaseTrigger scheduled release trigger
// swagger:model ScheduledReleaseTrigger
type ScheduledReleaseTrigger struct {
	ReleaseTriggerBase

	// Release schedule for Scheduled Release trigger type.
	Schedule *ReleaseSchedule `json:"schedule,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ScheduledReleaseTrigger) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ReleaseTriggerBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ReleaseTriggerBase = aO0

	// now for regular properties
	var propsScheduledReleaseTrigger struct {
		Schedule *ReleaseSchedule `json:"schedule,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsScheduledReleaseTrigger); err != nil {
		return err
	}
	m.Schedule = propsScheduledReleaseTrigger.Schedule

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ScheduledReleaseTrigger) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ReleaseTriggerBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsScheduledReleaseTrigger struct {
		Schedule *ReleaseSchedule `json:"schedule,omitempty"`
	}
	propsScheduledReleaseTrigger.Schedule = m.Schedule

	jsonDataPropsScheduledReleaseTrigger, errScheduledReleaseTrigger := swag.WriteJSON(propsScheduledReleaseTrigger)
	if errScheduledReleaseTrigger != nil {
		return nil, errScheduledReleaseTrigger
	}
	_parts = append(_parts, jsonDataPropsScheduledReleaseTrigger)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this scheduled release trigger
func (m *ScheduledReleaseTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ReleaseTriggerBase
	if err := m.ReleaseTriggerBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduledReleaseTrigger) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduledReleaseTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledReleaseTrigger) UnmarshalBinary(b []byte) error {
	var res ScheduledReleaseTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
