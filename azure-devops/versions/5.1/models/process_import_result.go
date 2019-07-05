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

// ProcessImportResult Describes the result of a Process Import request.
// swagger:model ProcessImportResult
type ProcessImportResult struct {

	// Check template existence result.
	CheckExistenceResult *CheckTemplateExistenceResult `json:"checkExistenceResult,omitempty"`

	// Help URL.
	HelpURL string `json:"helpUrl,omitempty"`

	// ID of the import operation.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Whether this imported process is new.
	IsNew bool `json:"isNew,omitempty"`

	// The promote job identifier.
	// Format: uuid
	PromoteJobID strfmt.UUID `json:"promoteJobId,omitempty"`

	// The list of validation results.
	ValidationResults []*ValidationIssue `json:"validationResults"`
}

// Validate validates this process import result
func (m *ProcessImportResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckExistenceResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromoteJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessImportResult) validateCheckExistenceResult(formats strfmt.Registry) error {

	if swag.IsZero(m.CheckExistenceResult) { // not required
		return nil
	}

	if m.CheckExistenceResult != nil {
		if err := m.CheckExistenceResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkExistenceResult")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessImportResult) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProcessImportResult) validatePromoteJobID(formats strfmt.Registry) error {

	if swag.IsZero(m.PromoteJobID) { // not required
		return nil
	}

	if err := validate.FormatOf("promoteJobId", "body", "uuid", m.PromoteJobID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProcessImportResult) validateValidationResults(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationResults) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidationResults); i++ {
		if swag.IsZero(m.ValidationResults[i]) { // not required
			continue
		}

		if m.ValidationResults[i] != nil {
			if err := m.ValidationResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessImportResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessImportResult) UnmarshalBinary(b []byte) error {
	var res ProcessImportResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
