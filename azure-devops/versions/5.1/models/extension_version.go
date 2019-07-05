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

// ExtensionVersion extension version
// swagger:model ExtensionVersion
type ExtensionVersion struct {

	// asset Uri
	AssetURI string `json:"assetUri,omitempty"`

	// badges
	Badges []*ExtensionBadge `json:"badges"`

	// fallback asset Uri
	FallbackAssetURI string `json:"fallbackAssetUri,omitempty"`

	// files
	Files []*ExtensionFile `json:"files"`

	// flags
	// Enum: [none validated]
	Flags interface{} `json:"flags,omitempty"`

	// last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// properties
	Properties []map[string]string `json:"properties"`

	// validation result message
	ValidationResultMessage string `json:"validationResultMessage,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// version description
	VersionDescription string `json:"versionDescription,omitempty"`
}

// Validate validates this extension version
func (m *ExtensionVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBadges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionVersion) validateBadges(formats strfmt.Registry) error {

	if swag.IsZero(m.Badges) { // not required
		return nil
	}

	for i := 0; i < len(m.Badges); i++ {
		if swag.IsZero(m.Badges[i]) { // not required
			continue
		}

		if m.Badges[i] != nil {
			if err := m.Badges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("badges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExtensionVersion) validateFiles(formats strfmt.Registry) error {

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

func (m *ExtensionVersion) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionVersion) UnmarshalBinary(b []byte) error {
	var res ExtensionVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}