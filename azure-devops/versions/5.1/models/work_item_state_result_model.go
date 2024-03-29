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

// WorkItemStateResultModel Class that represents a work item state result.
// swagger:model WorkItemStateResultModel
type WorkItemStateResultModel struct {

	// Work item state color.
	Color string `json:"color,omitempty"`

	// Work item state customization type.
	// Enum: [system inherited custom]
	CustomizationType interface{} `json:"customizationType,omitempty"`

	// If the Work item state is hidden.
	Hidden bool `json:"hidden,omitempty"`

	// Id of the Workitemstate.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Work item state name.
	Name string `json:"name,omitempty"`

	// Work item state order.
	Order int32 `json:"order,omitempty"`

	// Work item state statecategory.
	StateCategory string `json:"stateCategory,omitempty"`

	// Work item state url.
	URL string `json:"url,omitempty"`
}

// Validate validates this work item state result model
func (m *WorkItemStateResultModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkItemStateResultModel) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemStateResultModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemStateResultModel) UnmarshalBinary(b []byte) error {
	var res WorkItemStateResultModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
