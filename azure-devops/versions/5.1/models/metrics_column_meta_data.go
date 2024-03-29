// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MetricsColumnMetaData Meta data for a metrics column.
// swagger:model MetricsColumnMetaData
type MetricsColumnMetaData struct {

	// Name.
	ColumnName string `json:"columnName,omitempty"`

	// Data type.
	ColumnValueType string `json:"columnValueType,omitempty"`
}

// Validate validates this metrics column meta data
func (m *MetricsColumnMetaData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsColumnMetaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsColumnMetaData) UnmarshalBinary(b []byte) error {
	var res MetricsColumnMetaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
