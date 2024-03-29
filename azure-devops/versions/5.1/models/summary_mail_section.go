// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SummaryMailSection summary mail section
// swagger:model SummaryMailSection
type SummaryMailSection struct {

	// Html content of summary mail.
	HTMLContent string `json:"htmlContent,omitempty"`

	// Rank of the summary mail.
	Rank int32 `json:"rank,omitempty"`

	// Summary mail section type. MailSectionType has section types.
	// Enum: [details environments issues testResults workItems releaseInfo]
	SectionType interface{} `json:"sectionType,omitempty"`

	// Title of the summary mail.
	Title string `json:"title,omitempty"`
}

// Validate validates this summary mail section
func (m *SummaryMailSection) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SummaryMailSection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryMailSection) UnmarshalBinary(b []byte) error {
	var res SummaryMailSection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
