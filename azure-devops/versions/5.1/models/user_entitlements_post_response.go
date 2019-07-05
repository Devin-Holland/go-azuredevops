// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserEntitlementsPostResponse user entitlements post response
// swagger:model UserEntitlementsPostResponse
type UserEntitlementsPostResponse struct {
	UserEntitlementsResponseBase

	// Operation result.
	OperationResult *UserEntitlementOperationResult `json:"operationResult,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UserEntitlementsPostResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 UserEntitlementsResponseBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.UserEntitlementsResponseBase = aO0

	// now for regular properties
	var propsUserEntitlementsPostResponse struct {
		OperationResult *UserEntitlementOperationResult `json:"operationResult,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsUserEntitlementsPostResponse); err != nil {
		return err
	}
	m.OperationResult = propsUserEntitlementsPostResponse.OperationResult

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UserEntitlementsPostResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.UserEntitlementsResponseBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsUserEntitlementsPostResponse struct {
		OperationResult *UserEntitlementOperationResult `json:"operationResult,omitempty"`
	}
	propsUserEntitlementsPostResponse.OperationResult = m.OperationResult

	jsonDataPropsUserEntitlementsPostResponse, errUserEntitlementsPostResponse := swag.WriteJSON(propsUserEntitlementsPostResponse)
	if errUserEntitlementsPostResponse != nil {
		return nil, errUserEntitlementsPostResponse
	}
	_parts = append(_parts, jsonDataPropsUserEntitlementsPostResponse)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this user entitlements post response
func (m *UserEntitlementsPostResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UserEntitlementsResponseBase
	if err := m.UserEntitlementsResponseBase.Validate(formats); err != nil {
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

func (m *UserEntitlementsPostResponse) validateOperationResult(formats strfmt.Registry) error {

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
func (m *UserEntitlementsPostResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserEntitlementsPostResponse) UnmarshalBinary(b []byte) error {
	var res UserEntitlementsPostResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
