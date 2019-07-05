// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinationTestPlanCloneParams Destination Test Plan create parameters
// swagger:model DestinationTestPlanCloneParams
type DestinationTestPlanCloneParams struct {
	TestPlanCreateParams

	// Destination Project Name
	Project string `json:"project,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DestinationTestPlanCloneParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TestPlanCreateParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TestPlanCreateParams = aO0

	// now for regular properties
	var propsDestinationTestPlanCloneParams struct {
		Project string `json:"project,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsDestinationTestPlanCloneParams); err != nil {
		return err
	}
	m.Project = propsDestinationTestPlanCloneParams.Project

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DestinationTestPlanCloneParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TestPlanCreateParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsDestinationTestPlanCloneParams struct {
		Project string `json:"project,omitempty"`
	}
	propsDestinationTestPlanCloneParams.Project = m.Project

	jsonDataPropsDestinationTestPlanCloneParams, errDestinationTestPlanCloneParams := swag.WriteJSON(propsDestinationTestPlanCloneParams)
	if errDestinationTestPlanCloneParams != nil {
		return nil, errDestinationTestPlanCloneParams
	}
	_parts = append(_parts, jsonDataPropsDestinationTestPlanCloneParams)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this destination test plan clone params
func (m *DestinationTestPlanCloneParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TestPlanCreateParams
	if err := m.TestPlanCreateParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinationTestPlanCloneParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinationTestPlanCloneParams) UnmarshalBinary(b []byte) error {
	var res DestinationTestPlanCloneParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
