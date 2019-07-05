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

// Deployment deployment
// swagger:model Deployment
type Deployment struct {

	// Gets attempt number.
	Attempt int32 `json:"attempt,omitempty"`

	// Gets the date on which deployment is complete.
	// Format: date-time
	CompletedOn strfmt.DateTime `json:"completedOn,omitempty"`

	// Gets the list of condition associated with deployment.
	Conditions []*Condition `json:"conditions"`

	// Gets release definition environment id.
	DefinitionEnvironmentID int32 `json:"definitionEnvironmentId,omitempty"`

	// Gets status of the deployment.
	// Enum: [undefined notDeployed inProgress succeeded partiallySucceeded failed all]
	DeploymentStatus interface{} `json:"deploymentStatus,omitempty"`

	// Gets the unique identifier for deployment.
	ID int32 `json:"id,omitempty"`

	// Gets the identity who last modified the deployment.
	LastModifiedBy *IdentityRef `json:"lastModifiedBy,omitempty"`

	// Gets the date on which deployment is last modified.
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"lastModifiedOn,omitempty"`

	// Gets operation status of deployment.
	// Enum: [undefined queued scheduled pending approved rejected deferred queuedForAgent phaseInProgress phaseSucceeded phasePartiallySucceeded phaseFailed canceled phaseCanceled manualInterventionPending queuedForPipeline cancelling evaluatingGates gateFailed all]
	OperationStatus interface{} `json:"operationStatus,omitempty"`

	// Gets list of PostDeployApprovals.
	PostDeployApprovals []*ReleaseApproval `json:"postDeployApprovals"`

	// Gets list of PreDeployApprovals.
	PreDeployApprovals []*ReleaseApproval `json:"preDeployApprovals"`

	// Gets or sets project reference.
	ProjectReference *ProjectReference `json:"projectReference,omitempty"`

	// Gets the date on which deployment is queued.
	// Format: date-time
	QueuedOn strfmt.DateTime `json:"queuedOn,omitempty"`

	// Gets reason of deployment.
	// Enum: [none manual automated scheduled redeployTrigger]
	Reason interface{} `json:"reason,omitempty"`

	// Gets the reference of release.
	Release *ReleaseReference `json:"release,omitempty"`

	// Gets releaseDefinitionReference which specifies the reference of the release definition to which the deployment is associated.
	ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`

	// Gets releaseEnvironmentReference which specifies the reference of the release environment to which the deployment is associated.
	ReleaseEnvironment *ReleaseEnvironmentShallowReference `json:"releaseEnvironment,omitempty"`

	// Gets the identity who requested.
	RequestedBy *IdentityRef `json:"requestedBy,omitempty"`

	// Gets the identity for whom deployment is requested.
	RequestedFor *IdentityRef `json:"requestedFor,omitempty"`

	// Gets the date on which deployment is scheduled.
	// Format: date-time
	ScheduledDeploymentTime strfmt.DateTime `json:"scheduledDeploymentTime,omitempty"`

	// Gets the date on which deployment is started.
	// Format: date-time
	StartedOn strfmt.DateTime `json:"startedOn,omitempty"`
}

// Validate validates this deployment
func (m *Deployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostDeployApprovals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreDeployApprovals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueuedOn(formats); err != nil {
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

	if err := m.validateRequestedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedFor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledDeploymentTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deployment) validateCompletedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("completedOn", "body", "date-time", m.CompletedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Deployment) validateLastModifiedBy(formats strfmt.Registry) error {

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

func (m *Deployment) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validatePostDeployApprovals(formats strfmt.Registry) error {

	if swag.IsZero(m.PostDeployApprovals) { // not required
		return nil
	}

	for i := 0; i < len(m.PostDeployApprovals); i++ {
		if swag.IsZero(m.PostDeployApprovals[i]) { // not required
			continue
		}

		if m.PostDeployApprovals[i] != nil {
			if err := m.PostDeployApprovals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postDeployApprovals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Deployment) validatePreDeployApprovals(formats strfmt.Registry) error {

	if swag.IsZero(m.PreDeployApprovals) { // not required
		return nil
	}

	for i := 0; i < len(m.PreDeployApprovals); i++ {
		if swag.IsZero(m.PreDeployApprovals[i]) { // not required
			continue
		}

		if m.PreDeployApprovals[i] != nil {
			if err := m.PreDeployApprovals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preDeployApprovals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Deployment) validateProjectReference(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectReference) { // not required
		return nil
	}

	if m.ProjectReference != nil {
		if err := m.ProjectReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projectReference")
			}
			return err
		}
	}

	return nil
}

func (m *Deployment) validateQueuedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.QueuedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("queuedOn", "body", "date-time", m.QueuedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validateRelease(formats strfmt.Registry) error {

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

func (m *Deployment) validateReleaseDefinition(formats strfmt.Registry) error {

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

func (m *Deployment) validateReleaseEnvironment(formats strfmt.Registry) error {

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

func (m *Deployment) validateRequestedBy(formats strfmt.Registry) error {

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

func (m *Deployment) validateRequestedFor(formats strfmt.Registry) error {

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

func (m *Deployment) validateScheduledDeploymentTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledDeploymentTime) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduledDeploymentTime", "body", "date-time", m.ScheduledDeploymentTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Deployment) validateStartedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("startedOn", "body", "date-time", m.StartedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Deployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deployment) UnmarshalBinary(b []byte) error {
	var res Deployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
