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

// ResourceUsage resource usage
// swagger:model ResourceUsage
type ResourceUsage struct {

	// resource limit
	ResourceLimit *ResourceLimit `json:"resourceLimit,omitempty"`

	// running requests
	RunningRequests []*TaskAgentJobRequest `json:"runningRequests"`

	// used count
	UsedCount int32 `json:"usedCount,omitempty"`

	// used minutes
	UsedMinutes int32 `json:"usedMinutes,omitempty"`
}

// Validate validates this resource usage
func (m *ResourceUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunningRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceUsage) validateResourceLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceLimit) { // not required
		return nil
	}

	if m.ResourceLimit != nil {
		if err := m.ResourceLimit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceLimit")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceUsage) validateRunningRequests(formats strfmt.Registry) error {

	if swag.IsZero(m.RunningRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.RunningRequests); i++ {
		if swag.IsZero(m.RunningRequests[i]) { // not required
			continue
		}

		if m.RunningRequests[i] != nil {
			if err := m.RunningRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runningRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceUsage) UnmarshalBinary(b []byte) error {
	var res ResourceUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
