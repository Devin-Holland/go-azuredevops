// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ReportingWorkItemRevisionsFilter The class reprensents the reporting work item revision filer.
// swagger:model ReportingWorkItemRevisionsFilter
type ReportingWorkItemRevisionsFilter struct {

	// A list of fields to return in work item revisions. Omit this parameter to get all reportable fields.
	Fields []string `json:"fields"`

	// Include deleted work item in the result.
	IncludeDeleted bool `json:"includeDeleted,omitempty"`

	// Return an identity reference instead of a string value for identity fields.
	IncludeIdentityRef bool `json:"includeIdentityRef,omitempty"`

	// Include only the latest version of a work item, skipping over all previous revisions of the work item.
	IncludeLatestOnly bool `json:"includeLatestOnly,omitempty"`

	// Include tag reference instead of string value for System.Tags field
	IncludeTagRef bool `json:"includeTagRef,omitempty"`

	// A list of types to filter the results to specific work item types. Omit this parameter to get work item revisions of all work item types.
	Types []string `json:"types"`
}

// Validate validates this reporting work item revisions filter
func (m *ReportingWorkItemRevisionsFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportingWorkItemRevisionsFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingWorkItemRevisionsFilter) UnmarshalBinary(b []byte) error {
	var res ReportingWorkItemRevisionsFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}