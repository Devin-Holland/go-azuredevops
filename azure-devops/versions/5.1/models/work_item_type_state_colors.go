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

// WorkItemTypeStateColors State colors for a work item type
// swagger:model WorkItemTypeStateColors
type WorkItemTypeStateColors struct {

	// Work item type state colors
	StateColors []*WorkItemStateColor `json:"stateColors"`

	// Work item type name
	WorkItemTypeName string `json:"workItemTypeName,omitempty"`
}

// Validate validates this work item type state colors
func (m *WorkItemTypeStateColors) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStateColors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkItemTypeStateColors) validateStateColors(formats strfmt.Registry) error {

	if swag.IsZero(m.StateColors) { // not required
		return nil
	}

	for i := 0; i < len(m.StateColors); i++ {
		if swag.IsZero(m.StateColors[i]) { // not required
			continue
		}

		if m.StateColors[i] != nil {
			if err := m.StateColors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stateColors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemTypeStateColors) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemTypeStateColors) UnmarshalBinary(b []byte) error {
	var res WorkItemTypeStateColors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}