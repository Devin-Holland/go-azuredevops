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

// ProcessConfiguration Process Configurations for the project
// swagger:model ProcessConfiguration
type ProcessConfiguration struct {

	// Details about bug work items
	BugWorkItems *CategoryConfiguration `json:"bugWorkItems,omitempty"`

	// Details about portfolio backlogs
	PortfolioBacklogs []*CategoryConfiguration `json:"portfolioBacklogs"`

	// Details of requirement backlog
	RequirementBacklog *CategoryConfiguration `json:"requirementBacklog,omitempty"`

	// Details of task backlog
	TaskBacklog *CategoryConfiguration `json:"taskBacklog,omitempty"`

	// Type fields for the process configuration
	TypeFields map[string]WorkItemFieldReference `json:"typeFields,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this process configuration
func (m *ProcessConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBugWorkItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortfolioBacklogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequirementBacklog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskBacklog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessConfiguration) validateBugWorkItems(formats strfmt.Registry) error {

	if swag.IsZero(m.BugWorkItems) { // not required
		return nil
	}

	if m.BugWorkItems != nil {
		if err := m.BugWorkItems.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bugWorkItems")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessConfiguration) validatePortfolioBacklogs(formats strfmt.Registry) error {

	if swag.IsZero(m.PortfolioBacklogs) { // not required
		return nil
	}

	for i := 0; i < len(m.PortfolioBacklogs); i++ {
		if swag.IsZero(m.PortfolioBacklogs[i]) { // not required
			continue
		}

		if m.PortfolioBacklogs[i] != nil {
			if err := m.PortfolioBacklogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portfolioBacklogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProcessConfiguration) validateRequirementBacklog(formats strfmt.Registry) error {

	if swag.IsZero(m.RequirementBacklog) { // not required
		return nil
	}

	if m.RequirementBacklog != nil {
		if err := m.RequirementBacklog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requirementBacklog")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessConfiguration) validateTaskBacklog(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskBacklog) { // not required
		return nil
	}

	if m.TaskBacklog != nil {
		if err := m.TaskBacklog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskBacklog")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessConfiguration) validateTypeFields(formats strfmt.Registry) error {

	if swag.IsZero(m.TypeFields) { // not required
		return nil
	}

	for k := range m.TypeFields {

		if err := validate.Required("typeFields"+"."+k, "body", m.TypeFields[k]); err != nil {
			return err
		}
		if val, ok := m.TypeFields[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessConfiguration) UnmarshalBinary(b []byte) error {
	var res ProcessConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}