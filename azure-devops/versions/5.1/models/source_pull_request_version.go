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

// SourcePullRequestVersion source pull request version
// swagger:model SourcePullRequestVersion
type SourcePullRequestVersion struct {

	// Pull Request Iteration Id for which the release will publish status.
	IterationID string `json:"iterationId,omitempty"`

	// Pull Request Id for which the release will publish status.
	PullRequestID string `json:"pullRequestId,omitempty"`

	// Date and time of the pull request merge creation. It is required to keep timeline record of Releases created by pull request.
	// Format: date-time
	PullRequestMergedAt strfmt.DateTime `json:"pullRequestMergedAt,omitempty"`

	// Source branch of the Pull Request.
	SourceBranch string `json:"sourceBranch,omitempty"`

	// Source branch commit Id of the Pull Request for which the release will publish status.
	SourceBranchCommitID string `json:"sourceBranchCommitId,omitempty"`

	// Target branch of the Pull Request.
	TargetBranch string `json:"targetBranch,omitempty"`
}

// Validate validates this source pull request version
func (m *SourcePullRequestVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePullRequestMergedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourcePullRequestVersion) validatePullRequestMergedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.PullRequestMergedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("pullRequestMergedAt", "body", "date-time", m.PullRequestMergedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourcePullRequestVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourcePullRequestVersion) UnmarshalBinary(b []byte) error {
	var res SourcePullRequestVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}