// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReleaseApproval release approval
// swagger:model ReleaseApproval
type ReleaseApproval struct {

	// Gets or sets the type of approval.
	// Enum: [undefined preDeploy postDeploy all]
	ApprovalType interface{} `json:"approvalType,omitempty"`

	// Gets the identity who approved.
	ApprovedBy *IdentityRef `json:"approvedBy,omitempty"`

	// Gets or sets the identity who should approve.
	Approver *IdentityRef `json:"approver,omitempty"`

	// Gets or sets attempt which specifies as which deployment attempt it belongs.
	Attempt int32 `json:"attempt,omitempty"`

	// Gets or sets comments for approval.
	Comments string `json:"comments,omitempty"`

	// Gets date on which it got created.
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

	// Gets history which specifies all approvals associated with this approval.
	History []*ReleaseApprovalHistory `json:"history"`

	// Gets the unique identifier of this field.
	ID int32 `json:"id,omitempty"`

	// Gets or sets as approval is automated or not.
	IsAutomated bool `json:"isAutomated,omitempty"`

	// Gets date on which it got modified.
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modifiedOn,omitempty"`

	// Gets or sets rank which specifies the order of the approval. e.g. Same rank denotes parallel approval.
	Rank int32 `json:"rank,omitempty"`

	// Gets releaseReference which specifies the reference of the release to which this approval is associated.
	Release *ReleaseShallowReference `json:"release,omitempty"`

	// Gets releaseDefinitionReference which specifies the reference of the release definition to which this approval is associated.
	ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`

	// Gets releaseEnvironmentReference which specifies the reference of the release environment to which this approval is associated.
	ReleaseEnvironment *ReleaseEnvironmentShallowReference `json:"releaseEnvironment,omitempty"`

	// Gets the revision number.
	Revision int32 `json:"revision,omitempty"`

	// Gets or sets the status of the approval.
	// Enum: [undefined pending approved rejected reassigned canceled skipped]
	Status interface{} `json:"status,omitempty"`

	// Gets url to access the approval.
	URL string `json:"url,omitempty"`
}

// Validate validates this release approval
func (m *ReleaseApproval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApprovedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApprover(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseApproval) validateApprovedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ApprovedBy) { // not required
		return nil
	}

	if m.ApprovedBy != nil {
		if err := m.ApprovedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("approvedBy")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseApproval) validateApprover(formats strfmt.Registry) error {

	if swag.IsZero(m.Approver) { // not required
		return nil
	}

	if m.Approver != nil {
		if err := m.Approver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("approver")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseApproval) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("createdOn", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReleaseApproval) validateHistory(formats strfmt.Registry) error {

	if swag.IsZero(m.History) { // not required
		return nil
	}

	for i := 0; i < len(m.History); i++ {
		if swag.IsZero(m.History[i]) { // not required
			continue
		}

		if m.History[i] != nil {
			if err := m.History[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("history" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseApproval) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedOn", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReleaseApproval) validateRelease(formats strfmt.Registry) error {

	if swag.IsZero(m.Release) { // not required
		return nil
	}

	if m.Release != nil {
		if err := m.Release.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("release")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseApproval) validateReleaseDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseDefinition) { // not required
		return nil
	}

	if m.ReleaseDefinition != nil {
		if err := m.ReleaseDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releaseDefinition")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseApproval) validateReleaseEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseEnvironment) { // not required
		return nil
	}

	if m.ReleaseEnvironment != nil {
		if err := m.ReleaseEnvironment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releaseEnvironment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseApproval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseApproval) UnmarshalBinary(b []byte) error {
	var res ReleaseApproval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
