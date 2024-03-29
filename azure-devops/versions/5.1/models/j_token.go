// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// JToken Represents an abstract JSON token.
// swagger:model JToken
type JToken struct {

	// Get the first child token of this token.
	First *JToken `json:"first,omitempty"`

	// Gets a value indicating whether this token has child tokens.
	HasValues bool `json:"hasValues,omitempty"`

	// item
	Item *JToken `json:"item,omitempty"`

	// Get the last child token of this token.
	Last *JToken `json:"last,omitempty"`

	// Gets the next sibling token of this node.
	Next *JToken `json:"next,omitempty"`

	// Gets or sets the parent.
	Parent string `json:"parent,omitempty"`

	// Gets the path of the JSON token.
	Path string `json:"path,omitempty"`

	// Gets the previous sibling token of this node.
	Previous *JToken `json:"previous,omitempty"`

	// Gets the root JToken of this JToken.
	Root *JToken `json:"root,omitempty"`

	// Gets the node type for this JToken.
	Type string `json:"type,omitempty"`
}

// Validate validates this j token
func (m *JToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JToken) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(m.First) { // not required
		return nil
	}

	if m.First != nil {
		if err := m.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first")
			}
			return err
		}
	}

	return nil
}

func (m *JToken) validateItem(formats strfmt.Registry) error {

	if swag.IsZero(m.Item) { // not required
		return nil
	}

	if m.Item != nil {
		if err := m.Item.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

func (m *JToken) validateLast(formats strfmt.Registry) error {

	if swag.IsZero(m.Last) { // not required
		return nil
	}

	if m.Last != nil {
		if err := m.Last.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last")
			}
			return err
		}
	}

	return nil
}

func (m *JToken) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *JToken) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(m.Previous) { // not required
		return nil
	}

	if m.Previous != nil {
		if err := m.Previous.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("previous")
			}
			return err
		}
	}

	return nil
}

func (m *JToken) validateRoot(formats strfmt.Registry) error {

	if swag.IsZero(m.Root) { // not required
		return nil
	}

	if m.Root != nil {
		if err := m.Root.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("root")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JToken) UnmarshalBinary(b []byte) error {
	var res JToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
