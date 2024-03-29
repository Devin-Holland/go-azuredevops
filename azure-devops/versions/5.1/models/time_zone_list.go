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

// TimeZoneList time zone list
// swagger:model TimeZoneList
type TimeZoneList struct {

	// UTC timezone.
	UtcTimeZone *TimeZone `json:"utcTimeZone,omitempty"`

	// List of valid timezones.
	ValidTimeZones []*TimeZone `json:"validTimeZones"`
}

// Validate validates this time zone list
func (m *TimeZoneList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUtcTimeZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTimeZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeZoneList) validateUtcTimeZone(formats strfmt.Registry) error {

	if swag.IsZero(m.UtcTimeZone) { // not required
		return nil
	}

	if m.UtcTimeZone != nil {
		if err := m.UtcTimeZone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utcTimeZone")
			}
			return err
		}
	}

	return nil
}

func (m *TimeZoneList) validateValidTimeZones(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidTimeZones) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidTimeZones); i++ {
		if swag.IsZero(m.ValidTimeZones[i]) { // not required
			continue
		}

		if m.ValidTimeZones[i] != nil {
			if err := m.ValidTimeZones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validTimeZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeZoneList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeZoneList) UnmarshalBinary(b []byte) error {
	var res TimeZoneList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
