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

// PointLastResult point last result
// swagger:model PointLastResult
type PointLastResult struct {

	// last updated date
	// Format: date-time
	LastUpdatedDate strfmt.DateTime `json:"lastUpdatedDate,omitempty"`

	// point Id
	PointID int32 `json:"pointId,omitempty"`
}

// Validate validates this point last result
func (m *PointLastResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PointLastResult) validateLastUpdatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedDate", "body", "date-time", m.LastUpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PointLastResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PointLastResult) UnmarshalBinary(b []byte) error {
	var res PointLastResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
