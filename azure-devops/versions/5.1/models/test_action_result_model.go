// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestActionResultModel Represents a test step result.
// swagger:model TestActionResultModel
type TestActionResultModel struct {
	TestResultModelBase

	// Path identifier test step in test case workitem.
	ActionPath string `json:"actionPath,omitempty"`

	// Iteration ID of test action result.
	IterationID int32 `json:"iterationId,omitempty"`

	// Reference to shared step workitem.
	SharedStepModel *SharedStepModel `json:"sharedStepModel,omitempty"`

	// This is step Id of test case. For shared step, it is step Id of shared step in test case workitem; step Id in shared step. Example: TestCase workitem has two steps: 1) Normal step with Id = 1 2) Shared Step with Id = 2. Inside shared step: a) Normal Step with Id = 1 Value for StepIdentifier for First step: "1" Second step: "2;1"
	StepIdentifier string `json:"stepIdentifier,omitempty"`

	// Url of test action result.
	URL string `json:"url,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TestActionResultModel) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TestResultModelBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TestResultModelBase = aO0

	// now for regular properties
	var propsTestActionResultModel struct {
		ActionPath string `json:"actionPath,omitempty"`

		IterationID int32 `json:"iterationId,omitempty"`

		SharedStepModel *SharedStepModel `json:"sharedStepModel,omitempty"`

		StepIdentifier string `json:"stepIdentifier,omitempty"`

		URL string `json:"url,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsTestActionResultModel); err != nil {
		return err
	}
	m.ActionPath = propsTestActionResultModel.ActionPath

	m.IterationID = propsTestActionResultModel.IterationID

	m.SharedStepModel = propsTestActionResultModel.SharedStepModel

	m.StepIdentifier = propsTestActionResultModel.StepIdentifier

	m.URL = propsTestActionResultModel.URL

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TestActionResultModel) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TestResultModelBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTestActionResultModel struct {
		ActionPath string `json:"actionPath,omitempty"`

		IterationID int32 `json:"iterationId,omitempty"`

		SharedStepModel *SharedStepModel `json:"sharedStepModel,omitempty"`

		StepIdentifier string `json:"stepIdentifier,omitempty"`

		URL string `json:"url,omitempty"`
	}
	propsTestActionResultModel.ActionPath = m.ActionPath

	propsTestActionResultModel.IterationID = m.IterationID

	propsTestActionResultModel.SharedStepModel = m.SharedStepModel

	propsTestActionResultModel.StepIdentifier = m.StepIdentifier

	propsTestActionResultModel.URL = m.URL

	jsonDataPropsTestActionResultModel, errTestActionResultModel := swag.WriteJSON(propsTestActionResultModel)
	if errTestActionResultModel != nil {
		return nil, errTestActionResultModel
	}
	_parts = append(_parts, jsonDataPropsTestActionResultModel)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this test action result model
func (m *TestActionResultModel) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TestResultModelBase
	if err := m.TestResultModelBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedStepModel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestActionResultModel) validateSharedStepModel(formats strfmt.Registry) error {

	if swag.IsZero(m.SharedStepModel) { // not required
		return nil
	}

	if m.SharedStepModel != nil {
		if err := m.SharedStepModel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedStepModel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestActionResultModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestActionResultModel) UnmarshalBinary(b []byte) error {
	var res TestActionResultModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}