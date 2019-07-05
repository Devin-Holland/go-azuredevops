// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Operation Contains information about the progress or result of an async operation.
// swagger:model Operation
type Operation struct {
	OperationReference

	// Links to other related objects.
	Links *ReferenceLinks `json:"_links,omitempty"`

	// Detailed messaged about the status of an operation.
	DetailedMessage string `json:"detailedMessage,omitempty"`

	// Result message for an operation.
	ResultMessage string `json:"resultMessage,omitempty"`

	// URL to the operation result.
	ResultURL *OperationResultReference `json:"resultUrl,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Operation) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 OperationReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.OperationReference = aO0

	// now for regular properties
	var propsOperation struct {
		Links *ReferenceLinks `json:"_links,omitempty"`

		DetailedMessage string `json:"detailedMessage,omitempty"`

		ResultMessage string `json:"resultMessage,omitempty"`

		ResultURL *OperationResultReference `json:"resultUrl,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsOperation); err != nil {
		return err
	}
	m.Links = propsOperation.Links

	m.DetailedMessage = propsOperation.DetailedMessage

	m.ResultMessage = propsOperation.ResultMessage

	m.ResultURL = propsOperation.ResultURL

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Operation) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.OperationReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsOperation struct {
		Links *ReferenceLinks `json:"_links,omitempty"`

		DetailedMessage string `json:"detailedMessage,omitempty"`

		ResultMessage string `json:"resultMessage,omitempty"`

		ResultURL *OperationResultReference `json:"resultUrl,omitempty"`
	}
	propsOperation.Links = m.Links

	propsOperation.DetailedMessage = m.DetailedMessage

	propsOperation.ResultMessage = m.ResultMessage

	propsOperation.ResultURL = m.ResultURL

	jsonDataPropsOperation, errOperation := swag.WriteJSON(propsOperation)
	if errOperation != nil {
		return nil, errOperation
	}
	_parts = append(_parts, jsonDataPropsOperation)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this operation
func (m *Operation) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with OperationReference
	if err := m.OperationReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Operation) validateLinks(formats strfmt.Registry) error {

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

func (m *Operation) validateResultURL(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultURL) { // not required
		return nil
	}

	if m.ResultURL != nil {
		if err := m.ResultURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resultUrl")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Operation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Operation) UnmarshalBinary(b []byte) error {
	var res Operation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}