// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkItemRelationUpdates Descrives updates to a work item's relations.
// swagger:model WorkItemRelationUpdates
type WorkItemRelationUpdates struct {

	// List of newly added relations.
	Added []*WorkItemRelation `json:"added"`

	// List of removed relations.
	Removed []*WorkItemRelation `json:"removed"`

	// List of updated relations.
	Updated []*WorkItemRelation `json:"updated"`
}

// Validate validates this work item relation updates
func (m *WorkItemRelationUpdates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkItemRelationUpdates) validateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.Added) { // not required
		return nil
	}

	for i := 0; i < len(m.Added); i++ {
		if swag.IsZero(m.Added[i]) { // not required
			continue
		}

		if m.Added[i] != nil {
			if err := m.Added[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("added" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkItemRelationUpdates) validateRemoved(formats strfmt.Registry) error {

	if swag.IsZero(m.Removed) { // not required
		return nil
	}

	for i := 0; i < len(m.Removed); i++ {
		if swag.IsZero(m.Removed[i]) { // not required
			continue
		}

		if m.Removed[i] != nil {
			if err := m.Removed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("removed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkItemRelationUpdates) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	for i := 0; i < len(m.Updated); i++ {
		if swag.IsZero(m.Updated[i]) { // not required
			continue
		}

		if m.Updated[i] != nil {
			if err := m.Updated[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updated" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemRelationUpdates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemRelationUpdates) UnmarshalBinary(b []byte) error {
	var res WorkItemRelationUpdates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
