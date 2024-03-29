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

// BuildController build controller
// swagger:model BuildController
type BuildController struct {
	XamlBuildControllerReference

	// links
	Links *ReferenceLinks `json:"_links,omitempty"`

	// The date the controller was created.
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// The description of the controller.
	Description string `json:"description,omitempty"`

	// Indicates whether the controller is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// The status of the controller.
	// Enum: [unavailable available offline]
	Status interface{} `json:"status,omitempty"`

	// The date the controller was last updated.
	// Format: date-time
	UpdatedDate strfmt.DateTime `json:"updatedDate,omitempty"`

	// The controller's URI.
	URI string `json:"uri,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BuildController) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 XamlBuildControllerReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.XamlBuildControllerReference = aO0

	// now for regular properties
	var propsBuildController struct {
		Links *ReferenceLinks `json:"_links,omitempty"`

		CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

		Description string `json:"description,omitempty"`

		Enabled bool `json:"enabled,omitempty"`

		Status interface{} `json:"status,omitempty"`

		UpdatedDate strfmt.DateTime `json:"updatedDate,omitempty"`

		URI string `json:"uri,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsBuildController); err != nil {
		return err
	}
	m.Links = propsBuildController.Links

	m.CreatedDate = propsBuildController.CreatedDate

	m.Description = propsBuildController.Description

	m.Enabled = propsBuildController.Enabled

	m.Status = propsBuildController.Status

	m.UpdatedDate = propsBuildController.UpdatedDate

	m.URI = propsBuildController.URI

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BuildController) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.XamlBuildControllerReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsBuildController struct {
		Links *ReferenceLinks `json:"_links,omitempty"`

		CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

		Description string `json:"description,omitempty"`

		Enabled bool `json:"enabled,omitempty"`

		Status interface{} `json:"status,omitempty"`

		UpdatedDate strfmt.DateTime `json:"updatedDate,omitempty"`

		URI string `json:"uri,omitempty"`
	}
	propsBuildController.Links = m.Links

	propsBuildController.CreatedDate = m.CreatedDate

	propsBuildController.Description = m.Description

	propsBuildController.Enabled = m.Enabled

	propsBuildController.Status = m.Status

	propsBuildController.UpdatedDate = m.UpdatedDate

	propsBuildController.URI = m.URI

	jsonDataPropsBuildController, errBuildController := swag.WriteJSON(propsBuildController)
	if errBuildController != nil {
		return nil, errBuildController
	}
	_parts = append(_parts, jsonDataPropsBuildController)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this build controller
func (m *BuildController) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with XamlBuildControllerReference
	if err := m.XamlBuildControllerReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildController) validateLinks(formats strfmt.Registry) error {

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

func (m *BuildController) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildController) validateUpdatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedDate", "body", "date-time", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildController) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildController) UnmarshalBinary(b []byte) error {
	var res BuildController
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
