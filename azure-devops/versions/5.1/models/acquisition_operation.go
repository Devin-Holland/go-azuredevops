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

// AcquisitionOperation acquisition operation
// swagger:model AcquisitionOperation
type AcquisitionOperation struct {

	// State of the the AcquisitionOperation for the current user
	// Enum: [disallow allow completed]
	OperationState interface{} `json:"operationState,omitempty"`

	// AcquisitionOperationType: install, request, buy, etc...
	// Enum: [get install buy try request none purchaseRequest]
	OperationType interface{} `json:"operationType,omitempty"`

	// Optional reason to justify current state. Typically used with Disallow state.
	Reason string `json:"reason,omitempty"`

	// List of reasons indicating why the operation is not allowed.
	Reasons []*AcquisitionOperationDisallowReason `json:"reasons"`
}

// Validate validates this acquisition operation
func (m *AcquisitionOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReasons(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcquisitionOperation) validateReasons(formats strfmt.Registry) error {

	if swag.IsZero(m.Reasons) { // not required
		return nil
	}

	for i := 0; i < len(m.Reasons); i++ {
		if swag.IsZero(m.Reasons[i]) { // not required
			continue
		}

		if m.Reasons[i] != nil {
			if err := m.Reasons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AcquisitionOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcquisitionOperation) UnmarshalBinary(b []byte) error {
	var res AcquisitionOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
