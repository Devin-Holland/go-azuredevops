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

// UpstreamSource Upstream source definition, including its Identity, package type, and other associated information.
// swagger:model UpstreamSource
type UpstreamSource struct {

	// UTC date that this upstream was deleted.
	// Format: date-time
	DeletedDate strfmt.DateTime `json:"deletedDate,omitempty"`

	// Locator for connecting to the upstream source in a user friendly format, that may potentially change over time
	DisplayLocation string `json:"displayLocation,omitempty"`

	// Identity of the upstream source.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// For an internal upstream type, track the Azure DevOps organization that contains it.
	// Format: uuid
	InternalUpstreamCollectionID strfmt.UUID `json:"internalUpstreamCollectionId,omitempty"`

	// For an internal upstream type, track the feed id being referenced.
	// Format: uuid
	InternalUpstreamFeedID strfmt.UUID `json:"internalUpstreamFeedId,omitempty"`

	// For an internal upstream type, track the view of the feed being referenced.
	// Format: uuid
	InternalUpstreamViewID strfmt.UUID `json:"internalUpstreamViewId,omitempty"`

	// Consistent locator for connecting to the upstream source.
	Location string `json:"location,omitempty"`

	// Display name.
	Name string `json:"name,omitempty"`

	// Package type associated with the upstream source.
	Protocol string `json:"protocol,omitempty"`

	// Source type, such as Public or Internal.
	// Enum: [public internal]
	UpstreamSourceType interface{} `json:"upstreamSourceType,omitempty"`
}

// Validate validates this upstream source
func (m *UpstreamSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeletedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalUpstreamCollectionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalUpstreamFeedID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalUpstreamViewID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpstreamSource) validateDeletedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("deletedDate", "body", "date-time", m.DeletedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpstreamSource) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpstreamSource) validateInternalUpstreamCollectionID(formats strfmt.Registry) error {

	if swag.IsZero(m.InternalUpstreamCollectionID) { // not required
		return nil
	}

	if err := validate.FormatOf("internalUpstreamCollectionId", "body", "uuid", m.InternalUpstreamCollectionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpstreamSource) validateInternalUpstreamFeedID(formats strfmt.Registry) error {

	if swag.IsZero(m.InternalUpstreamFeedID) { // not required
		return nil
	}

	if err := validate.FormatOf("internalUpstreamFeedId", "body", "uuid", m.InternalUpstreamFeedID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpstreamSource) validateInternalUpstreamViewID(formats strfmt.Registry) error {

	if swag.IsZero(m.InternalUpstreamViewID) { // not required
		return nil
	}

	if err := validate.FormatOf("internalUpstreamViewId", "body", "uuid", m.InternalUpstreamViewID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpstreamSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpstreamSource) UnmarshalBinary(b []byte) error {
	var res UpstreamSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
