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

// XamlBuildDefinition xaml build definition
// swagger:model XamlBuildDefinition
type XamlBuildDefinition struct {
	DefinitionReference

	// links
	Links *ReferenceLinks `json:"_links,omitempty"`

	// Batch size of the definition
	BatchSize int32 `json:"batchSize,omitempty"`

	// build args
	BuildArgs string `json:"buildArgs,omitempty"`

	// The continuous integration quiet period
	ContinuousIntegrationQuietPeriod int32 `json:"continuousIntegrationQuietPeriod,omitempty"`

	// The build controller
	Controller *BuildController `json:"controller,omitempty"`

	// The date this definition was created
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

	// Default drop location for builds from this definition
	DefaultDropLocation string `json:"defaultDropLocation,omitempty"`

	// Description of the definition
	Description string `json:"description,omitempty"`

	// The last build on this definition
	LastBuild *XamlBuildReference `json:"lastBuild,omitempty"`

	// The repository
	Repository *BuildRepository `json:"repository,omitempty"`

	// The reasons supported by the template
	// Enum: [none manual individualCI batchedCI schedule scheduleForced userCreated validateShelveset checkInShelveset pullRequest buildCompletion triggered all]
	SupportedReasons interface{} `json:"supportedReasons,omitempty"`

	// How builds are triggered from this definition
	// Enum: [none continuousIntegration batchedContinuousIntegration schedule gatedCheckIn batchedGatedCheckIn pullRequest buildCompletion all]
	TriggerType interface{} `json:"triggerType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *XamlBuildDefinition) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DefinitionReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DefinitionReference = aO0

	// now for regular properties
	var propsXamlBuildDefinition struct {
		Links *ReferenceLinks `json:"_links,omitempty"`

		BatchSize int32 `json:"batchSize,omitempty"`

		BuildArgs string `json:"buildArgs,omitempty"`

		ContinuousIntegrationQuietPeriod int32 `json:"continuousIntegrationQuietPeriod,omitempty"`

		Controller *BuildController `json:"controller,omitempty"`

		CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

		DefaultDropLocation string `json:"defaultDropLocation,omitempty"`

		Description string `json:"description,omitempty"`

		LastBuild *XamlBuildReference `json:"lastBuild,omitempty"`

		Repository *BuildRepository `json:"repository,omitempty"`

		SupportedReasons interface{} `json:"supportedReasons,omitempty"`

		TriggerType interface{} `json:"triggerType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsXamlBuildDefinition); err != nil {
		return err
	}
	m.Links = propsXamlBuildDefinition.Links

	m.BatchSize = propsXamlBuildDefinition.BatchSize

	m.BuildArgs = propsXamlBuildDefinition.BuildArgs

	m.ContinuousIntegrationQuietPeriod = propsXamlBuildDefinition.ContinuousIntegrationQuietPeriod

	m.Controller = propsXamlBuildDefinition.Controller

	m.CreatedOn = propsXamlBuildDefinition.CreatedOn

	m.DefaultDropLocation = propsXamlBuildDefinition.DefaultDropLocation

	m.Description = propsXamlBuildDefinition.Description

	m.LastBuild = propsXamlBuildDefinition.LastBuild

	m.Repository = propsXamlBuildDefinition.Repository

	m.SupportedReasons = propsXamlBuildDefinition.SupportedReasons

	m.TriggerType = propsXamlBuildDefinition.TriggerType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m XamlBuildDefinition) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.DefinitionReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsXamlBuildDefinition struct {
		Links *ReferenceLinks `json:"_links,omitempty"`

		BatchSize int32 `json:"batchSize,omitempty"`

		BuildArgs string `json:"buildArgs,omitempty"`

		ContinuousIntegrationQuietPeriod int32 `json:"continuousIntegrationQuietPeriod,omitempty"`

		Controller *BuildController `json:"controller,omitempty"`

		CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

		DefaultDropLocation string `json:"defaultDropLocation,omitempty"`

		Description string `json:"description,omitempty"`

		LastBuild *XamlBuildReference `json:"lastBuild,omitempty"`

		Repository *BuildRepository `json:"repository,omitempty"`

		SupportedReasons interface{} `json:"supportedReasons,omitempty"`

		TriggerType interface{} `json:"triggerType,omitempty"`
	}
	propsXamlBuildDefinition.Links = m.Links

	propsXamlBuildDefinition.BatchSize = m.BatchSize

	propsXamlBuildDefinition.BuildArgs = m.BuildArgs

	propsXamlBuildDefinition.ContinuousIntegrationQuietPeriod = m.ContinuousIntegrationQuietPeriod

	propsXamlBuildDefinition.Controller = m.Controller

	propsXamlBuildDefinition.CreatedOn = m.CreatedOn

	propsXamlBuildDefinition.DefaultDropLocation = m.DefaultDropLocation

	propsXamlBuildDefinition.Description = m.Description

	propsXamlBuildDefinition.LastBuild = m.LastBuild

	propsXamlBuildDefinition.Repository = m.Repository

	propsXamlBuildDefinition.SupportedReasons = m.SupportedReasons

	propsXamlBuildDefinition.TriggerType = m.TriggerType

	jsonDataPropsXamlBuildDefinition, errXamlBuildDefinition := swag.WriteJSON(propsXamlBuildDefinition)
	if errXamlBuildDefinition != nil {
		return nil, errXamlBuildDefinition
	}
	_parts = append(_parts, jsonDataPropsXamlBuildDefinition)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this xaml build definition
func (m *XamlBuildDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DefinitionReference
	if err := m.DefinitionReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateController(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *XamlBuildDefinition) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *XamlBuildDefinition) validateController(formats strfmt.Registry) error {

	if swag.IsZero(m.Controller) { // not required
		return nil
	}

	if m.Controller != nil {
		if err := m.Controller.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controller")
			}
			return err
		}
	}

	return nil
}

func (m *XamlBuildDefinition) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("createdOn", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *XamlBuildDefinition) validateLastBuild(formats strfmt.Registry) error {

	if swag.IsZero(m.LastBuild) { // not required
		return nil
	}

	if m.LastBuild != nil {
		if err := m.LastBuild.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastBuild")
			}
			return err
		}
	}

	return nil
}

func (m *XamlBuildDefinition) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *XamlBuildDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *XamlBuildDefinition) UnmarshalBinary(b []byte) error {
	var res XamlBuildDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}