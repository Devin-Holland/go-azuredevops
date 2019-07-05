// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MemberEntitlementsPostResponse member entitlements post response
// swagger:model MemberEntitlementsPostResponse
type MemberEntitlementsPostResponse struct {
	MemberEntitlementsResponseBase

	// Operation result
	OperationResult *OperationResult `json:"operationResult,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MemberEntitlementsPostResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MemberEntitlementsResponseBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MemberEntitlementsResponseBase = aO0

	// now for regular properties
	var propsMemberEntitlementsPostResponse struct {
		OperationResult *OperationResult `json:"operationResult,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsMemberEntitlementsPostResponse); err != nil {
		return err
	}
	m.OperationResult = propsMemberEntitlementsPostResponse.OperationResult

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MemberEntitlementsPostResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.MemberEntitlementsResponseBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsMemberEntitlementsPostResponse struct {
		OperationResult *OperationResult `json:"operationResult,omitempty"`
	}
	propsMemberEntitlementsPostResponse.OperationResult = m.OperationResult

	jsonDataPropsMemberEntitlementsPostResponse, errMemberEntitlementsPostResponse := swag.WriteJSON(propsMemberEntitlementsPostResponse)
	if errMemberEntitlementsPostResponse != nil {
		return nil, errMemberEntitlementsPostResponse
	}
	_parts = append(_parts, jsonDataPropsMemberEntitlementsPostResponse)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this member entitlements post response
func (m *MemberEntitlementsPostResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MemberEntitlementsResponseBase
	if err := m.MemberEntitlementsResponseBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemberEntitlementsPostResponse) validateOperationResult(formats strfmt.Registry) error {

	if swag.IsZero(m.OperationResult) { // not required
		return nil
	}

	if m.OperationResult != nil {
		if err := m.OperationResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationResult")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemberEntitlementsPostResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberEntitlementsPostResponse) UnmarshalBinary(b []byte) error {
	var res MemberEntitlementsPostResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
