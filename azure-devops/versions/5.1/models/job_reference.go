// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// JobReference Job in pipeline. This is related to matrixing in YAML.
// swagger:model JobReference
type JobReference struct {

	// Attempt number of the job
	Attempt int32 `json:"attempt,omitempty"`

	// Matrixing in YAML generates copies of a job with different inputs in matrix. JobName is the name of those input. Maximum supported length for name is 256 character.
	JobName string `json:"jobName,omitempty"`
}

// Validate validates this job reference
func (m *JobReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobReference) UnmarshalBinary(b []byte) error {
	var res JobReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
