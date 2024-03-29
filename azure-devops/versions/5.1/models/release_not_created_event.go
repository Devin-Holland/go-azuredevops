// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReleaseNotCreatedEvent release not created event
// swagger:model ReleaseNotCreatedEvent
type ReleaseNotCreatedEvent struct {

	// definition reference
	DefinitionReference *ReleaseDefinitionShallowReference `json:"definitionReference,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// release reason
	// Enum: [none manual continuousIntegration schedule pullRequest]
	ReleaseReason interface{} `json:"releaseReason,omitempty"`

	// requested by
	RequestedBy *IdentityRef `json:"requestedBy,omitempty"`
}

// Validate validates this release not created event
func (m *ReleaseNotCreatedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefinitionReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseNotCreatedEvent) validateDefinitionReference(formats strfmt.Registry) error {

	if swag.IsZero(m.DefinitionReference) { // not required
		return nil
	}

	if m.DefinitionReference != nil {
		if err := m.DefinitionReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("definitionReference")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseNotCreatedEvent) validateRequestedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedBy) { // not required
		return nil
	}

	if m.RequestedBy != nil {
		if err := m.RequestedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedBy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseNotCreatedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseNotCreatedEvent) UnmarshalBinary(b []byte) error {
	var res ReleaseNotCreatedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
