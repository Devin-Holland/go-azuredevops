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

// ExtensionManifest Base class for extension properties which are shared by the extension manifest and the extension model
// swagger:model ExtensionManifest
type ExtensionManifest struct {

	// Uri used as base for other relative uri's defined in extension
	BaseURI string `json:"baseUri,omitempty"`

	// List of shared constraints defined by this extension
	Constraints []*ContributionConstraint `json:"constraints"`

	// List of contribution types defined by this extension
	ContributionTypes []*ContributionType `json:"contributionTypes"`

	// List of contributions made by this extension
	Contributions []*Contribution `json:"contributions"`

	// List of explicit demands required by this extension
	Demands []string `json:"demands"`

	// Collection of endpoints that get called when particular extension events occur
	EventCallbacks *ExtensionEventCallbackCollection `json:"eventCallbacks,omitempty"`

	// Secondary location that can be used as base for other relative uri's defined in extension
	FallbackBaseURI string `json:"fallbackBaseUri,omitempty"`

	// Language Culture Name set by the Gallery
	Language string `json:"language,omitempty"`

	// How this extension behaves with respect to licensing
	Licensing *ExtensionLicensing `json:"licensing,omitempty"`

	// Version of the extension manifest format/content
	ManifestVersion float64 `json:"manifestVersion,omitempty"`

	// Default user claims applied to all contributions (except the ones which have been speficied restrictedTo explicitly) to control the visibility of a contribution.
	RestrictedTo []string `json:"restrictedTo"`

	// List of all oauth scopes required by this extension
	Scopes []string `json:"scopes"`

	// The ServiceInstanceType(Guid) of the VSTS service that must be available to an account in order for the extension to be installed
	// Format: uuid
	ServiceInstanceType strfmt.UUID `json:"serviceInstanceType,omitempty"`
}

// Validate validates this extension manifest
func (m *ExtensionManifest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContributionTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContributions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventCallbacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicensing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionManifest) validateConstraints(formats strfmt.Registry) error {

	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	for i := 0; i < len(m.Constraints); i++ {
		if swag.IsZero(m.Constraints[i]) { // not required
			continue
		}

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExtensionManifest) validateContributionTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ContributionTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ContributionTypes); i++ {
		if swag.IsZero(m.ContributionTypes[i]) { // not required
			continue
		}

		if m.ContributionTypes[i] != nil {
			if err := m.ContributionTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contributionTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExtensionManifest) validateContributions(formats strfmt.Registry) error {

	if swag.IsZero(m.Contributions) { // not required
		return nil
	}

	for i := 0; i < len(m.Contributions); i++ {
		if swag.IsZero(m.Contributions[i]) { // not required
			continue
		}

		if m.Contributions[i] != nil {
			if err := m.Contributions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contributions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExtensionManifest) validateEventCallbacks(formats strfmt.Registry) error {

	if swag.IsZero(m.EventCallbacks) { // not required
		return nil
	}

	if m.EventCallbacks != nil {
		if err := m.EventCallbacks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventCallbacks")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionManifest) validateLicensing(formats strfmt.Registry) error {

	if swag.IsZero(m.Licensing) { // not required
		return nil
	}

	if m.Licensing != nil {
		if err := m.Licensing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licensing")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionManifest) validateServiceInstanceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceInstanceType) { // not required
		return nil
	}

	if err := validate.FormatOf("serviceInstanceType", "body", "uuid", m.ServiceInstanceType.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionManifest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionManifest) UnmarshalBinary(b []byte) error {
	var res ExtensionManifest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}