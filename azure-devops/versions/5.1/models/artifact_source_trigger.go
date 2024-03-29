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

// ArtifactSourceTrigger artifact source trigger
// swagger:model ArtifactSourceTrigger
type ArtifactSourceTrigger struct {
	ReleaseTriggerBase

	// Artifact source alias for Artifact Source trigger type
	ArtifactAlias string `json:"artifactAlias,omitempty"`

	// trigger conditions
	TriggerConditions []*ArtifactFilter `json:"triggerConditions"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ArtifactSourceTrigger) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ReleaseTriggerBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ReleaseTriggerBase = aO0

	// now for regular properties
	var propsArtifactSourceTrigger struct {
		ArtifactAlias string `json:"artifactAlias,omitempty"`

		TriggerConditions []*ArtifactFilter `json:"triggerConditions"`
	}
	if err := swag.ReadJSON(raw, &propsArtifactSourceTrigger); err != nil {
		return err
	}
	m.ArtifactAlias = propsArtifactSourceTrigger.ArtifactAlias

	m.TriggerConditions = propsArtifactSourceTrigger.TriggerConditions

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ArtifactSourceTrigger) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ReleaseTriggerBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsArtifactSourceTrigger struct {
		ArtifactAlias string `json:"artifactAlias,omitempty"`

		TriggerConditions []*ArtifactFilter `json:"triggerConditions"`
	}
	propsArtifactSourceTrigger.ArtifactAlias = m.ArtifactAlias

	propsArtifactSourceTrigger.TriggerConditions = m.TriggerConditions

	jsonDataPropsArtifactSourceTrigger, errArtifactSourceTrigger := swag.WriteJSON(propsArtifactSourceTrigger)
	if errArtifactSourceTrigger != nil {
		return nil, errArtifactSourceTrigger
	}
	_parts = append(_parts, jsonDataPropsArtifactSourceTrigger)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this artifact source trigger
func (m *ArtifactSourceTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ReleaseTriggerBase
	if err := m.ReleaseTriggerBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggerConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArtifactSourceTrigger) validateTriggerConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.TriggerConditions); i++ {
		if swag.IsZero(m.TriggerConditions[i]) { // not required
			continue
		}

		if m.TriggerConditions[i] != nil {
			if err := m.TriggerConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("triggerConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactSourceTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactSourceTrigger) UnmarshalBinary(b []byte) error {
	var res ArtifactSourceTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
