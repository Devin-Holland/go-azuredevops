// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WorkItemTypeTemplateUpdateModel Describes a update work item type template request body.
// swagger:model WorkItemTypeTemplateUpdateModel
type WorkItemTypeTemplateUpdateModel struct {

	// Describes the type of the action for the update request.
	// Enum: [import validate]
	ActionType interface{} `json:"actionType,omitempty"`

	// Methodology to which the template belongs, eg. Agile, Scrum, CMMI.
	Methodology string `json:"methodology,omitempty"`

	// String representation of the work item type template.
	Template string `json:"template,omitempty"`

	// The type of the template described in the request body.
	// Enum: [workItemType globalWorkflow]
	TemplateType interface{} `json:"templateType,omitempty"`
}

// Validate validates this work item type template update model
func (m *WorkItemTypeTemplateUpdateModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemTypeTemplateUpdateModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemTypeTemplateUpdateModel) UnmarshalBinary(b []byte) error {
	var res WorkItemTypeTemplateUpdateModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}