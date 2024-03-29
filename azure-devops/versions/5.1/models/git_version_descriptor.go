// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GitVersionDescriptor git version descriptor
// swagger:model GitVersionDescriptor
type GitVersionDescriptor struct {

	// Version string identifier (name of tag/branch, SHA1 of commit)
	Version string `json:"version,omitempty"`

	// Version options - Specify additional modifiers to version (e.g Previous)
	// Enum: [none previousChange firstParent]
	VersionOptions interface{} `json:"versionOptions,omitempty"`

	// Version type (branch, tag, or commit). Determines how Id is interpreted
	// Enum: [branch tag commit]
	VersionType interface{} `json:"versionType,omitempty"`
}

// Validate validates this git version descriptor
func (m *GitVersionDescriptor) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitVersionDescriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitVersionDescriptor) UnmarshalBinary(b []byte) error {
	var res GitVersionDescriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
