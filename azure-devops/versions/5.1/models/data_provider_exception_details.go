// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DataProviderExceptionDetails data provider exception details
// swagger:model DataProviderExceptionDetails
type DataProviderExceptionDetails struct {

	// The type of the exception that was thrown.
	ExceptionType string `json:"exceptionType,omitempty"`

	// Message that is associated with the exception.
	Message string `json:"message,omitempty"`

	// The StackTrace from the exception turned into a string.
	StackTrace string `json:"stackTrace,omitempty"`
}

// Validate validates this data provider exception details
func (m *DataProviderExceptionDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataProviderExceptionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataProviderExceptionDetails) UnmarshalBinary(b []byte) error {
	var res DataProviderExceptionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
