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

// TeamBacklogMapping Mapping of teams to the corresponding work item category
// swagger:model TeamBacklogMapping
type TeamBacklogMapping struct {

	// category reference name
	CategoryReferenceName string `json:"categoryReferenceName,omitempty"`

	// team Id
	// Format: uuid
	TeamID strfmt.UUID `json:"teamId,omitempty"`
}

// Validate validates this team backlog mapping
func (m *TeamBacklogMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeamID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamBacklogMapping) validateTeamID(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamID) { // not required
		return nil
	}

	if err := validate.FormatOf("teamId", "body", "uuid", m.TeamID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamBacklogMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamBacklogMapping) UnmarshalBinary(b []byte) error {
	var res TeamBacklogMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}