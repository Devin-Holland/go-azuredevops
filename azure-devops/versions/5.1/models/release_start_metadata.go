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

// ReleaseStartMetadata release start metadata
// swagger:model ReleaseStartMetadata
type ReleaseStartMetadata struct {

	// Sets list of artifact to create a release.
	Artifacts []*ArtifactMetadata `json:"artifacts"`

	// Sets definition Id to create a release.
	DefinitionID int32 `json:"definitionId,omitempty"`

	// Sets description to create a release.
	Description string `json:"description,omitempty"`

	// Sets list of environments meta data.
	EnvironmentsMetadata []*ReleaseStartEnvironmentMetadata `json:"environmentsMetadata"`

	// Sets 'true' to create release in draft mode, 'false' otherwise.
	IsDraft bool `json:"isDraft,omitempty"`

	// Sets list of environments to manual as condition.
	ManualEnvironments []string `json:"manualEnvironments"`

	// properties
	Properties *PropertiesCollection `json:"properties,omitempty"`

	// Sets reason to create a release.
	// Enum: [none manual continuousIntegration schedule pullRequest]
	Reason interface{} `json:"reason,omitempty"`

	// Sets list of release variables to be overridden at deployment time.
	Variables map[string]ConfigurationVariableValue `json:"variables,omitempty"`
}

// Validate validates this release start metadata
func (m *ReleaseStartMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentsMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseStartMetadata) validateArtifacts(formats strfmt.Registry) error {

	if swag.IsZero(m.Artifacts) { // not required
		return nil
	}

	for i := 0; i < len(m.Artifacts); i++ {
		if swag.IsZero(m.Artifacts[i]) { // not required
			continue
		}

		if m.Artifacts[i] != nil {
			if err := m.Artifacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseStartMetadata) validateEnvironmentsMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentsMetadata) { // not required
		return nil
	}

	for i := 0; i < len(m.EnvironmentsMetadata); i++ {
		if swag.IsZero(m.EnvironmentsMetadata[i]) { // not required
			continue
		}

		if m.EnvironmentsMetadata[i] != nil {
			if err := m.EnvironmentsMetadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environmentsMetadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseStartMetadata) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseStartMetadata) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for k := range m.Variables {

		if err := validate.Required("variables"+"."+k, "body", m.Variables[k]); err != nil {
			return err
		}
		if val, ok := m.Variables[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseStartMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseStartMetadata) UnmarshalBinary(b []byte) error {
	var res ReleaseStartMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}