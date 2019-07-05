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

// Profile A user profile.
// swagger:model Profile
type Profile struct {

	// The attributes of this profile.
	ApplicationContainer *AttributesContainer `json:"applicationContainer,omitempty"`

	// The core attributes of this profile.
	CoreAttributes map[string]CoreProfileAttribute `json:"coreAttributes,omitempty"`

	// The maximum revision number of any attribute.
	CoreRevision int32 `json:"coreRevision,omitempty"`

	// The unique identifier of the profile.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The current state of the profile.
	// Enum: [custom customReadOnly readOnly]
	ProfileState interface{} `json:"profileState,omitempty"`

	// The maximum revision number of any attribute.
	Revision int32 `json:"revision,omitempty"`

	// The time at which this profile was last changed.
	// Format: date-time
	TimeStamp strfmt.DateTime `json:"timeStamp,omitempty"`
}

// Validate validates this profile
func (m *Profile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoreAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeStamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Profile) validateApplicationContainer(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicationContainer) { // not required
		return nil
	}

	if m.ApplicationContainer != nil {
		if err := m.ApplicationContainer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationContainer")
			}
			return err
		}
	}

	return nil
}

func (m *Profile) validateCoreAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.CoreAttributes) { // not required
		return nil
	}

	for k := range m.CoreAttributes {

		if err := validate.Required("coreAttributes"+"."+k, "body", m.CoreAttributes[k]); err != nil {
			return err
		}
		if val, ok := m.CoreAttributes[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Profile) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Profile) validateTimeStamp(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeStamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timeStamp", "body", "date-time", m.TimeStamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Profile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Profile) UnmarshalBinary(b []byte) error {
	var res Profile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}