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

// DeploymentAttempt deployment attempt
// swagger:model DeploymentAttempt
type DeploymentAttempt struct {

	// Deployment attempt.
	Attempt int32 `json:"attempt,omitempty"`

	// ID of the deployment.
	DeploymentID int32 `json:"deploymentId,omitempty"`

	// Specifies whether deployment has started or not.
	HasStarted bool `json:"hasStarted,omitempty"`

	// ID of deployment.
	ID int32 `json:"id,omitempty"`

	// All the issues related to the deployment.
	Issues []*Issue `json:"issues"`

	// Identity who last modified this deployment.
	LastModifiedBy *IdentityRef `json:"lastModifiedBy,omitempty"`

	// Time when this deployment last modified.
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"lastModifiedOn,omitempty"`

	// Deployment opeartion status.
	// Enum: [undefined queued scheduled pending approved rejected deferred queuedForAgent phaseInProgress phaseSucceeded phasePartiallySucceeded phaseFailed canceled phaseCanceled manualInterventionPending queuedForPipeline cancelling evaluatingGates gateFailed all]
	OperationStatus interface{} `json:"operationStatus,omitempty"`

	// Post deployment gates that executed in this deployment.
	PostDeploymentGates *ReleaseGates `json:"postDeploymentGates,omitempty"`

	// Pre deployment gates that executed in this deployment.
	PreDeploymentGates *ReleaseGates `json:"preDeploymentGates,omitempty"`

	// When this deployment queued on.
	// Format: date-time
	QueuedOn strfmt.DateTime `json:"queuedOn,omitempty"`

	// Reason for the deployment.
	// Enum: [none manual automated scheduled redeployTrigger]
	Reason interface{} `json:"reason,omitempty"`

	// List of release deployphases executed in this deployment.
	ReleaseDeployPhases []*ReleaseDeployPhase `json:"releaseDeployPhases"`

	// Identity who requested this deployment.
	RequestedBy *IdentityRef `json:"requestedBy,omitempty"`

	// Identity for this deployment requested.
	RequestedFor *IdentityRef `json:"requestedFor,omitempty"`

	// status of the deployment.
	// Enum: [undefined notDeployed inProgress succeeded partiallySucceeded failed all]
	Status interface{} `json:"status,omitempty"`
}

// Validate validates this deployment attempt
func (m *DeploymentAttempt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostDeploymentGates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreDeploymentGates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueuedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDeployPhases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedFor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentAttempt) validateIssues(formats strfmt.Registry) error {

	if swag.IsZero(m.Issues) { // not required
		return nil
	}

	for i := 0; i < len(m.Issues); i++ {
		if swag.IsZero(m.Issues[i]) { // not required
			continue
		}

		if m.Issues[i] != nil {
			if err := m.Issues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("issues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentAttempt) validateLastModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedBy) { // not required
		return nil
	}

	if m.LastModifiedBy != nil {
		if err := m.LastModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastModifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentAttempt) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentAttempt) validatePostDeploymentGates(formats strfmt.Registry) error {

	if swag.IsZero(m.PostDeploymentGates) { // not required
		return nil
	}

	if m.PostDeploymentGates != nil {
		if err := m.PostDeploymentGates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDeploymentGates")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentAttempt) validatePreDeploymentGates(formats strfmt.Registry) error {

	if swag.IsZero(m.PreDeploymentGates) { // not required
		return nil
	}

	if m.PreDeploymentGates != nil {
		if err := m.PreDeploymentGates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preDeploymentGates")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentAttempt) validateQueuedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.QueuedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("queuedOn", "body", "date-time", m.QueuedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentAttempt) validateReleaseDeployPhases(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseDeployPhases) { // not required
		return nil
	}

	for i := 0; i < len(m.ReleaseDeployPhases); i++ {
		if swag.IsZero(m.ReleaseDeployPhases[i]) { // not required
			continue
		}

		if m.ReleaseDeployPhases[i] != nil {
			if err := m.ReleaseDeployPhases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("releaseDeployPhases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentAttempt) validateRequestedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedBy) { // not required
		return nil
	}

	if m.RequestedBy != nil {
		if err := m.RequestedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedBy")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentAttempt) validateRequestedFor(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedFor) { // not required
		return nil
	}

	if m.RequestedFor != nil {
		if err := m.RequestedFor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedFor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentAttempt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentAttempt) UnmarshalBinary(b []byte) error {
	var res DeploymentAttempt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
