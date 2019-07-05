// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CommentMention comment mention
// swagger:model CommentMention
type CommentMention struct {
	WorkItemTrackingResource

	// The artifact portion of the parsed text. (i.e. the work item's id)
	ArtifactID string `json:"artifactId,omitempty"`

	// The type the parser assigned to the mention. (i.e. person, work item, etc)
	ArtifactType string `json:"artifactType,omitempty"`

	// The comment id of the mention.
	CommentID int32 `json:"commentId,omitempty"`

	// The resolved target of the mention. An example of this could be a user's tfid
	TargetID string `json:"targetId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CommentMention) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkItemTrackingResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkItemTrackingResource = aO0

	// now for regular properties
	var propsCommentMention struct {
		ArtifactID string `json:"artifactId,omitempty"`

		ArtifactType string `json:"artifactType,omitempty"`

		CommentID int32 `json:"commentId,omitempty"`

		TargetID string `json:"targetId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsCommentMention); err != nil {
		return err
	}
	m.ArtifactID = propsCommentMention.ArtifactID

	m.ArtifactType = propsCommentMention.ArtifactType

	m.CommentID = propsCommentMention.CommentID

	m.TargetID = propsCommentMention.TargetID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CommentMention) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WorkItemTrackingResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsCommentMention struct {
		ArtifactID string `json:"artifactId,omitempty"`

		ArtifactType string `json:"artifactType,omitempty"`

		CommentID int32 `json:"commentId,omitempty"`

		TargetID string `json:"targetId,omitempty"`
	}
	propsCommentMention.ArtifactID = m.ArtifactID

	propsCommentMention.ArtifactType = m.ArtifactType

	propsCommentMention.CommentID = m.CommentID

	propsCommentMention.TargetID = m.TargetID

	jsonDataPropsCommentMention, errCommentMention := swag.WriteJSON(propsCommentMention)
	if errCommentMention != nil {
		return nil, errCommentMention
	}
	_parts = append(_parts, jsonDataPropsCommentMention)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this comment mention
func (m *CommentMention) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkItemTrackingResource
	if err := m.WorkItemTrackingResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CommentMention) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentMention) UnmarshalBinary(b []byte) error {
	var res CommentMention
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}