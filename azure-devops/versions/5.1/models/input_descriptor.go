// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InputDescriptor Describes an input for subscriptions.
// swagger:model InputDescriptor
type InputDescriptor struct {

	// The ids of all inputs that the value of this input is dependent on.
	DependencyInputIds []string `json:"dependencyInputIds"`

	// Description of what this input is used for
	Description string `json:"description,omitempty"`

	// The group localized name to which this input belongs and can be shown as a header for the container that will include all the inputs in the group.
	GroupName string `json:"groupName,omitempty"`

	// If true, the value information for this input is dynamic and should be fetched when the value of dependency inputs change.
	HasDynamicValueInformation bool `json:"hasDynamicValueInformation,omitempty"`

	// Identifier for the subscription input
	ID string `json:"id,omitempty"`

	// Mode in which the value of this input should be entered
	// Enum: [none textBox passwordBox combo radioButtons checkBox textArea]
	InputMode interface{} `json:"inputMode,omitempty"`

	// Gets whether this input is confidential, such as for a password or application key
	IsConfidential bool `json:"isConfidential,omitempty"`

	// Localized name which can be shown as a label for the subscription input
	Name string `json:"name,omitempty"`

	// Custom properties for the input which can be used by the service provider
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Underlying data type for the input value. When this value is specified, InputMode, Validation and Values are optional.
	Type string `json:"type,omitempty"`

	// Gets whether this input is included in the default generated action description.
	UseInDefaultDescription bool `json:"useInDefaultDescription,omitempty"`

	// Information to use to validate this input's value
	Validation *InputValidation `json:"validation,omitempty"`

	// A hint for input value. It can be used in the UI as the input placeholder.
	ValueHint string `json:"valueHint,omitempty"`

	// Information about possible values for this input
	Values *InputValues `json:"values,omitempty"`
}

// Validate validates this input descriptor
func (m *InputDescriptor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InputDescriptor) validateValidation(formats strfmt.Registry) error {

	if swag.IsZero(m.Validation) { // not required
		return nil
	}

	if m.Validation != nil {
		if err := m.Validation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validation")
			}
			return err
		}
	}

	return nil
}

func (m *InputDescriptor) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	if m.Values != nil {
		if err := m.Values.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("values")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InputDescriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InputDescriptor) UnmarshalBinary(b []byte) error {
	var res InputDescriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
