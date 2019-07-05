// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SupportedTrigger supported trigger
// swagger:model SupportedTrigger
type SupportedTrigger struct {

	// The default interval to wait between polls (only relevant when NotificationType is Polling).
	DefaultPollingInterval int32 `json:"defaultPollingInterval,omitempty"`

	// How the trigger is notified of changes.
	NotificationType string `json:"notificationType,omitempty"`

	// The capabilities supported by this trigger.
	SupportedCapabilities map[string]interface{} `json:"supportedCapabilities,omitempty"`

	// The type of trigger.
	// Enum: [none continuousIntegration batchedContinuousIntegration schedule gatedCheckIn batchedGatedCheckIn pullRequest buildCompletion all]
	Type interface{} `json:"type,omitempty"`
}

// Validate validates this supported trigger
func (m *SupportedTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSupportedCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// additional properties value enum
var supportedTriggerSupportedCapabilitiesValueEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["unsupported","supported","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		supportedTriggerSupportedCapabilitiesValueEnum = append(supportedTriggerSupportedCapabilitiesValueEnum, v)
	}
}

func (m *SupportedTrigger) validateSupportedCapabilitiesValueEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, supportedTriggerSupportedCapabilitiesValueEnum); err != nil {
		return err
	}
	return nil
}

func (m *SupportedTrigger) validateSupportedCapabilities(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedCapabilities) { // not required
		return nil
	}

	for k := range m.SupportedCapabilities {

		if err := m.validateSupportedCapabilitiesValueEnum("supportedCapabilities"+"."+k, "body", m.SupportedCapabilities[k]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupportedTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportedTrigger) UnmarshalBinary(b []byte) error {
	var res SupportedTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
