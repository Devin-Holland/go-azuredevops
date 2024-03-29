// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EnvironmentOptions environment options
// swagger:model EnvironmentOptions
type EnvironmentOptions struct {

	// Gets and sets as the auto link workitems or not.
	AutoLinkWorkItems bool `json:"autoLinkWorkItems,omitempty"`

	// Gets and sets as the badge enabled or not.
	BadgeEnabled bool `json:"badgeEnabled,omitempty"`

	// Gets and sets as the publish deployment status or not.
	PublishDeploymentStatus bool `json:"publishDeploymentStatus,omitempty"`

	// Gets and sets as the.pull request deployment enabled or not.
	PullRequestDeploymentEnabled bool `json:"pullRequestDeploymentEnabled,omitempty"`
}

// Validate validates this environment options
func (m *EnvironmentOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentOptions) UnmarshalBinary(b []byte) error {
	var res EnvironmentOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
