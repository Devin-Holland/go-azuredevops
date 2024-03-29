// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestCaseMetadata2 test case metadata2
// swagger:model TestCaseMetadata2
type TestCaseMetadata2 struct {

	// container
	Container string `json:"container,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project Id
	// Format: uuid
	ProjectID strfmt.UUID `json:"projectId,omitempty"`

	// test metadata Id
	TestMetadataID int32 `json:"testMetadataId,omitempty"`
}

// Validate validates this test case metadata2
func (m *TestCaseMetadata2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCaseMetadata2) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectId", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCaseMetadata2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCaseMetadata2) UnmarshalBinary(b []byte) error {
	var res TestCaseMetadata2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
