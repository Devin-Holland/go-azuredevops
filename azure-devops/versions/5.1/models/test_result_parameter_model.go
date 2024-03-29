// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestResultParameterModel Test parameter information in a test iteration.
// swagger:model TestResultParameterModel
type TestResultParameterModel struct {

	// Test step path where parameter is referenced.
	ActionPath string `json:"actionPath,omitempty"`

	// Iteration ID.
	IterationID int32 `json:"iterationId,omitempty"`

	// Name of parameter.
	ParameterName string `json:"parameterName,omitempty"`

	// This is step Id of test case. For shared step, it is step Id of shared step in test case workitem; step Id in shared step. Example: TestCase workitem has two steps: 1) Normal step with Id = 1 2) Shared Step with Id = 2. Inside shared step: a) Normal Step with Id = 1 Value for StepIdentifier for First step: "1" Second step: "2;1"
	StepIdentifier string `json:"stepIdentifier,omitempty"`

	// Url of test parameter.
	URL string `json:"url,omitempty"`

	// Value of parameter.
	Value string `json:"value,omitempty"`
}

// Validate validates this test result parameter model
func (m *TestResultParameterModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestResultParameterModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultParameterModel) UnmarshalBinary(b []byte) error {
	var res TestResultParameterModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
