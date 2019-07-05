// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PointsReference2 points reference2
// swagger:model PointsReference2
type PointsReference2 struct {

	// plan Id
	PlanID int32 `json:"planId,omitempty"`

	// point Id
	PointID int32 `json:"pointId,omitempty"`
}

// Validate validates this points reference2
func (m *PointsReference2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PointsReference2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PointsReference2) UnmarshalBinary(b []byte) error {
	var res PointsReference2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}