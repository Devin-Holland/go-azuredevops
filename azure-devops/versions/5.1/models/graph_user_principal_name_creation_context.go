// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GraphUserPrincipalNameCreationContext Use this type to create a new user using the principal name as a reference to an existing user from an external AD or AAD backed provider. This is the subset of GraphUser fields required for creation of a GraphUser for the AD and AAD use case when looking up the user by its principal name in the backing provider.
// swagger:model GraphUserPrincipalNameCreationContext
type GraphUserPrincipalNameCreationContext struct {
	GraphUserCreationContext

	// This should be the principal name or upn of the user in the source AD or AAD provider. Example: jamal@contoso.com Team Services will communicate with the source provider to fill all other fields on creation.
	PrincipalName string `json:"principalName,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GraphUserPrincipalNameCreationContext) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GraphUserCreationContext
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GraphUserCreationContext = aO0

	// now for regular properties
	var propsGraphUserPrincipalNameCreationContext struct {
		PrincipalName string `json:"principalName,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsGraphUserPrincipalNameCreationContext); err != nil {
		return err
	}
	m.PrincipalName = propsGraphUserPrincipalNameCreationContext.PrincipalName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GraphUserPrincipalNameCreationContext) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.GraphUserCreationContext)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsGraphUserPrincipalNameCreationContext struct {
		PrincipalName string `json:"principalName,omitempty"`
	}
	propsGraphUserPrincipalNameCreationContext.PrincipalName = m.PrincipalName

	jsonDataPropsGraphUserPrincipalNameCreationContext, errGraphUserPrincipalNameCreationContext := swag.WriteJSON(propsGraphUserPrincipalNameCreationContext)
	if errGraphUserPrincipalNameCreationContext != nil {
		return nil, errGraphUserPrincipalNameCreationContext
	}
	_parts = append(_parts, jsonDataPropsGraphUserPrincipalNameCreationContext)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this graph user principal name creation context
func (m *GraphUserPrincipalNameCreationContext) Validate(formats strfmt.Registry) error {
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
func (m *GraphUserPrincipalNameCreationContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphUserPrincipalNameCreationContext) UnmarshalBinary(b []byte) error {
	var res GraphUserPrincipalNameCreationContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
