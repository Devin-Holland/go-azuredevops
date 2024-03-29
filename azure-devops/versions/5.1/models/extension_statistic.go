// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ExtensionStatistic extension statistic
// swagger:model ExtensionStatistic
type ExtensionStatistic struct {

	// statistic name
	StatisticName string `json:"statisticName,omitempty"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this extension statistic
func (m *ExtensionStatistic) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionStatistic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionStatistic) UnmarshalBinary(b []byte) error {
	var res ExtensionStatistic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
