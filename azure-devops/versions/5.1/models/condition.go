// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Condition condition
// swagger:model Condition
type Condition struct {

	// Gets or sets the condition type.
	// Enum: [undefined event environmentState artifact]
	ConditionType interface{} `json:"conditionType,omitempty"`

	// Gets or sets the name of the condition. e.g. 'ReleaseStarted'.
	Name string `json:"name,omitempty"`

	// Gets or set value of the condition.
	Value string `json:"value,omitempty"`
}

// Validate validates this condition
func (m *Condition) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Condition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Condition) UnmarshalBinary(b []byte) error {
	var res Condition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
