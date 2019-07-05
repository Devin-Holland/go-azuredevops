// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Filter Describes a filter bucket item representing the total matches of search result, name and id.
// swagger:model Filter
type Filter struct {

	// Id of the filter bucket.
	ID string `json:"id,omitempty"`

	// Name of the filter bucket.
	Name string `json:"name,omitempty"`

	// Count of matches in the filter bucket.
	ResultCount int32 `json:"resultCount,omitempty"`
}

// Validate validates this filter
func (m *Filter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Filter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Filter) UnmarshalBinary(b []byte) error {
	var res Filter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}