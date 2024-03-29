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

// TaskEvent task event
// swagger:model TaskEvent
type TaskEvent struct {
	JobEvent

	// task Id
	// Format: uuid
	TaskID strfmt.UUID `json:"taskId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TaskEvent) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 JobEvent
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.JobEvent = aO0

	// now for regular properties
	var propsTaskEvent struct {
		TaskID strfmt.UUID `json:"taskId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsTaskEvent); err != nil {
		return err
	}
	m.TaskID = propsTaskEvent.TaskID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TaskEvent) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.JobEvent)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTaskEvent struct {
		TaskID strfmt.UUID `json:"taskId,omitempty"`
	}
	propsTaskEvent.TaskID = m.TaskID

	jsonDataPropsTaskEvent, errTaskEvent := swag.WriteJSON(propsTaskEvent)
	if errTaskEvent != nil {
		return nil, errTaskEvent
	}
	_parts = append(_parts, jsonDataPropsTaskEvent)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this task event
func (m *TaskEvent) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with JobEvent
	if err := m.JobEvent.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskEvent) validateTaskID(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskID) { // not required
		return nil
	}

	if err := validate.FormatOf("taskId", "body", "uuid", m.TaskID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskEvent) UnmarshalBinary(b []byte) error {
	var res TaskEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
