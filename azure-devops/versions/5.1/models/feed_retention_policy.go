// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FeedRetentionPolicy Retention policy settings.
// swagger:model FeedRetentionPolicy
type FeedRetentionPolicy struct {

	// This attribute is deprecated and is not honoured by retention
	AgeLimitInDays int32 `json:"ageLimitInDays,omitempty"`

	// Maximum versions to preserve per package and package type.
	CountLimit int32 `json:"countLimit,omitempty"`

	// Number of days to preserve a package version after its latest download.
	DaysToKeepRecentlyDownloadedPackages int32 `json:"daysToKeepRecentlyDownloadedPackages,omitempty"`
}

// Validate validates this feed retention policy
func (m *FeedRetentionPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeedRetentionPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedRetentionPolicy) UnmarshalBinary(b []byte) error {
	var res FeedRetentionPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
