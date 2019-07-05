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

// CommentVersion Represents a specific version of a comment on a work item.
// swagger:model CommentVersion
type CommentVersion struct {
	WorkItemTrackingResource

	// IdentityRef of the creator of the comment.
	CreatedBy *IdentityRef `json:"createdBy,omitempty"`

	// The creation date of the comment.
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// Effective Date/time value for adding the comment. Can be optionally different from CreatedDate.
	// Format: date-time
	CreatedOnBehalfDate strfmt.DateTime `json:"createdOnBehalfDate,omitempty"`

	// Identity on whose behalf this comment has been added. Can be optionally different from CreatedBy.
	CreatedOnBehalfOf *IdentityRef `json:"createdOnBehalfOf,omitempty"`

	// The id assigned to the comment.
	ID int32 `json:"id,omitempty"`

	// Indicates if the comment has been deleted at this version.
	IsDeleted bool `json:"isDeleted,omitempty"`

	// IdentityRef of the user who modified the comment at this version.
	ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`

	// The modification date of the comment for this version.
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// The rendered content of the comment at this version.
	RenderedText string `json:"renderedText,omitempty"`

	// The text of the comment at this version.
	Text string `json:"text,omitempty"`

	// The version number.
	Version int32 `json:"version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CommentVersion) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkItemTrackingResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkItemTrackingResource = aO0

	// now for regular properties
	var propsCommentVersion struct {
		CreatedBy *IdentityRef `json:"createdBy,omitempty"`

		CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

		CreatedOnBehalfDate strfmt.DateTime `json:"createdOnBehalfDate,omitempty"`

		CreatedOnBehalfOf *IdentityRef `json:"createdOnBehalfOf,omitempty"`

		ID int32 `json:"id,omitempty"`

		IsDeleted bool `json:"isDeleted,omitempty"`

		ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`

		ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

		RenderedText string `json:"renderedText,omitempty"`

		Text string `json:"text,omitempty"`

		Version int32 `json:"version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsCommentVersion); err != nil {
		return err
	}
	m.CreatedBy = propsCommentVersion.CreatedBy

	m.CreatedDate = propsCommentVersion.CreatedDate

	m.CreatedOnBehalfDate = propsCommentVersion.CreatedOnBehalfDate

	m.CreatedOnBehalfOf = propsCommentVersion.CreatedOnBehalfOf

	m.ID = propsCommentVersion.ID

	m.IsDeleted = propsCommentVersion.IsDeleted

	m.ModifiedBy = propsCommentVersion.ModifiedBy

	m.ModifiedDate = propsCommentVersion.ModifiedDate

	m.RenderedText = propsCommentVersion.RenderedText

	m.Text = propsCommentVersion.Text

	m.Version = propsCommentVersion.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CommentVersion) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WorkItemTrackingResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsCommentVersion struct {
		CreatedBy *IdentityRef `json:"createdBy,omitempty"`

		CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

		CreatedOnBehalfDate strfmt.DateTime `json:"createdOnBehalfDate,omitempty"`

		CreatedOnBehalfOf *IdentityRef `json:"createdOnBehalfOf,omitempty"`

		ID int32 `json:"id,omitempty"`

		IsDeleted bool `json:"isDeleted,omitempty"`

		ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`

		ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

		RenderedText string `json:"renderedText,omitempty"`

		Text string `json:"text,omitempty"`

		Version int32 `json:"version,omitempty"`
	}
	propsCommentVersion.CreatedBy = m.CreatedBy

	propsCommentVersion.CreatedDate = m.CreatedDate

	propsCommentVersion.CreatedOnBehalfDate = m.CreatedOnBehalfDate

	propsCommentVersion.CreatedOnBehalfOf = m.CreatedOnBehalfOf

	propsCommentVersion.ID = m.ID

	propsCommentVersion.IsDeleted = m.IsDeleted

	propsCommentVersion.ModifiedBy = m.ModifiedBy

	propsCommentVersion.ModifiedDate = m.ModifiedDate

	propsCommentVersion.RenderedText = m.RenderedText

	propsCommentVersion.Text = m.Text

	propsCommentVersion.Version = m.Version

	jsonDataPropsCommentVersion, errCommentVersion := swag.WriteJSON(propsCommentVersion)
	if errCommentVersion != nil {
		return nil, errCommentVersion
	}
	_parts = append(_parts, jsonDataPropsCommentVersion)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this comment version
func (m *CommentVersion) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkItemTrackingResource
	if err := m.WorkItemTrackingResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOnBehalfDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOnBehalfOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommentVersion) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *CommentVersion) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CommentVersion) validateCreatedOnBehalfDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOnBehalfDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdOnBehalfDate", "body", "date-time", m.CreatedOnBehalfDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CommentVersion) validateCreatedOnBehalfOf(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOnBehalfOf) { // not required
		return nil
	}

	if m.CreatedOnBehalfOf != nil {
		if err := m.CreatedOnBehalfOf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdOnBehalfOf")
			}
			return err
		}
	}

	return nil
}

func (m *CommentVersion) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *CommentVersion) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommentVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentVersion) UnmarshalBinary(b []byte) error {
	var res CommentVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}