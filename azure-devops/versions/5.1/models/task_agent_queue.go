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

// TaskAgentQueue An agent queue.
// swagger:model TaskAgentQueue
type TaskAgentQueue struct {

	// ID of the queue
	ID int32 `json:"id,omitempty"`

	// Name of the queue
	Name string `json:"name,omitempty"`

	// Pool reference for this queue
	Pool *TaskAgentPoolReference `json:"pool,omitempty"`

	// Project ID
	// Format: uuid
	ProjectID strfmt.UUID `json:"projectId,omitempty"`
}

// Validate validates this task agent queue
func (m *TaskAgentQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAgentQueue) validatePool(formats strfmt.Registry) error {

	if swag.IsZero(m.Pool) { // not required
		return nil
	}

	if m.Pool != nil {
		if err := m.Pool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pool")
			}
			return err
		}
	}

	return nil
}

func (m *TaskAgentQueue) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectId", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAgentQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAgentQueue) UnmarshalBinary(b []byte) error {
	var res TaskAgentQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
