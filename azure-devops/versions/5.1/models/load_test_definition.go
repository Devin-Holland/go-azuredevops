// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LoadTestDefinition load test definition
// swagger:model LoadTestDefinition
type LoadTestDefinition struct {

	// agent count
	AgentCount int32 `json:"agentCount,omitempty"`

	// browser mixs
	BrowserMixs []*BrowserMix `json:"browserMixs"`

	// core count
	CoreCount int32 `json:"coreCount,omitempty"`

	// cores per agent
	CoresPerAgent int32 `json:"coresPerAgent,omitempty"`

	// load generation geo locations
	LoadGenerationGeoLocations []*LoadGenerationGeoLocation `json:"loadGenerationGeoLocations"`

	// load pattern name
	LoadPatternName string `json:"loadPatternName,omitempty"`

	// load test name
	LoadTestName string `json:"loadTestName,omitempty"`

	// max vusers
	MaxVusers int32 `json:"maxVusers,omitempty"`

	// run duration
	RunDuration int32 `json:"runDuration,omitempty"`

	// sampling rate
	SamplingRate int32 `json:"samplingRate,omitempty"`

	// think time
	ThinkTime int32 `json:"thinkTime,omitempty"`

	// urls
	Urls []string `json:"urls"`
}

// Validate validates this load test definition
func (m *LoadTestDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrowserMixs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadGenerationGeoLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadTestDefinition) validateBrowserMixs(formats strfmt.Registry) error {

	if swag.IsZero(m.BrowserMixs) { // not required
		return nil
	}

	for i := 0; i < len(m.BrowserMixs); i++ {
		if swag.IsZero(m.BrowserMixs[i]) { // not required
			continue
		}

		if m.BrowserMixs[i] != nil {
			if err := m.BrowserMixs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("browserMixs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoadTestDefinition) validateLoadGenerationGeoLocations(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadGenerationGeoLocations) { // not required
		return nil
	}

	for i := 0; i < len(m.LoadGenerationGeoLocations); i++ {
		if swag.IsZero(m.LoadGenerationGeoLocations[i]) { // not required
			continue
		}

		if m.LoadGenerationGeoLocations[i] != nil {
			if err := m.LoadGenerationGeoLocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("loadGenerationGeoLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadTestDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadTestDefinition) UnmarshalBinary(b []byte) error {
	var res LoadTestDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
