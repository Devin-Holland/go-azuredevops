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

// JobEvent job event
// swagger:model JobEvent
type JobEvent struct {

	// job Id
	// Format: uuid
	JobID strfmt.UUID `json:"jobId,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this job event
func (m *JobEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobEvent) validateJobID(formats strfmt.Registry) error {

	if swag.IsZero(m.JobID) { // not required
		return nil
	}

	if err := validate.FormatOf("jobId", "body", "uuid", m.JobID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobEvent) UnmarshalBinary(b []byte) error {
	var res JobEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
