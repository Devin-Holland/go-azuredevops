// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestRunCounterInstance test run counter instance
// swagger:model TestRunCounterInstance
type TestRunCounterInstance struct {

	// category name
	CategoryName string `json:"categoryName,omitempty"`

	// counter instance Id
	CounterInstanceID string `json:"counterInstanceId,omitempty"`

	// counter name
	CounterName string `json:"counterName,omitempty"`

	// counter units
	CounterUnits string `json:"counterUnits,omitempty"`

	// instance name
	InstanceName string `json:"instanceName,omitempty"`

	// is preselected counter
	IsPreselectedCounter bool `json:"isPreselectedCounter,omitempty"`

	// machine name
	MachineName string `json:"machineName,omitempty"`

	// part of counter groups
	PartOfCounterGroups []string `json:"partOfCounterGroups"`

	// summary data
	SummaryData *WebInstanceSummaryData `json:"summaryData,omitempty"`

	// unique name
	UniqueName string `json:"uniqueName,omitempty"`
}

// Validate validates this test run counter instance
func (m *TestRunCounterInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummaryData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestRunCounterInstance) validateSummaryData(formats strfmt.Registry) error {

	if swag.IsZero(m.SummaryData) { // not required
		return nil
	}

	if m.SummaryData != nil {
		if err := m.SummaryData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestRunCounterInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestRunCounterInstance) UnmarshalBinary(b []byte) error {
	var res TestRunCounterInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}