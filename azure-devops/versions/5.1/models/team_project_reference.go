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

// TeamProjectReference Represents a shallow reference to a TeamProject.
// swagger:model TeamProjectReference
type TeamProjectReference struct {

	// Project abbreviation.
	Abbreviation string `json:"abbreviation,omitempty"`

	// Url to default team identity image.
	DefaultTeamImageURL string `json:"defaultTeamImageUrl,omitempty"`

	// The project's description (if any).
	Description string `json:"description,omitempty"`

	// Project identifier.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Project last update time.
	// Format: date-time
	LastUpdateTime strfmt.DateTime `json:"lastUpdateTime,omitempty"`

	// Project name.
	Name string `json:"name,omitempty"`

	// Project revision.
	Revision int64 `json:"revision,omitempty"`

	// Project state.
	// Enum: [deleting new wellFormed createPending all unchanged deleted]
	State interface{} `json:"state,omitempty"`

	// Url to the full version of the object.
	URL string `json:"url,omitempty"`

	// Project visibility.
	// Enum: [private public]
	Visibility interface{} `json:"visibility,omitempty"`
}

// Validate validates this team project reference
func (m *TeamProjectReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamProjectReference) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TeamProjectReference) validateLastUpdateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdateTime", "body", "date-time", m.LastUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamProjectReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamProjectReference) UnmarshalBinary(b []byte) error {
	var res TeamProjectReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}