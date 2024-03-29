// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestLogStoreEndpointDetails Represents Test Log store endpoint details.
// swagger:model TestLogStoreEndpointDetails
type TestLogStoreEndpointDetails struct {

	// Test log store connection Uri.
	EndpointSASURI string `json:"endpointSASUri,omitempty"`

	// Test log store endpoint type.
	// Enum: [root file]
	EndpointType interface{} `json:"endpointType,omitempty"`

	// Test log store status code
	// Enum: [success failed fileAlreadyExists invalidInput invalidFileName invalidContainer transferFailed featureDisabled buildDoesNotExist runDoesNotExist containerNotCreated apiNotSupported fileSizeExceeds containerNotFound fileNotFound directoryNotFound]
	Status interface{} `json:"status,omitempty"`
}

// Validate validates this test log store endpoint details
func (m *TestLogStoreEndpointDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestLogStoreEndpointDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestLogStoreEndpointDetails) UnmarshalBinary(b []byte) error {
	var res TestLogStoreEndpointDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
