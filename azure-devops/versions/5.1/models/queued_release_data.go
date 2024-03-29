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

// QueuedReleaseData queued release data
// swagger:model QueuedReleaseData
type QueuedReleaseData struct {

	// Project ID of the release.
	// Format: uuid
	ProjectID strfmt.UUID `json:"projectId,omitempty"`

	// Release queue position.
	QueuePosition int32 `json:"queuePosition,omitempty"`

	// Queued release ID.
	ReleaseID int32 `json:"releaseId,omitempty"`
}

// Validate validates this queued release data
func (m *QueuedReleaseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueuedReleaseData) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectId", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueuedReleaseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueuedReleaseData) UnmarshalBinary(b []byte) error {
	var res QueuedReleaseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
