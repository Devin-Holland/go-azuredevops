// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RuleAction Action to take when the rule is triggered.
// swagger:model RuleAction
type RuleAction struct {

	// Type of action to take when the rule is triggered.
	// Enum: [makeRequired makeReadOnly setDefaultValue setDefaultFromClock setDefaultFromCurrentUser setDefaultFromField copyValue copyFromClock copyFromCurrentUser copyFromField setValueToEmpty copyFromServerClock copyFromServerCurrentUser]
	ActionType interface{} `json:"actionType,omitempty"`

	// Field on which the action should be taken.
	TargetField string `json:"targetField,omitempty"`

	// Value to apply on target field, once the action is taken.
	Value string `json:"value,omitempty"`
}

// Validate validates this rule action
func (m *RuleAction) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuleAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleAction) UnmarshalBinary(b []byte) error {
	var res RuleAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
