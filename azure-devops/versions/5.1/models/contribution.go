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

// Contribution An individual contribution made by an extension
// swagger:model Contribution
type Contribution struct {
	ContributionBase

	// List of constraints (filters) that should be applied to the availability of this contribution
	Constraints []*ContributionConstraint `json:"constraints"`

	// Includes is a set of contributions that should have this contribution included in their targets list.
	Includes []string `json:"includes"`

	// Properties/attributes of this contribution
	Properties *JObject `json:"properties,omitempty"`

	// List of demanded claims in order for the user to see this contribution (like anonymous, public, member...).
	RestrictedTo []string `json:"restrictedTo"`

	// The ids of the contribution(s) that this contribution targets. (parent contributions)
	Targets []string `json:"targets"`

	// Id of the Contribution Type
	Type string `json:"type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Contribution) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ContributionBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ContributionBase = aO0

	// now for regular properties
	var propsContribution struct {
		Constraints []*ContributionConstraint `json:"constraints"`

		Includes []string `json:"includes"`

		Properties *JObject `json:"properties,omitempty"`

		RestrictedTo []string `json:"restrictedTo"`

		Targets []string `json:"targets"`

		Type string `json:"type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsContribution); err != nil {
		return err
	}
	m.Constraints = propsContribution.Constraints

	m.Includes = propsContribution.Includes

	m.Properties = propsContribution.Properties

	m.RestrictedTo = propsContribution.RestrictedTo

	m.Targets = propsContribution.Targets

	m.Type = propsContribution.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Contribution) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ContributionBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsContribution struct {
		Constraints []*ContributionConstraint `json:"constraints"`

		Includes []string `json:"includes"`

		Properties *JObject `json:"properties,omitempty"`

		RestrictedTo []string `json:"restrictedTo"`

		Targets []string `json:"targets"`

		Type string `json:"type,omitempty"`
	}
	propsContribution.Constraints = m.Constraints

	propsContribution.Includes = m.Includes

	propsContribution.Properties = m.Properties

	propsContribution.RestrictedTo = m.RestrictedTo

	propsContribution.Targets = m.Targets

	propsContribution.Type = m.Type

	jsonDataPropsContribution, errContribution := swag.WriteJSON(propsContribution)
	if errContribution != nil {
		return nil, errContribution
	}
	_parts = append(_parts, jsonDataPropsContribution)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this contribution
func (m *Contribution) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ContributionBase
	if err := m.ContributionBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Contribution) validateConstraints(formats strfmt.Registry) error {

	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	for i := 0; i < len(m.Constraints); i++ {
		if swag.IsZero(m.Constraints[i]) { // not required
			continue
		}

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Contribution) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Contribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Contribution) UnmarshalBinary(b []byte) error {
	var res Contribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
