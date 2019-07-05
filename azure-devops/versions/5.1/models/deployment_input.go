// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeploymentInput deployment input
// swagger:model DeploymentInput
type DeploymentInput struct {
	BaseDeploymentInput

	// Artifacts that downloaded during job execution.
	ArtifactsDownloadInput *ArtifactsDownloadInput `json:"artifactsDownloadInput,omitempty"`

	// List demands that needs to meet to execute the job.
	Demands []*Demand `json:"demands"`

	// Indicates whether to include access token in deployment job or not.
	EnableAccessToken bool `json:"enableAccessToken,omitempty"`

	// Id of the pool on which job get executed.
	QueueID int32 `json:"queueId,omitempty"`

	// Indicates whether artifacts downloaded while job execution or not.
	SkipArtifactsDownload bool `json:"skipArtifactsDownload,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DeploymentInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseDeploymentInput
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseDeploymentInput = aO0

	// now for regular properties
	var propsDeploymentInput struct {
		ArtifactsDownloadInput *ArtifactsDownloadInput `json:"artifactsDownloadInput,omitempty"`

		Demands []*Demand `json:"demands"`

		EnableAccessToken bool `json:"enableAccessToken,omitempty"`

		QueueID int32 `json:"queueId,omitempty"`

		SkipArtifactsDownload bool `json:"skipArtifactsDownload,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsDeploymentInput); err != nil {
		return err
	}
	m.ArtifactsDownloadInput = propsDeploymentInput.ArtifactsDownloadInput

	m.Demands = propsDeploymentInput.Demands

	m.EnableAccessToken = propsDeploymentInput.EnableAccessToken

	m.QueueID = propsDeploymentInput.QueueID

	m.SkipArtifactsDownload = propsDeploymentInput.SkipArtifactsDownload

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DeploymentInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BaseDeploymentInput)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsDeploymentInput struct {
		ArtifactsDownloadInput *ArtifactsDownloadInput `json:"artifactsDownloadInput,omitempty"`

		Demands []*Demand `json:"demands"`

		EnableAccessToken bool `json:"enableAccessToken,omitempty"`

		QueueID int32 `json:"queueId,omitempty"`

		SkipArtifactsDownload bool `json:"skipArtifactsDownload,omitempty"`
	}
	propsDeploymentInput.ArtifactsDownloadInput = m.ArtifactsDownloadInput

	propsDeploymentInput.Demands = m.Demands

	propsDeploymentInput.EnableAccessToken = m.EnableAccessToken

	propsDeploymentInput.QueueID = m.QueueID

	propsDeploymentInput.SkipArtifactsDownload = m.SkipArtifactsDownload

	jsonDataPropsDeploymentInput, errDeploymentInput := swag.WriteJSON(propsDeploymentInput)
	if errDeploymentInput != nil {
		return nil, errDeploymentInput
	}
	_parts = append(_parts, jsonDataPropsDeploymentInput)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this deployment input
func (m *DeploymentInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseDeploymentInput
	if err := m.BaseDeploymentInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtifactsDownloadInput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDemands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentInput) validateArtifactsDownloadInput(formats strfmt.Registry) error {

	if swag.IsZero(m.ArtifactsDownloadInput) { // not required
		return nil
	}

	if m.ArtifactsDownloadInput != nil {
		if err := m.ArtifactsDownloadInput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactsDownloadInput")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentInput) validateDemands(formats strfmt.Registry) error {

	if swag.IsZero(m.Demands) { // not required
		return nil
	}

	for i := 0; i < len(m.Demands); i++ {
		if swag.IsZero(m.Demands[i]) { // not required
			continue
		}

		if m.Demands[i] != nil {
			if err := m.Demands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("demands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentInput) UnmarshalBinary(b []byte) error {
	var res DeploymentInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}