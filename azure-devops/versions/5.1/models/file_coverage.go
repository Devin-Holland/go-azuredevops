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

// FileCoverage file coverage
// swagger:model FileCoverage
type FileCoverage struct {

	// List of line blocks along with their coverage status
	LineBlocksCoverage []*LineBlockCoverage `json:"lineBlocksCoverage"`

	// File path for which coverage information is sought for
	Path string `json:"path,omitempty"`
}

// Validate validates this file coverage
func (m *FileCoverage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineBlocksCoverage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileCoverage) validateLineBlocksCoverage(formats strfmt.Registry) error {

	if swag.IsZero(m.LineBlocksCoverage) { // not required
		return nil
	}

	for i := 0; i < len(m.LineBlocksCoverage); i++ {
		if swag.IsZero(m.LineBlocksCoverage[i]) { // not required
			continue
		}

		if m.LineBlocksCoverage[i] != nil {
			if err := m.LineBlocksCoverage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lineBlocksCoverage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileCoverage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileCoverage) UnmarshalBinary(b []byte) error {
	var res FileCoverage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}