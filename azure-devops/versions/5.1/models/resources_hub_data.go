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

// ResourcesHubData resources hub data
// swagger:model ResourcesHubData
type ResourcesHubData struct {

	// continuation token
	ContinuationToken string `json:"continuationToken,omitempty"`

	// resource filter options
	ResourceFilterOptions *ResourceFilterOptions `json:"resourceFilterOptions,omitempty"`

	// resource filters
	ResourceFilters *ResourceFilters `json:"resourceFilters,omitempty"`

	// resource items
	ResourceItems []*ResourceItem `json:"resourceItems"`
}

// Validate validates this resources hub data
func (m *ResourcesHubData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceFilterOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcesHubData) validateResourceFilterOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceFilterOptions) { // not required
		return nil
	}

	if m.ResourceFilterOptions != nil {
		if err := m.ResourceFilterOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceFilterOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ResourcesHubData) validateResourceFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceFilters) { // not required
		return nil
	}

	if m.ResourceFilters != nil {
		if err := m.ResourceFilters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceFilters")
			}
			return err
		}
	}

	return nil
}

func (m *ResourcesHubData) validateResourceItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceItems); i++ {
		if swag.IsZero(m.ResourceItems[i]) { // not required
			continue
		}

		if m.ResourceItems[i] != nil {
			if err := m.ResourceItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourcesHubData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcesHubData) UnmarshalBinary(b []byte) error {
	var res ResourcesHubData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
