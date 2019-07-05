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

// Change Represents a change associated with a build.
// swagger:model Change
type Change struct {

	// The author of the change.
	Author *IdentityRef `json:"author,omitempty"`

	// The type of source. "TfsVersionControl", "TfsGit", etc.
	ChangeType string `json:"changeType,omitempty"`

	// The location of a user-friendly representation of the resource.
	DisplayURI string `json:"displayUri,omitempty"`

	// Something that identifies the change. For a commit, this would be the SHA1. For a TFVC changeset, this would be the changeset id.
	ID string `json:"id,omitempty"`

	// The location of the full representation of the resource.
	Location string `json:"location,omitempty"`

	// A description of the change. This might be a commit message or changeset description.
	Message string `json:"message,omitempty"`

	// The person or process that pushed the change.
	PushedBy *IdentityRef `json:"pushedBy,omitempty"`

	// A timestamp for the change.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this change
func (m *Change) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePushedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Change) validateAuthor(formats strfmt.Registry) error {

	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {
		if err := m.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("author")
			}
			return err
		}
	}

	return nil
}

func (m *Change) validatePushedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.PushedBy) { // not required
		return nil
	}

	if m.PushedBy != nil {
		if err := m.PushedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pushedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Change) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Change) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Change) UnmarshalBinary(b []byte) error {
	var res Change
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}