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

// CounterSample counter sample
// swagger:model CounterSample
type CounterSample struct {

	// base value
	BaseValue int64 `json:"baseValue,omitempty"`

	// computed value
	ComputedValue string `json:"computedValue,omitempty"`

	// counter frequency
	CounterFrequency int64 `json:"counterFrequency,omitempty"`

	// counter instance Id
	CounterInstanceID string `json:"counterInstanceId,omitempty"`

	// counter type
	CounterType string `json:"counterType,omitempty"`

	// interval end date
	// Format: date-time
	IntervalEndDate strfmt.DateTime `json:"intervalEndDate,omitempty"`

	// interval number
	IntervalNumber int32 `json:"intervalNumber,omitempty"`

	// raw value
	RawValue int64 `json:"rawValue,omitempty"`

	// system frequency
	SystemFrequency int64 `json:"systemFrequency,omitempty"`

	// time stamp
	TimeStamp int64 `json:"timeStamp,omitempty"`
}

// Validate validates this counter sample
func (m *CounterSample) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntervalEndDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CounterSample) validateIntervalEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.IntervalEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("intervalEndDate", "body", "date-time", m.IntervalEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CounterSample) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CounterSample) UnmarshalBinary(b []byte) error {
	var res CounterSample
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
