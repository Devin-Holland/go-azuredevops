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

// WorkItemComment Comment on Work Item
// swagger:model WorkItemComment
type WorkItemComment struct {
	WorkItemTrackingResource

	// Identity of user who added the comment.
	RevisedBy *IdentityReference `json:"revisedBy,omitempty"`

	// The date of comment.
	// Format: date-time
	RevisedDate strfmt.DateTime `json:"revisedDate,omitempty"`

	// The work item revision number.
	Revision int32 `json:"revision,omitempty"`

	// The text of the comment.
	Text string `json:"text,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkItemComment) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkItemTrackingResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkItemTrackingResource = aO0

	// now for regular properties
	var propsWorkItemComment struct {
		RevisedBy *IdentityReference `json:"revisedBy,omitempty"`

		RevisedDate strfmt.DateTime `json:"revisedDate,omitempty"`

		Revision int32 `json:"revision,omitempty"`

		Text string `json:"text,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsWorkItemComment); err != nil {
		return err
	}
	m.RevisedBy = propsWorkItemComment.RevisedBy

	m.RevisedDate = propsWorkItemComment.RevisedDate

	m.Revision = propsWorkItemComment.Revision

	m.Text = propsWorkItemComment.Text

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkItemComment) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WorkItemTrackingResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsWorkItemComment struct {
		RevisedBy *IdentityReference `json:"revisedBy,omitempty"`

		RevisedDate strfmt.DateTime `json:"revisedDate,omitempty"`

		Revision int32 `json:"revision,omitempty"`

		Text string `json:"text,omitempty"`
	}
	propsWorkItemComment.RevisedBy = m.RevisedBy

	propsWorkItemComment.RevisedDate = m.RevisedDate

	propsWorkItemComment.Revision = m.Revision

	propsWorkItemComment.Text = m.Text

	jsonDataPropsWorkItemComment, errWorkItemComment := swag.WriteJSON(propsWorkItemComment)
	if errWorkItemComment != nil {
		return nil, errWorkItemComment
	}
	_parts = append(_parts, jsonDataPropsWorkItemComment)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this work item comment
func (m *WorkItemComment) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkItemTrackingResource
	if err := m.WorkItemTrackingResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevisedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevisedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkItemComment) validateRevisedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.RevisedBy) { // not required
		return nil
	}

	if m.RevisedBy != nil {
		if err := m.RevisedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revisedBy")
			}
			return err
		}
	}

	return nil
}

func (m *WorkItemComment) validateRevisedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RevisedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("revisedDate", "body", "date-time", m.RevisedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemComment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemComment) UnmarshalBinary(b []byte) error {
	var res WorkItemComment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
