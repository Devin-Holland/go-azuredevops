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

// ReleaseSchedule release schedule
// swagger:model ReleaseSchedule
type ReleaseSchedule struct {

	// Days of the week to release.
	// Enum: [none monday tuesday wednesday thursday friday saturday sunday all]
	DaysToRelease interface{} `json:"daysToRelease,omitempty"`

	// Team Foundation Job Definition Job Id.
	// Format: uuid
	JobID strfmt.UUID `json:"jobId,omitempty"`

	// Flag to determine if this schedule should only release if the associated artifact has been changed or release definition changed.
	ScheduleOnlyWithChanges bool `json:"scheduleOnlyWithChanges,omitempty"`

	// Local time zone hour to start.
	StartHours int32 `json:"startHours,omitempty"`

	// Local time zone minute to start.
	StartMinutes int32 `json:"startMinutes,omitempty"`

	// Time zone Id of release schedule, such as 'UTC'.
	TimeZoneID string `json:"timeZoneId,omitempty"`
}

// Validate validates this release schedule
func (m *ReleaseSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseSchedule) validateJobID(formats strfmt.Registry) error {

	if swag.IsZero(m.JobID) { // not required
		return nil
	}

	if err := validate.FormatOf("jobId", "body", "uuid", m.JobID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseSchedule) UnmarshalBinary(b []byte) error {
	var res ReleaseSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}