// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AcquisitionOperationDisallowReason acquisition operation disallow reason
// swagger:model AcquisitionOperationDisallowReason
type AcquisitionOperationDisallowReason struct {

	// User-friendly message clarifying the reason for disallowance
	Message string `json:"message,omitempty"`

	// Type of reason for disallowance - AlreadyInstalled, UnresolvedDemand, etc.
	Type string `json:"type,omitempty"`
}

// Validate validates this acquisition operation disallow reason
func (m *AcquisitionOperationDisallowReason) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AcquisitionOperationDisallowReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcquisitionOperationDisallowReason) UnmarshalBinary(b []byte) error {
	var res AcquisitionOperationDisallowReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
