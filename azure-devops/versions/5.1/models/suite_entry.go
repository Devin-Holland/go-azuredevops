// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SuiteEntry A suite entry defines properties for a test suite.
// swagger:model SuiteEntry
type SuiteEntry struct {
	SuiteEntryUpdateParams

	// Id for the test suite.
	SuiteID int32 `json:"suiteId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SuiteEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SuiteEntryUpdateParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SuiteEntryUpdateParams = aO0

	// now for regular properties
	var propsSuiteEntry struct {
		SuiteID int32 `json:"suiteId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsSuiteEntry); err != nil {
		return err
	}
	m.SuiteID = propsSuiteEntry.SuiteID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SuiteEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SuiteEntryUpdateParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsSuiteEntry struct {
		SuiteID int32 `json:"suiteId,omitempty"`
	}
	propsSuiteEntry.SuiteID = m.SuiteID

	jsonDataPropsSuiteEntry, errSuiteEntry := swag.WriteJSON(propsSuiteEntry)
	if errSuiteEntry != nil {
		return nil, errSuiteEntry
	}
	_parts = append(_parts, jsonDataPropsSuiteEntry)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this suite entry
func (m *SuiteEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SuiteEntryUpdateParams
	if err := m.SuiteEntryUpdateParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SuiteEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuiteEntry) UnmarshalBinary(b []byte) error {
	var res SuiteEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}