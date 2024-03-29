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

// AccessControlListsCollection A list of AccessControlList. An AccessControlList is meant to associate a set of AccessControlEntries with a security token and its inheritance settings.
// swagger:model AccessControlListsCollection
type AccessControlListsCollection struct {
	AccessControlListsCollectionAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AccessControlListsCollection) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AccessControlListsCollectionAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AccessControlListsCollectionAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AccessControlListsCollection) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.AccessControlListsCollectionAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this access control lists collection
func (m *AccessControlListsCollection) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AccessControlListsCollectionAllOf0
	if err := m.AccessControlListsCollectionAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AccessControlListsCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessControlListsCollection) UnmarshalBinary(b []byte) error {
	var res AccessControlListsCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccessControlListsCollectionAllOf0 access control lists collection all of0
// swagger:model AccessControlListsCollectionAllOf0
type AccessControlListsCollectionAllOf0 []*AccessControlList

// Validate validates this access control lists collection all of0
func (m AccessControlListsCollectionAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
