// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestRunWithDtlEnvEvent test run with dtl env event
// swagger:model TestRunWithDtlEnvEvent
type TestRunWithDtlEnvEvent struct {
	TestRunEvent

	// configuration ids
	ConfigurationIds []int32 `json:"configurationIds"`

	// mapped test run event type
	MappedTestRunEventType string `json:"mappedTestRunEventType,omitempty"`

	// run timeout
	RunTimeout string `json:"runTimeout,omitempty"`

	// test configurations mapping
	TestConfigurationsMapping string `json:"testConfigurationsMapping,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TestRunWithDtlEnvEvent) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TestRunEvent
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TestRunEvent = aO0

	// now for regular properties
	var propsTestRunWithDtlEnvEvent struct {
		ConfigurationIds []int32 `json:"configurationIds"`

		MappedTestRunEventType string `json:"mappedTestRunEventType,omitempty"`

		RunTimeout string `json:"runTimeout,omitempty"`

		TestConfigurationsMapping string `json:"testConfigurationsMapping,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsTestRunWithDtlEnvEvent); err != nil {
		return err
	}
	m.ConfigurationIds = propsTestRunWithDtlEnvEvent.ConfigurationIds

	m.MappedTestRunEventType = propsTestRunWithDtlEnvEvent.MappedTestRunEventType

	m.RunTimeout = propsTestRunWithDtlEnvEvent.RunTimeout

	m.TestConfigurationsMapping = propsTestRunWithDtlEnvEvent.TestConfigurationsMapping

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TestRunWithDtlEnvEvent) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TestRunEvent)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTestRunWithDtlEnvEvent struct {
		ConfigurationIds []int32 `json:"configurationIds"`

		MappedTestRunEventType string `json:"mappedTestRunEventType,omitempty"`

		RunTimeout string `json:"runTimeout,omitempty"`

		TestConfigurationsMapping string `json:"testConfigurationsMapping,omitempty"`
	}
	propsTestRunWithDtlEnvEvent.ConfigurationIds = m.ConfigurationIds

	propsTestRunWithDtlEnvEvent.MappedTestRunEventType = m.MappedTestRunEventType

	propsTestRunWithDtlEnvEvent.RunTimeout = m.RunTimeout

	propsTestRunWithDtlEnvEvent.TestConfigurationsMapping = m.TestConfigurationsMapping

	jsonDataPropsTestRunWithDtlEnvEvent, errTestRunWithDtlEnvEvent := swag.WriteJSON(propsTestRunWithDtlEnvEvent)
	if errTestRunWithDtlEnvEvent != nil {
		return nil, errTestRunWithDtlEnvEvent
	}
	_parts = append(_parts, jsonDataPropsTestRunWithDtlEnvEvent)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this test run with dtl env event
func (m *TestRunWithDtlEnvEvent) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TestRunEvent
	if err := m.TestRunEvent.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TestRunWithDtlEnvEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestRunWithDtlEnvEvent) UnmarshalBinary(b []byte) error {
	var res TestRunWithDtlEnvEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
