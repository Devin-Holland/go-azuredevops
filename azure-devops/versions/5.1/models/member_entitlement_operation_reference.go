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

// MemberEntitlementOperationReference member entitlement operation reference
// swagger:model MemberEntitlementOperationReference
type MemberEntitlementOperationReference struct {
	OperationReference

	// Operation completed with success or failure
	Completed bool `json:"completed,omitempty"`

	// True if all operations were successful
	HaveResultsSucceeded bool `json:"haveResultsSucceeded,omitempty"`

	// List of results for each operation
	Results []*OperationResult `json:"results"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MemberEntitlementOperationReference) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 OperationReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.OperationReference = aO0

	// now for regular properties
	var propsMemberEntitlementOperationReference struct {
		Completed bool `json:"completed,omitempty"`

		HaveResultsSucceeded bool `json:"haveResultsSucceeded,omitempty"`

		Results []*OperationResult `json:"results"`
	}
	if err := swag.ReadJSON(raw, &propsMemberEntitlementOperationReference); err != nil {
		return err
	}
	m.Completed = propsMemberEntitlementOperationReference.Completed

	m.HaveResultsSucceeded = propsMemberEntitlementOperationReference.HaveResultsSucceeded

	m.Results = propsMemberEntitlementOperationReference.Results

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MemberEntitlementOperationReference) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.OperationReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsMemberEntitlementOperationReference struct {
		Completed bool `json:"completed,omitempty"`

		HaveResultsSucceeded bool `json:"haveResultsSucceeded,omitempty"`

		Results []*OperationResult `json:"results"`
	}
	propsMemberEntitlementOperationReference.Completed = m.Completed

	propsMemberEntitlementOperationReference.HaveResultsSucceeded = m.HaveResultsSucceeded

	propsMemberEntitlementOperationReference.Results = m.Results

	jsonDataPropsMemberEntitlementOperationReference, errMemberEntitlementOperationReference := swag.WriteJSON(propsMemberEntitlementOperationReference)
	if errMemberEntitlementOperationReference != nil {
		return nil, errMemberEntitlementOperationReference
	}
	_parts = append(_parts, jsonDataPropsMemberEntitlementOperationReference)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this member entitlement operation reference
func (m *MemberEntitlementOperationReference) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with OperationReference
	if err := m.OperationReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemberEntitlementOperationReference) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemberEntitlementOperationReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberEntitlementOperationReference) UnmarshalBinary(b []byte) error {
	var res MemberEntitlementOperationReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
