// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RunFilter This class is used to provide the filters used for discovery
// swagger:model RunFilter
type RunFilter struct {

	// filter for the test case sources (test containers)
	SourceFilter string `json:"sourceFilter,omitempty"`

	// filter for the test cases
	TestCaseFilter string `json:"testCaseFilter,omitempty"`
}

// Validate validates this run filter
func (m *RunFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunFilter) UnmarshalBinary(b []byte) error {
	var res RunFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
