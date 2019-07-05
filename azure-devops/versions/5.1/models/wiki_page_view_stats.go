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

// WikiPageViewStats Defines properties for wiki page view stats.
// swagger:model WikiPageViewStats
type WikiPageViewStats struct {

	// Wiki page view count.
	Count int32 `json:"count,omitempty"`

	// Wiki page last viewed time.
	// Format: date-time
	LastViewedTime strfmt.DateTime `json:"lastViewedTime,omitempty"`

	// Wiki page path.
	Path string `json:"path,omitempty"`
}

// Validate validates this wiki page view stats
func (m *WikiPageViewStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastViewedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WikiPageViewStats) validateLastViewedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastViewedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastViewedTime", "body", "date-time", m.LastViewedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WikiPageViewStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WikiPageViewStats) UnmarshalBinary(b []byte) error {
	var res WikiPageViewStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
