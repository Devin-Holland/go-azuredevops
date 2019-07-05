// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BuildTagsAddedEvent build tags added event
// swagger:model BuildTagsAddedEvent
type BuildTagsAddedEvent struct {
	BuildUpdatedEvent

	// all tags
	AllTags []string `json:"allTags"`

	// new tags
	NewTags []string `json:"newTags"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BuildTagsAddedEvent) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuildUpdatedEvent
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuildUpdatedEvent = aO0

	// now for regular properties
	var propsBuildTagsAddedEvent struct {
		AllTags []string `json:"allTags"`

		NewTags []string `json:"newTags"`
	}
	if err := swag.ReadJSON(raw, &propsBuildTagsAddedEvent); err != nil {
		return err
	}
	m.AllTags = propsBuildTagsAddedEvent.AllTags

	m.NewTags = propsBuildTagsAddedEvent.NewTags

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BuildTagsAddedEvent) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BuildUpdatedEvent)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsBuildTagsAddedEvent struct {
		AllTags []string `json:"allTags"`

		NewTags []string `json:"newTags"`
	}
	propsBuildTagsAddedEvent.AllTags = m.AllTags

	propsBuildTagsAddedEvent.NewTags = m.NewTags

	jsonDataPropsBuildTagsAddedEvent, errBuildTagsAddedEvent := swag.WriteJSON(propsBuildTagsAddedEvent)
	if errBuildTagsAddedEvent != nil {
		return nil, errBuildTagsAddedEvent
	}
	_parts = append(_parts, jsonDataPropsBuildTagsAddedEvent)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this build tags added event
func (m *BuildTagsAddedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuildUpdatedEvent
	if err := m.BuildUpdatedEvent.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BuildTagsAddedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildTagsAddedEvent) UnmarshalBinary(b []byte) error {
	var res BuildTagsAddedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}