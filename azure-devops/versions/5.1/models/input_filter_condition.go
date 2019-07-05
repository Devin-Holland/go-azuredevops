// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InputFilterCondition An expression which can be applied to filter a list of subscription inputs
// swagger:model InputFilterCondition
type InputFilterCondition struct {

	// Whether or not to do a case sensitive match
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// The Id of the input to filter on
	InputID string `json:"inputId,omitempty"`

	// The "expected" input value to compare with the actual input value
	InputValue string `json:"inputValue,omitempty"`

	// The operator applied between the expected and actual input value
	// Enum: [equals notEquals]
	Operator interface{} `json:"operator,omitempty"`
}

// Validate validates this input filter condition
func (m *InputFilterCondition) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InputFilterCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InputFilterCondition) UnmarshalBinary(b []byte) error {
	var res InputFilterCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
