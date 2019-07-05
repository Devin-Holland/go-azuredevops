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

// SaveCachedPackagesData save cached packages data
// swagger:model SaveCachedPackagesData
type SaveCachedPackagesData struct {
	FeedBatchOperationData

	// normalized package names
	NormalizedPackageNames []string `json:"normalizedPackageNames"`

	// views for promotion
	ViewsForPromotion []strfmt.UUID `json:"viewsForPromotion"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SaveCachedPackagesData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 FeedBatchOperationData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.FeedBatchOperationData = aO0

	// now for regular properties
	var propsSaveCachedPackagesData struct {
		NormalizedPackageNames []string `json:"normalizedPackageNames"`

		ViewsForPromotion []strfmt.UUID `json:"viewsForPromotion"`
	}
	if err := swag.ReadJSON(raw, &propsSaveCachedPackagesData); err != nil {
		return err
	}
	m.NormalizedPackageNames = propsSaveCachedPackagesData.NormalizedPackageNames

	m.ViewsForPromotion = propsSaveCachedPackagesData.ViewsForPromotion

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SaveCachedPackagesData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.FeedBatchOperationData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsSaveCachedPackagesData struct {
		NormalizedPackageNames []string `json:"normalizedPackageNames"`

		ViewsForPromotion []strfmt.UUID `json:"viewsForPromotion"`
	}
	propsSaveCachedPackagesData.NormalizedPackageNames = m.NormalizedPackageNames

	propsSaveCachedPackagesData.ViewsForPromotion = m.ViewsForPromotion

	jsonDataPropsSaveCachedPackagesData, errSaveCachedPackagesData := swag.WriteJSON(propsSaveCachedPackagesData)
	if errSaveCachedPackagesData != nil {
		return nil, errSaveCachedPackagesData
	}
	_parts = append(_parts, jsonDataPropsSaveCachedPackagesData)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this save cached packages data
func (m *SaveCachedPackagesData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FeedBatchOperationData

	if err := m.validateViewsForPromotion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SaveCachedPackagesData) validateViewsForPromotion(formats strfmt.Registry) error {

	if swag.IsZero(m.ViewsForPromotion) { // not required
		return nil
	}

	for i := 0; i < len(m.ViewsForPromotion); i++ {

		if err := validate.FormatOf("viewsForPromotion"+"."+strconv.Itoa(i), "body", "uuid", m.ViewsForPromotion[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SaveCachedPackagesData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SaveCachedPackagesData) UnmarshalBinary(b []byte) error {
	var res SaveCachedPackagesData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
