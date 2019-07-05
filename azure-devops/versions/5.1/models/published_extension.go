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

// PublishedExtension published extension
// swagger:model PublishedExtension
type PublishedExtension struct {

	// categories
	Categories []string `json:"categories"`

	// deployment type
	// Enum: [exe msi vsix referralLink]
	DeploymentType interface{} `json:"deploymentType,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// extension Id
	// Format: uuid
	ExtensionID strfmt.UUID `json:"extensionId,omitempty"`

	// extension name
	ExtensionName string `json:"extensionName,omitempty"`

	// flags
	// Enum: [none disabled builtIn validated trusted paid public multiVersion system preview unpublished trial locked hidden]
	Flags interface{} `json:"flags,omitempty"`

	// installation targets
	InstallationTargets []*InstallationTarget `json:"installationTargets"`

	// last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// long description
	LongDescription string `json:"longDescription,omitempty"`

	// Date on which the extension was first uploaded.
	// Format: date-time
	PublishedDate strfmt.DateTime `json:"publishedDate,omitempty"`

	// publisher
	Publisher *PublisherFacts `json:"publisher,omitempty"`

	// Date on which the extension first went public.
	// Format: date-time
	ReleaseDate strfmt.DateTime `json:"releaseDate,omitempty"`

	// shared with
	SharedWith []*ExtensionShare `json:"sharedWith"`

	// short description
	ShortDescription string `json:"shortDescription,omitempty"`

	// statistics
	Statistics []*ExtensionStatistic `json:"statistics"`

	// tags
	Tags []string `json:"tags"`

	// versions
	Versions []*ExtensionVersion `json:"versions"`
}

// Validate validates this published extension
func (m *PublishedExtension) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallationTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublisher(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublishedExtension) validateExtensionID(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtensionID) { // not required
		return nil
	}

	if err := validate.FormatOf("extensionId", "body", "uuid", m.ExtensionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublishedExtension) validateInstallationTargets(formats strfmt.Registry) error {

	if swag.IsZero(m.InstallationTargets) { // not required
		return nil
	}

	for i := 0; i < len(m.InstallationTargets); i++ {
		if swag.IsZero(m.InstallationTargets[i]) { // not required
			continue
		}

		if m.InstallationTargets[i] != nil {
			if err := m.InstallationTargets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("installationTargets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PublishedExtension) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublishedExtension) validatePublishedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("publishedDate", "body", "date-time", m.PublishedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublishedExtension) validatePublisher(formats strfmt.Registry) error {

	if swag.IsZero(m.Publisher) { // not required
		return nil
	}

	if m.Publisher != nil {
		if err := m.Publisher.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publisher")
			}
			return err
		}
	}

	return nil
}

func (m *PublishedExtension) validateReleaseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("releaseDate", "body", "date-time", m.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublishedExtension) validateSharedWith(formats strfmt.Registry) error {

	if swag.IsZero(m.SharedWith) { // not required
		return nil
	}

	for i := 0; i < len(m.SharedWith); i++ {
		if swag.IsZero(m.SharedWith[i]) { // not required
			continue
		}

		if m.SharedWith[i] != nil {
			if err := m.SharedWith[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sharedWith" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PublishedExtension) validateStatistics(formats strfmt.Registry) error {

	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	for i := 0; i < len(m.Statistics); i++ {
		if swag.IsZero(m.Statistics[i]) { // not required
			continue
		}

		if m.Statistics[i] != nil {
			if err := m.Statistics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statistics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PublishedExtension) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	for i := 0; i < len(m.Versions); i++ {
		if swag.IsZero(m.Versions[i]) { // not required
			continue
		}

		if m.Versions[i] != nil {
			if err := m.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublishedExtension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublishedExtension) UnmarshalBinary(b []byte) error {
	var res PublishedExtension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}