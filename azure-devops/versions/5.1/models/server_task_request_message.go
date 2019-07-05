// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerTaskRequestMessage server task request message
// swagger:model ServerTaskRequestMessage
type ServerTaskRequestMessage struct {
	JobRequestMessage

	// task definition
	TaskDefinition *TaskDefinition `json:"taskDefinition,omitempty"`

	// task instance
	TaskInstance *TaskInstance `json:"taskInstance,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ServerTaskRequestMessage) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 JobRequestMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.JobRequestMessage = aO0

	// now for regular properties
	var propsServerTaskRequestMessage struct {
		TaskDefinition *TaskDefinition `json:"taskDefinition,omitempty"`

		TaskInstance *TaskInstance `json:"taskInstance,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsServerTaskRequestMessage); err != nil {
		return err
	}
	m.TaskDefinition = propsServerTaskRequestMessage.TaskDefinition

	m.TaskInstance = propsServerTaskRequestMessage.TaskInstance

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ServerTaskRequestMessage) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.JobRequestMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsServerTaskRequestMessage struct {
		TaskDefinition *TaskDefinition `json:"taskDefinition,omitempty"`

		TaskInstance *TaskInstance `json:"taskInstance,omitempty"`
	}
	propsServerTaskRequestMessage.TaskDefinition = m.TaskDefinition

	propsServerTaskRequestMessage.TaskInstance = m.TaskInstance

	jsonDataPropsServerTaskRequestMessage, errServerTaskRequestMessage := swag.WriteJSON(propsServerTaskRequestMessage)
	if errServerTaskRequestMessage != nil {
		return nil, errServerTaskRequestMessage
	}
	_parts = append(_parts, jsonDataPropsServerTaskRequestMessage)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this server task request message
func (m *ServerTaskRequestMessage) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with JobRequestMessage
	if err := m.JobRequestMessage.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerTaskRequestMessage) validateTaskDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskDefinition) { // not required
		return nil
	}

	if m.TaskDefinition != nil {
		if err := m.TaskDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskDefinition")
			}
			return err
		}
	}

	return nil
}

func (m *ServerTaskRequestMessage) validateTaskInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskInstance) { // not required
		return nil
	}

	if m.TaskInstance != nil {
		if err := m.TaskInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerTaskRequestMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerTaskRequestMessage) UnmarshalBinary(b []byte) error {
	var res ServerTaskRequestMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
