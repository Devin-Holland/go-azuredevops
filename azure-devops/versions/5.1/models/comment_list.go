// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CommentList Represents a list of work item comments.
// swagger:model CommentList
type CommentList struct {
	WorkItemTrackingResource

	// List of comments in the current batch.
	Comments []*Comment `json:"comments"`

	// A string token that can be used to retrieving next page of comments if available. Otherwise null.
	ContinuationToken string `json:"continuationToken,omitempty"`

	// The count of comments in the current batch.
	Count int32 `json:"count,omitempty"`

	// Uri to the next page of comments if it is available. Otherwise null.
	NextPage string `json:"nextPage,omitempty"`

	// Total count of comments on a work item.
	TotalCount int32 `json:"totalCount,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CommentList) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkItemTrackingResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkItemTrackingResource = aO0

	// now for regular properties
	var propsCommentList struct {
		Comments []*Comment `json:"comments"`

		ContinuationToken string `json:"continuationToken,omitempty"`

		Count int32 `json:"count,omitempty"`

		NextPage string `json:"nextPage,omitempty"`

		TotalCount int32 `json:"totalCount,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsCommentList); err != nil {
		return err
	}
	m.Comments = propsCommentList.Comments

	m.ContinuationToken = propsCommentList.ContinuationToken

	m.Count = propsCommentList.Count

	m.NextPage = propsCommentList.NextPage

	m.TotalCount = propsCommentList.TotalCount

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CommentList) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WorkItemTrackingResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsCommentList struct {
		Comments []*Comment `json:"comments"`

		ContinuationToken string `json:"continuationToken,omitempty"`

		Count int32 `json:"count,omitempty"`

		NextPage string `json:"nextPage,omitempty"`

		TotalCount int32 `json:"totalCount,omitempty"`
	}
	propsCommentList.Comments = m.Comments

	propsCommentList.ContinuationToken = m.ContinuationToken

	propsCommentList.Count = m.Count

	propsCommentList.NextPage = m.NextPage

	propsCommentList.TotalCount = m.TotalCount

	jsonDataPropsCommentList, errCommentList := swag.WriteJSON(propsCommentList)
	if errCommentList != nil {
		return nil, errCommentList
	}
	_parts = append(_parts, jsonDataPropsCommentList)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this comment list
func (m *CommentList) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkItemTrackingResource
	if err := m.WorkItemTrackingResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommentList) validateComments(formats strfmt.Registry) error {

	if swag.IsZero(m.Comments) { // not required
		return nil
	}

	for i := 0; i < len(m.Comments); i++ {
		if swag.IsZero(m.Comments[i]) { // not required
			continue
		}

		if m.Comments[i] != nil {
			if err := m.Comments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommentList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentList) UnmarshalBinary(b []byte) error {
	var res CommentList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
