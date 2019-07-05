// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GraphCachePolicies graph cache policies
// swagger:model GraphCachePolicies
type GraphCachePolicies struct {

	// Size of the cache
	CacheSize int32 `json:"cacheSize,omitempty"`
}

// Validate validates this graph cache policies
func (m *GraphCachePolicies) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GraphCachePolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphCachePolicies) UnmarshalBinary(b []byte) error {
	var res GraphCachePolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
