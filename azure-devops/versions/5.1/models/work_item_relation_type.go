// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkItemRelationType Represents the work item type relatiion type.
// swagger:model WorkItemRelationType
type WorkItemRelationType struct {
	WorkItemTrackingReference

	// The collection of relation type attributes.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkItemRelationType) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkItemTrackingReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkItemTrackingReference = aO0

	// now for regular properties
	var propsWorkItemRelationType struct {
		Attributes map[string]interface{} `json:"attributes,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsWorkItemRelationType); err != nil {
		return err
	}
	m.Attributes = propsWorkItemRelationType.Attributes

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkItemRelationType) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WorkItemTrackingReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsWorkItemRelationType struct {
		Attributes map[string]interface{} `json:"attributes,omitempty"`
	}
	propsWorkItemRelationType.Attributes = m.Attributes

	jsonDataPropsWorkItemRelationType, errWorkItemRelationType := swag.WriteJSON(propsWorkItemRelationType)
	if errWorkItemRelationType != nil {
		return nil, errWorkItemRelationType
	}
	_parts = append(_parts, jsonDataPropsWorkItemRelationType)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this work item relation type
func (m *WorkItemRelationType) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkItemTrackingReference
	if err := m.WorkItemTrackingReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemRelationType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemRelationType) UnmarshalBinary(b []byte) error {
	var res WorkItemRelationType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}