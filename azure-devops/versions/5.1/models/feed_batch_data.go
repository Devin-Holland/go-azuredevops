// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FeedBatchData feed batch data
// swagger:model FeedBatchData
type FeedBatchData struct {

	// data
	Data FeedBatchOperationData `json:"data,omitempty"`

	// operation
	// Enum: [saveCachedPackages]
	Operation interface{} `json:"operation,omitempty"`
}

// Validate validates this feed batch data
func (m *FeedBatchData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeedBatchData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedBatchData) UnmarshalBinary(b []byte) error {
	var res FeedBatchData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}