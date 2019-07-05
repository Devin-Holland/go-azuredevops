// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DownloadAttachmentsRequest download attachments request
// swagger:model DownloadAttachmentsRequest
type DownloadAttachmentsRequest struct {

	// ids
	Ids []int32 `json:"ids"`

	// lengths
	Lengths []int64 `json:"lengths"`
}

// Validate validates this download attachments request
func (m *DownloadAttachmentsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DownloadAttachmentsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DownloadAttachmentsRequest) UnmarshalBinary(b []byte) error {
	var res DownloadAttachmentsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}