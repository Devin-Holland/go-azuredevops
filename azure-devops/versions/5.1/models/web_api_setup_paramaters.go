// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WebAPISetupParamaters web Api setup paramaters
// swagger:model WebApiSetupParamaters
type WebAPISetupParamaters struct {

	// configurations
	Configurations map[string]string `json:"configurations,omitempty"`
}

// Validate validates this web Api setup paramaters
func (m *WebAPISetupParamaters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebAPISetupParamaters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebAPISetupParamaters) UnmarshalBinary(b []byte) error {
	var res WebAPISetupParamaters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
