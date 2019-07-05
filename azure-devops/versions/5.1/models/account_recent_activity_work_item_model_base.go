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

// AccountRecentActivityWorkItemModelBase Represents Work Item Recent Activity
// swagger:model AccountRecentActivityWorkItemModelBase
type AccountRecentActivityWorkItemModelBase struct {

	// Date of the last Activity by the user
	// Format: date-time
	ActivityDate strfmt.DateTime `json:"activityDate,omitempty"`

	// Type of the activity
	// Enum: [visited edited deleted restored]
	ActivityType interface{} `json:"activityType,omitempty"`

	// Last changed date of the work item
	// Format: date-time
	ChangedDate strfmt.DateTime `json:"changedDate,omitempty"`

	// Work Item Id
	ID int32 `json:"id,omitempty"`

	// TeamFoundationId of the user this activity belongs to
	// Format: uuid
	IdentityID strfmt.UUID `json:"identityId,omitempty"`

	// State of the work item
	State string `json:"state,omitempty"`

	// Team project the work item belongs to
	TeamProject string `json:"teamProject,omitempty"`

	// Title of the work item
	Title string `json:"title,omitempty"`

	// Type of Work Item
	WorkItemType string `json:"workItemType,omitempty"`
}

// Validate validates this account recent activity work item model base
func (m *AccountRecentActivityWorkItemModelBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRecentActivityWorkItemModelBase) validateActivityDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityDate) { // not required
		return nil
	}

	if err := validate.FormatOf("activityDate", "body", "date-time", m.ActivityDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountRecentActivityWorkItemModelBase) validateChangedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("changedDate", "body", "date-time", m.ChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountRecentActivityWorkItemModelBase) validateIdentityID(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentityID) { // not required
		return nil
	}

	if err := validate.FormatOf("identityId", "body", "uuid", m.IdentityID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRecentActivityWorkItemModelBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRecentActivityWorkItemModelBase) UnmarshalBinary(b []byte) error {
	var res AccountRecentActivityWorkItemModelBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}