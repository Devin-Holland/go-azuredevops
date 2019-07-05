// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ReorderOperation Represents a reorder request for one or more work items.
// swagger:model ReorderOperation
type ReorderOperation struct {

	// IDs of the work items to be reordered.  Must be valid WorkItem Ids.
	Ids []int32 `json:"ids"`

	// IterationPath for reorder operation. This is only used when we reorder from the Iteration Backlog
	IterationPath string `json:"iterationPath,omitempty"`

	// ID of the work item that should be after the reordered items. Can use 0 to specify the end of the list.
	NextID int32 `json:"nextId,omitempty"`

	// Parent ID for all of the work items involved in this operation. Can use 0 to indicate the items don't have a parent.
	ParentID int32 `json:"parentId,omitempty"`

	// ID of the work item that should be before the reordered items. Can use 0 to specify the beginning of the list.
	PreviousID int32 `json:"previousId,omitempty"`
}

// Validate validates this reorder operation
func (m *ReorderOperation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReorderOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReorderOperation) UnmarshalBinary(b []byte) error {
	var res ReorderOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
