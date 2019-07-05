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

// ServiceEndpointExecutionData Represents service endpoint execution data.
// swagger:model ServiceEndpointExecutionData
type ServiceEndpointExecutionData struct {

	// Gets the definition of service endpoint execution owner.
	Definition *TaskOrchestrationOwner `json:"definition,omitempty"`

	// Gets the finish time of service endpoint execution.
	// Format: date-time
	FinishTime strfmt.DateTime `json:"finishTime,omitempty"`

	// Gets the Id of service endpoint execution data.
	ID int64 `json:"id,omitempty"`

	// Gets the owner of service endpoint execution data.
	Owner *TaskOrchestrationOwner `json:"owner,omitempty"`

	// Gets the plan type of service endpoint execution data.
	PlanType string `json:"planType,omitempty"`

	// Gets the result of service endpoint execution.
	// Enum: [succeeded succeededWithIssues failed canceled skipped abandoned]
	Result interface{} `json:"result,omitempty"`

	// Gets the start time of service endpoint execution.
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`
}

// Validate validates this service endpoint execution data
func (m *ServiceEndpointExecutionData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceEndpointExecutionData) validateDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.Definition) { // not required
		return nil
	}

	if m.Definition != nil {
		if err := m.Definition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("definition")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceEndpointExecutionData) validateFinishTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("finishTime", "body", "date-time", m.FinishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceEndpointExecutionData) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceEndpointExecutionData) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceEndpointExecutionData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceEndpointExecutionData) UnmarshalBinary(b []byte) error {
	var res ServiceEndpointExecutionData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
