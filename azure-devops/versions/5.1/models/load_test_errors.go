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

// LoadTestErrors load test errors
// swagger:model LoadTestErrors
type LoadTestErrors struct {

	// count
	Count int32 `json:"count,omitempty"`

	// occurrences
	Occurrences int32 `json:"occurrences,omitempty"`

	// types
	Types []*Type `json:"types"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this load test errors
func (m *LoadTestErrors) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadTestErrors) validateTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.Types) { // not required
		return nil
	}

	for i := 0; i < len(m.Types); i++ {
		if swag.IsZero(m.Types[i]) { // not required
			continue
		}

		if m.Types[i] != nil {
			if err := m.Types[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadTestErrors) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadTestErrors) UnmarshalBinary(b []byte) error {
	var res LoadTestErrors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
