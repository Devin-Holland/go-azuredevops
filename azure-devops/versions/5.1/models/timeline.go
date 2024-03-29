// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Timeline timeline
// swagger:model Timeline
type Timeline struct {
	TimelineReference

	// last changed by
	// Format: uuid
	LastChangedBy strfmt.UUID `json:"lastChangedBy,omitempty"`

	// last changed on
	// Format: date-time
	LastChangedOn strfmt.DateTime `json:"lastChangedOn,omitempty"`

	// records
	Records []*TimelineRecord `json:"records"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Timeline) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TimelineReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TimelineReference = aO0

	// now for regular properties
	var propsTimeline struct {
		LastChangedBy strfmt.UUID `json:"lastChangedBy,omitempty"`

		LastChangedOn strfmt.DateTime `json:"lastChangedOn,omitempty"`

		Records []*TimelineRecord `json:"records"`
	}
	if err := swag.ReadJSON(raw, &propsTimeline); err != nil {
		return err
	}
	m.LastChangedBy = propsTimeline.LastChangedBy

	m.LastChangedOn = propsTimeline.LastChangedOn

	m.Records = propsTimeline.Records

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Timeline) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TimelineReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTimeline struct {
		LastChangedBy strfmt.UUID `json:"lastChangedBy,omitempty"`

		LastChangedOn strfmt.DateTime `json:"lastChangedOn,omitempty"`

		Records []*TimelineRecord `json:"records"`
	}
	propsTimeline.LastChangedBy = m.LastChangedBy

	propsTimeline.LastChangedOn = m.LastChangedOn

	propsTimeline.Records = m.Records

	jsonDataPropsTimeline, errTimeline := swag.WriteJSON(propsTimeline)
	if errTimeline != nil {
		return nil, errTimeline
	}
	_parts = append(_parts, jsonDataPropsTimeline)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this timeline
func (m *Timeline) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TimelineReference
	if err := m.TimelineReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastChangedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastChangedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Timeline) validateLastChangedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastChangedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("lastChangedBy", "body", "uuid", m.LastChangedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Timeline) validateLastChangedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastChangedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("lastChangedOn", "body", "date-time", m.LastChangedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Timeline) validateRecords(formats strfmt.Registry) error {

	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Timeline) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Timeline) UnmarshalBinary(b []byte) error {
	var res Timeline
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
