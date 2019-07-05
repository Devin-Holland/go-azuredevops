// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ScheduleTrigger Represents a schedule trigger.
// swagger:model ScheduleTrigger
type ScheduleTrigger struct {
	BuildTrigger

	// schedules
	Schedules []*Schedule `json:"schedules"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ScheduleTrigger) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuildTrigger
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuildTrigger = aO0

	// now for regular properties
	var propsScheduleTrigger struct {
		Schedules []*Schedule `json:"schedules"`
	}
	if err := swag.ReadJSON(raw, &propsScheduleTrigger); err != nil {
		return err
	}
	m.Schedules = propsScheduleTrigger.Schedules

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ScheduleTrigger) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BuildTrigger)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsScheduleTrigger struct {
		Schedules []*Schedule `json:"schedules"`
	}
	propsScheduleTrigger.Schedules = m.Schedules

	jsonDataPropsScheduleTrigger, errScheduleTrigger := swag.WriteJSON(propsScheduleTrigger)
	if errScheduleTrigger != nil {
		return nil, errScheduleTrigger
	}
	_parts = append(_parts, jsonDataPropsScheduleTrigger)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this schedule trigger
func (m *ScheduleTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuildTrigger
	if err := m.BuildTrigger.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleTrigger) validateSchedules(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ScheduleTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleTrigger) UnmarshalBinary(b []byte) error {
	var res ScheduleTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
