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

// Widget Widget data
// swagger:model Widget
type Widget struct {

	// links
	Links *ReferenceLinks `json:"_links,omitempty"`

	// Refers to the allowed sizes for the widget. This gets populated when user wants to configure the widget
	AllowedSizes []*WidgetSize `json:"allowedSizes"`

	// Read-Only Property from Dashboard Service. Indicates if settings are blocked for the current user.
	AreSettingsBlockedForUser bool `json:"areSettingsBlockedForUser,omitempty"`

	// Refers to unique identifier of a feature artifact. Used for pinning+unpinning a specific artifact.
	ArtifactID string `json:"artifactId,omitempty"`

	// configuration contribution Id
	ConfigurationContributionID string `json:"configurationContributionId,omitempty"`

	// configuration contribution relative Id
	ConfigurationContributionRelativeID string `json:"configurationContributionRelativeId,omitempty"`

	// content Uri
	ContentURI string `json:"contentUri,omitempty"`

	// The id of the underlying contribution defining the supplied Widget Configuration.
	ContributionID string `json:"contributionId,omitempty"`

	// Optional partial dashboard content, to support exchanging dashboard-level version ETag for widget-level APIs
	Dashboard *Dashboard `json:"dashboard,omitempty"`

	// e tag
	ETag string `json:"eTag,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// is enabled
	IsEnabled bool `json:"isEnabled,omitempty"`

	// is name configurable
	IsNameConfigurable bool `json:"isNameConfigurable,omitempty"`

	// lightbox options
	LightboxOptions *LightboxOptions `json:"lightboxOptions,omitempty"`

	// loading image Url
	LoadingImageURL string `json:"loadingImageUrl,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// position
	Position *WidgetPosition `json:"position,omitempty"`

	// settings
	Settings string `json:"settings,omitempty"`

	// settings version
	SettingsVersion *SemanticVersion `json:"settingsVersion,omitempty"`

	// size
	Size *WidgetSize `json:"size,omitempty"`

	// type Id
	TypeID string `json:"typeId,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this widget
func (m *Widget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowedSizes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLightboxOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettingsVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Widget) validateLinks(formats strfmt.Registry) error {

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

func (m *Widget) validateAllowedSizes(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowedSizes) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowedSizes); i++ {
		if swag.IsZero(m.AllowedSizes[i]) { // not required
			continue
		}

		if m.AllowedSizes[i] != nil {
			if err := m.AllowedSizes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedSizes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Widget) validateDashboard(formats strfmt.Registry) error {

	if swag.IsZero(m.Dashboard) { // not required
		return nil
	}

	if m.Dashboard != nil {
		if err := m.Dashboard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard")
			}
			return err
		}
	}

	return nil
}

func (m *Widget) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Widget) validateLightboxOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.LightboxOptions) { // not required
		return nil
	}

	if m.LightboxOptions != nil {
		if err := m.LightboxOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lightboxOptions")
			}
			return err
		}
	}

	return nil
}

func (m *Widget) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *Widget) validateSettingsVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.SettingsVersion) { // not required
		return nil
	}

	if m.SettingsVersion != nil {
		if err := m.SettingsVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settingsVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Widget) validateSize(formats strfmt.Registry) error {

	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Widget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Widget) UnmarshalBinary(b []byte) error {
	var res Widget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
