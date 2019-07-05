// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CounterSampleQueryDetails counter sample query details
// swagger:model CounterSampleQueryDetails
type CounterSampleQueryDetails struct {

	// counter instance Id
	CounterInstanceID string `json:"counterInstanceId,omitempty"`

	// from interval
	FromInterval int32 `json:"fromInterval,omitempty"`

	// to interval
	ToInterval int32 `json:"toInterval,omitempty"`
}

// Validate validates this counter sample query details
func (m *CounterSampleQueryDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CounterSampleQueryDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CounterSampleQueryDetails) UnmarshalBinary(b []byte) error {
	var res CounterSampleQueryDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
