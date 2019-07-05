// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ContributionPropertyDescription Description about a property of a contribution type
// swagger:model ContributionPropertyDescription
type ContributionPropertyDescription struct {

	// Description of the property
	Description string `json:"description,omitempty"`

	// Name of the property
	Name string `json:"name,omitempty"`

	// True if this property is required
	Required bool `json:"required,omitempty"`

	// The type of value used for this property
	// Enum: [unknown string uri guid boolean integer double dateTime dictionary array object]
	Type interface{} `json:"type,omitempty"`
}

// Validate validates this contribution property description
func (m *ContributionPropertyDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContributionPropertyDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContributionPropertyDescription) UnmarshalBinary(b []byte) error {
	var res ContributionPropertyDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
