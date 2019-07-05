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

// DeliveryViewPropertyCollection Collection of properties, specific to the DeliveryTimelineView
// swagger:model DeliveryViewPropertyCollection
type DeliveryViewPropertyCollection struct {

	// Card settings
	CardSettings *CardSettings `json:"cardSettings,omitempty"`

	// Field criteria
	Criteria []*FilterClause `json:"criteria"`

	// Markers. Will be missing/null if there are no markers.
	Markers []*Marker `json:"markers"`

	// Team backlog mappings
	TeamBacklogMappings []*TeamBacklogMapping `json:"teamBacklogMappings"`
}

// Validate validates this delivery view property collection
func (m *DeliveryViewPropertyCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamBacklogMappings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeliveryViewPropertyCollection) validateCardSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.CardSettings) { // not required
		return nil
	}

	if m.CardSettings != nil {
		if err := m.CardSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cardSettings")
			}
			return err
		}
	}

	return nil
}

func (m *DeliveryViewPropertyCollection) validateCriteria(formats strfmt.Registry) error {

	if swag.IsZero(m.Criteria) { // not required
		return nil
	}

	for i := 0; i < len(m.Criteria); i++ {
		if swag.IsZero(m.Criteria[i]) { // not required
			continue
		}

		if m.Criteria[i] != nil {
			if err := m.Criteria[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeliveryViewPropertyCollection) validateMarkers(formats strfmt.Registry) error {

	if swag.IsZero(m.Markers) { // not required
		return nil
	}

	for i := 0; i < len(m.Markers); i++ {
		if swag.IsZero(m.Markers[i]) { // not required
			continue
		}

		if m.Markers[i] != nil {
			if err := m.Markers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("markers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeliveryViewPropertyCollection) validateTeamBacklogMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamBacklogMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.TeamBacklogMappings); i++ {
		if swag.IsZero(m.TeamBacklogMappings[i]) { // not required
			continue
		}

		if m.TeamBacklogMappings[i] != nil {
			if err := m.TeamBacklogMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teamBacklogMappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeliveryViewPropertyCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeliveryViewPropertyCollection) UnmarshalBinary(b []byte) error {
	var res DeliveryViewPropertyCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}