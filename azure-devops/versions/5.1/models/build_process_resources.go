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

// BuildProcessResources Represents resources used by a build process.
// swagger:model BuildProcessResources
type BuildProcessResources struct {

	// endpoints
	Endpoints []*ServiceEndpointReference `json:"endpoints"`

	// files
	Files []*SecureFileReference `json:"files"`

	// queues
	Queues []*AgentPoolQueueReference `json:"queues"`

	// variable groups
	VariableGroups []*VariableGroupReference `json:"variableGroups"`
}

// Validate validates this build process resources
func (m *BuildProcessResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariableGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildProcessResources) validateEndpoints(formats strfmt.Registry) error {

	if swag.IsZero(m.Endpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.Endpoints); i++ {
		if swag.IsZero(m.Endpoints[i]) { // not required
			continue
		}

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuildProcessResources) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuildProcessResources) validateQueues(formats strfmt.Registry) error {

	if swag.IsZero(m.Queues) { // not required
		return nil
	}

	for i := 0; i < len(m.Queues); i++ {
		if swag.IsZero(m.Queues[i]) { // not required
			continue
		}

		if m.Queues[i] != nil {
			if err := m.Queues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuildProcessResources) validateVariableGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.VariableGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.VariableGroups); i++ {
		if swag.IsZero(m.VariableGroups[i]) { // not required
			continue
		}

		if m.VariableGroups[i] != nil {
			if err := m.VariableGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variableGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildProcessResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildProcessResources) UnmarshalBinary(b []byte) error {
	var res BuildProcessResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
