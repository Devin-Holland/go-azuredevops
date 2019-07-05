// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProtocolMetadata Extended metadata for a specific package type.
// swagger:model ProtocolMetadata
type ProtocolMetadata struct {

	// Extended metadata for a specific package type, formatted to the associated schema version definition.
	Data interface{} `json:"data,omitempty"`

	// Schema version.
	SchemaVersion int32 `json:"schemaVersion,omitempty"`
}

// Validate validates this protocol metadata
func (m *ProtocolMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtocolMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtocolMetadata) UnmarshalBinary(b []byte) error {
	var res ProtocolMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}