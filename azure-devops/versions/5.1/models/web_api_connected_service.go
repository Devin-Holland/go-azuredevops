// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WebAPIConnectedService web Api connected service
// swagger:model WebApiConnectedService
type WebAPIConnectedService struct {
	WebAPIConnectedServiceRef

	// The user who did the OAuth authentication to created this service
	AuthenticatedBy *IdentityRef `json:"authenticatedBy,omitempty"`

	// Extra description on the service.
	Description string `json:"description,omitempty"`

	// Friendly Name of service connection
	FriendlyName string `json:"friendlyName,omitempty"`

	// Id/Name of the connection service. For Ex: Subscription Id for Azure Connection
	ID string `json:"id,omitempty"`

	// The kind of service.
	Kind string `json:"kind,omitempty"`

	// The project associated with this service
	Project *TeamProjectReference `json:"project,omitempty"`

	// Optional uri to connect directly to the service such as https://windows.azure.com
	ServiceURI string `json:"serviceUri,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WebAPIConnectedService) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WebAPIConnectedServiceRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WebAPIConnectedServiceRef = aO0

	// now for regular properties
	var propsWebAPIConnectedService struct {
		AuthenticatedBy *IdentityRef `json:"authenticatedBy,omitempty"`

		Description string `json:"description,omitempty"`

		FriendlyName string `json:"friendlyName,omitempty"`

		ID string `json:"id,omitempty"`

		Kind string `json:"kind,omitempty"`

		Project *TeamProjectReference `json:"project,omitempty"`

		ServiceURI string `json:"serviceUri,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsWebAPIConnectedService); err != nil {
		return err
	}
	m.AuthenticatedBy = propsWebAPIConnectedService.AuthenticatedBy

	m.Description = propsWebAPIConnectedService.Description

	m.FriendlyName = propsWebAPIConnectedService.FriendlyName

	m.ID = propsWebAPIConnectedService.ID

	m.Kind = propsWebAPIConnectedService.Kind

	m.Project = propsWebAPIConnectedService.Project

	m.ServiceURI = propsWebAPIConnectedService.ServiceURI

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WebAPIConnectedService) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WebAPIConnectedServiceRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsWebAPIConnectedService struct {
		AuthenticatedBy *IdentityRef `json:"authenticatedBy,omitempty"`

		Description string `json:"description,omitempty"`

		FriendlyName string `json:"friendlyName,omitempty"`

		ID string `json:"id,omitempty"`

		Kind string `json:"kind,omitempty"`

		Project *TeamProjectReference `json:"project,omitempty"`

		ServiceURI string `json:"serviceUri,omitempty"`
	}
	propsWebAPIConnectedService.AuthenticatedBy = m.AuthenticatedBy

	propsWebAPIConnectedService.Description = m.Description

	propsWebAPIConnectedService.FriendlyName = m.FriendlyName

	propsWebAPIConnectedService.ID = m.ID

	propsWebAPIConnectedService.Kind = m.Kind

	propsWebAPIConnectedService.Project = m.Project

	propsWebAPIConnectedService.ServiceURI = m.ServiceURI

	jsonDataPropsWebAPIConnectedService, errWebAPIConnectedService := swag.WriteJSON(propsWebAPIConnectedService)
	if errWebAPIConnectedService != nil {
		return nil, errWebAPIConnectedService
	}
	_parts = append(_parts, jsonDataPropsWebAPIConnectedService)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this web Api connected service
func (m *WebAPIConnectedService) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WebAPIConnectedServiceRef
	if err := m.WebAPIConnectedServiceRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebAPIConnectedService) validateAuthenticatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticatedBy) { // not required
		return nil
	}

	if m.AuthenticatedBy != nil {
		if err := m.AuthenticatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *WebAPIConnectedService) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebAPIConnectedService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebAPIConnectedService) UnmarshalBinary(b []byte) error {
	var res WebAPIConnectedService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}