// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestResultMetaDataUpdateResponse test result meta data update response
// swagger:model TestResultMetaDataUpdateResponse
type TestResultMetaDataUpdateResponse struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this test result meta data update response
func (m *TestResultMetaDataUpdateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestResultMetaDataUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultMetaDataUpdateResponse) UnmarshalBinary(b []byte) error {
	var res TestResultMetaDataUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
