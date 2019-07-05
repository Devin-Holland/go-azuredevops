// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ResultUpdateResponseModel result update response model
// swagger:model ResultUpdateResponseModel
type ResultUpdateResponseModel struct {

	// revision
	Revision int32 `json:"revision,omitempty"`
}

// Validate validates this result update response model
func (m *ResultUpdateResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResultUpdateResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultUpdateResponseModel) UnmarshalBinary(b []byte) error {
	var res ResultUpdateResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
