// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyConfiguration The full policy configuration with settings.
// swagger:model PolicyConfiguration
type PolicyConfiguration struct {
	VersionedPolicyConfigurationRef

	// The links to other objects related to this object.
	Links *ReferenceLinks `json:"_links,omitempty"`

	// A reference to the identity that created the policy.
	CreatedBy *IdentityRef `json:"createdBy,omitempty"`

	// The date and time when the policy was created.
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// Indicates whether the policy is blocking.
	IsBlocking bool `json:"isBlocking,omitempty"`

	// Indicates whether the policy has been (soft) deleted.
	IsDeleted bool `json:"isDeleted,omitempty"`

	// Indicates whether the policy is enabled.
	IsEnabled bool `json:"isEnabled,omitempty"`

	// The policy configuration settings.
	Settings *JObject `json:"settings,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PolicyConfiguration) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 VersionedPolicyConfigurationRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.VersionedPolicyConfigurationRef = aO0

	// now for regular properties
	var propsPolicyConfiguration struct {
		Links *ReferenceLinks `json:"_links,omitempty"`

		CreatedBy *IdentityRef `json:"createdBy,omitempty"`

		CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

		IsBlocking bool `json:"isBlocking,omitempty"`

		IsDeleted bool `json:"isDeleted,omitempty"`

		IsEnabled bool `json:"isEnabled,omitempty"`

		Settings *JObject `json:"settings,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsPolicyConfiguration); err != nil {
		return err
	}
	m.Links = propsPolicyConfiguration.Links

	m.CreatedBy = propsPolicyConfiguration.CreatedBy

	m.CreatedDate = propsPolicyConfiguration.CreatedDate

	m.IsBlocking = propsPolicyConfiguration.IsBlocking

	m.IsDeleted = propsPolicyConfiguration.IsDeleted

	m.IsEnabled = propsPolicyConfiguration.IsEnabled

	m.Settings = propsPolicyConfiguration.Settings

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PolicyConfiguration) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.VersionedPolicyConfigurationRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsPolicyConfiguration struct {
		Links *ReferenceLinks `json:"_links,omitempty"`

		CreatedBy *IdentityRef `json:"createdBy,omitempty"`

		CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

		IsBlocking bool `json:"isBlocking,omitempty"`

		IsDeleted bool `json:"isDeleted,omitempty"`

		IsEnabled bool `json:"isEnabled,omitempty"`

		Settings *JObject `json:"settings,omitempty"`
	}
	propsPolicyConfiguration.Links = m.Links

	propsPolicyConfiguration.CreatedBy = m.CreatedBy

	propsPolicyConfiguration.CreatedDate = m.CreatedDate

	propsPolicyConfiguration.IsBlocking = m.IsBlocking

	propsPolicyConfiguration.IsDeleted = m.IsDeleted

	propsPolicyConfiguration.IsEnabled = m.IsEnabled

	propsPolicyConfiguration.Settings = m.Settings

	jsonDataPropsPolicyConfiguration, errPolicyConfiguration := swag.WriteJSON(propsPolicyConfiguration)
	if errPolicyConfiguration != nil {
		return nil, errPolicyConfiguration
	}
	_parts = append(_parts, jsonDataPropsPolicyConfiguration)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this policy configuration
func (m *PolicyConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VersionedPolicyConfigurationRef
	if err := m.VersionedPolicyConfigurationRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyConfiguration) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PolicyConfiguration) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *PolicyConfiguration) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PolicyConfiguration) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyConfiguration) UnmarshalBinary(b []byte) error {
	var res PolicyConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
