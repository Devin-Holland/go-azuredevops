// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TestResultAttachmentIdentity test result attachment identity
// swagger:model TestResultAttachmentIdentity
type TestResultAttachmentIdentity struct {

	// attachment Id
	AttachmentID int32 `json:"attachmentId,omitempty"`

	// session Id
	SessionID int32 `json:"sessionId,omitempty"`

	// test result Id
	TestResultID int32 `json:"testResultId,omitempty"`

	// test run Id
	TestRunID int32 `json:"testRunId,omitempty"`
}

// Validate validates this test result attachment identity
func (m *TestResultAttachmentIdentity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestResultAttachmentIdentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultAttachmentIdentity) UnmarshalBinary(b []byte) error {
	var res TestResultAttachmentIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
