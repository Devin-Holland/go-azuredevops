// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GraphUserOriginIDCreationContext Use this type to create a new user using the OriginID as a reference to an existing user from an external AD or AAD backed provider. This is the subset of GraphUser fields required for creation of a GraphUser for the AD and AAD use case when looking up the user by its unique ID in the backing provider.
// swagger:model GraphUserOriginIdCreationContext
type GraphUserOriginIDCreationContext struct {
	GraphUserCreationContext

	// This should be the name of the origin provider. Example: github.com
	Origin string `json:"origin,omitempty"`

	// This should be the object id or sid of the user from the source AD or AAD provider. Example: d47d025a-ce2f-4a79-8618-e8862ade30dd Team Services will communicate with the source provider to fill all other fields on creation.
	OriginID string `json:"originId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GraphUserOriginIDCreationContext) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GraphUserCreationContext
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GraphUserCreationContext = aO0

	// now for regular properties
	var propsGraphUserOriginIDCreationContext struct {
		Origin string `json:"origin,omitempty"`

		OriginID string `json:"originId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsGraphUserOriginIDCreationContext); err != nil {
		return err
	}
	m.Origin = propsGraphUserOriginIDCreationContext.Origin

	m.OriginID = propsGraphUserOriginIDCreationContext.OriginID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GraphUserOriginIDCreationContext) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.GraphUserCreationContext)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsGraphUserOriginIDCreationContext struct {
		Origin string `json:"origin,omitempty"`

		OriginID string `json:"originId,omitempty"`
	}
	propsGraphUserOriginIDCreationContext.Origin = m.Origin

	propsGraphUserOriginIDCreationContext.OriginID = m.OriginID

	jsonDataPropsGraphUserOriginIDCreationContext, errGraphUserOriginIDCreationContext := swag.WriteJSON(propsGraphUserOriginIDCreationContext)
	if errGraphUserOriginIDCreationContext != nil {
		return nil, errGraphUserOriginIDCreationContext
	}
	_parts = append(_parts, jsonDataPropsGraphUserOriginIDCreationContext)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this graph user origin Id creation context
func (m *GraphUserOriginIDCreationContext) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GraphUserCreationContext
	if err := m.GraphUserCreationContext.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GraphUserOriginIDCreationContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphUserOriginIDCreationContext) UnmarshalBinary(b []byte) error {
	var res GraphUserOriginIDCreationContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}