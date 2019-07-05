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

// UsersSummary Summary of licenses and extensions assigned to users in the organization
// swagger:model UsersSummary
type UsersSummary struct {

	// Available Access Levels
	AvailableAccessLevels []*AccessLevel `json:"availableAccessLevels"`

	// Summary of Extensions in the organization
	Extensions []*ExtensionSummaryData `json:"extensions"`

	// Group Options
	GroupOptions []*GroupOption `json:"groupOptions"`

	// Summary of Licenses in the organization
	Licenses []*LicenseSummaryData `json:"licenses"`

	// Summary of Projects in the organization
	ProjectRefs []*ProjectRef `json:"projectRefs"`
}

// Validate validates this users summary
func (m *UsersSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableAccessLevels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectRefs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsersSummary) validateAvailableAccessLevels(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailableAccessLevels) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableAccessLevels); i++ {
		if swag.IsZero(m.AvailableAccessLevels[i]) { // not required
			continue
		}

		if m.AvailableAccessLevels[i] != nil {
			if err := m.AvailableAccessLevels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availableAccessLevels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UsersSummary) validateExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Extensions); i++ {
		if swag.IsZero(m.Extensions[i]) { // not required
			continue
		}

		if m.Extensions[i] != nil {
			if err := m.Extensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UsersSummary) validateGroupOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupOptions); i++ {
		if swag.IsZero(m.GroupOptions[i]) { // not required
			continue
		}

		if m.GroupOptions[i] != nil {
			if err := m.GroupOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groupOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UsersSummary) validateLicenses(formats strfmt.Registry) error {

	if swag.IsZero(m.Licenses) { // not required
		return nil
	}

	for i := 0; i < len(m.Licenses); i++ {
		if swag.IsZero(m.Licenses[i]) { // not required
			continue
		}

		if m.Licenses[i] != nil {
			if err := m.Licenses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("licenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UsersSummary) validateProjectRefs(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectRefs) { // not required
		return nil
	}

	for i := 0; i < len(m.ProjectRefs); i++ {
		if swag.IsZero(m.ProjectRefs[i]) { // not required
			continue
		}

		if m.ProjectRefs[i] != nil {
			if err := m.ProjectRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsersSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsersSummary) UnmarshalBinary(b []byte) error {
	var res UsersSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}