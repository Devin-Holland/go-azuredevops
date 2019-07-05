// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClientContributionNode Representaion of a ContributionNode that can be used for serialized to clients.
// swagger:model ClientContributionNode
type ClientContributionNode struct {

	// List of ids for contributions which are children to the current contribution.
	Children []string `json:"children"`

	// Contribution associated with this node.
	Contribution *ClientContribution `json:"contribution,omitempty"`

	// List of ids for contributions which are parents to the current contribution.
	Parents []string `json:"parents"`
}

// Validate validates this client contribution node
func (m *ClientContributionNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContribution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientContributionNode) validateContribution(formats strfmt.Registry) error {

	if swag.IsZero(m.Contribution) { // not required
		return nil
	}

	if m.Contribution != nil {
		if err := m.Contribution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contribution")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientContributionNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientContributionNode) UnmarshalBinary(b []byte) error {
	var res ClientContributionNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}