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

// CounterInstanceSamples counter instance samples
// swagger:model CounterInstanceSamples
type CounterInstanceSamples struct {

	// count
	Count int32 `json:"count,omitempty"`

	// counter instance Id
	CounterInstanceID string `json:"counterInstanceId,omitempty"`

	// next refresh time
	// Format: date-time
	NextRefreshTime strfmt.DateTime `json:"nextRefreshTime,omitempty"`

	// values
	Values []*CounterSample `json:"values"`
}

// Validate validates this counter instance samples
func (m *CounterInstanceSamples) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNextRefreshTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CounterInstanceSamples) validateNextRefreshTime(formats strfmt.Registry) error {

	if swag.IsZero(m.NextRefreshTime) { // not required
		return nil
	}

	if err := validate.FormatOf("nextRefreshTime", "body", "date-time", m.NextRefreshTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CounterInstanceSamples) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CounterInstanceSamples) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CounterInstanceSamples) UnmarshalBinary(b []byte) error {
	var res CounterInstanceSamples
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
