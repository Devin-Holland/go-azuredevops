// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DebugEntryCreateBatch A batch of debug entry to create.
// swagger:model DebugEntryCreateBatch
type DebugEntryCreateBatch struct {

	// Defines what to do when a debug entry in the batch already exists.
	// Enum: [throwIfExists skipIfExists overwriteIfExists]
	CreateBehavior interface{} `json:"createBehavior,omitempty"`

	// The debug entries.
	DebugEntries []*DebugEntry `json:"debugEntries"`
}

// Validate validates this debug entry create batch
func (m *DebugEntryCreateBatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDebugEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebugEntryCreateBatch) validateDebugEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.DebugEntries) { // not required
		return nil
	}

	for i := 0; i < len(m.DebugEntries); i++ {
		if swag.IsZero(m.DebugEntries[i]) { // not required
			continue
		}

		if m.DebugEntries[i] != nil {
			if err := m.DebugEntries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("debugEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DebugEntryCreateBatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebugEntryCreateBatch) UnmarshalBinary(b []byte) error {
	var res DebugEntryCreateBatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
