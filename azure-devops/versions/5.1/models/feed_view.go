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

// FeedView A view on top of a feed.
// swagger:model FeedView
type FeedView struct {

	// Related REST links.
	Links *ReferenceLinks `json:"_links,omitempty"`

	// Id of the view.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name of the view.
	Name string `json:"name,omitempty"`

	// Type of view.
	// Enum: [none release implicit]
	Type interface{} `json:"type,omitempty"`

	// Url of the view.
	URL string `json:"url,omitempty"`

	// Visibility status of the view.
	// Enum: [private collection organization]
	Visibility interface{} `json:"visibility,omitempty"`
}

// Validate validates this feed view
func (m *FeedView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeedView) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *FeedView) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeedView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedView) UnmarshalBinary(b []byte) error {
	var res FeedView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
