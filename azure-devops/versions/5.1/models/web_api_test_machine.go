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

// WebAPITestMachine web Api test machine
// swagger:model WebApiTestMachine
type WebAPITestMachine struct {

	// last heart beat
	// Format: date-time
	LastHeartBeat strfmt.DateTime `json:"lastHeartBeat,omitempty"`

	// machine name
	MachineName string `json:"machineName,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this web Api test machine
func (m *WebAPITestMachine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastHeartBeat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebAPITestMachine) validateLastHeartBeat(formats strfmt.Registry) error {

	if swag.IsZero(m.LastHeartBeat) { // not required
		return nil
	}

	if err := validate.FormatOf("lastHeartBeat", "body", "date-time", m.LastHeartBeat.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebAPITestMachine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebAPITestMachine) UnmarshalBinary(b []byte) error {
	var res WebAPITestMachine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
