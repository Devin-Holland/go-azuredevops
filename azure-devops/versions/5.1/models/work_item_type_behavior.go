// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkItemTypeBehavior Association between a work item type and it's behavior
// swagger:model WorkItemTypeBehavior
type WorkItemTypeBehavior struct {

	// Reference to the behavior of a work item type
	Behavior *WorkItemBehaviorReference `json:"behavior,omitempty"`

	// If true the work item type is the default work item type in the behavior
	IsDefault bool `json:"isDefault,omitempty"`

	// URL of the work item type behavior
	URL string `json:"url,omitempty"`
}

// Validate validates this work item type behavior
func (m *WorkItemTypeBehavior) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBehavior(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkItemTypeBehavior) validateBehavior(formats strfmt.Registry) error {

	if swag.IsZero(m.Behavior) { // not required
		return nil
	}

	if m.Behavior != nil {
		if err := m.Behavior.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("behavior")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemTypeBehavior) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemTypeBehavior) UnmarshalBinary(b []byte) error {
	var res WorkItemTypeBehavior
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
