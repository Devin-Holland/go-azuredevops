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

// BuildAgent build agent
// swagger:model BuildAgent
type BuildAgent struct {

	// build directory
	BuildDirectory string `json:"buildDirectory,omitempty"`

	// controller
	Controller *XamlBuildControllerReference `json:"controller,omitempty"`

	// created date
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// message queue Url
	MessageQueueURL string `json:"messageQueueUrl,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// reserved for build
	ReservedForBuild string `json:"reservedForBuild,omitempty"`

	// server
	Server *XamlBuildServerReference `json:"server,omitempty"`

	// status
	// Enum: [unavailable available offline]
	Status interface{} `json:"status,omitempty"`

	// status message
	StatusMessage string `json:"statusMessage,omitempty"`

	// updated date
	// Format: date-time
	UpdatedDate strfmt.DateTime `json:"updatedDate,omitempty"`

	// uri
	URI string `json:"uri,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this build agent
func (m *BuildAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateController(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
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

func (m *BuildAgent) validateController(formats strfmt.Registry) error {

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

func (m *BuildAgent) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildAgent) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *BuildAgent) validateUpdatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedDate", "body", "date-time", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildAgent) UnmarshalBinary(b []byte) error {
	var res BuildAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}