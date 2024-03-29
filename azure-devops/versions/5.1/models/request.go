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

// Request Symbol request.
// swagger:model Request
type Request struct {
	ResourceBase

	// An optional human-facing description.
	Description string `json:"description,omitempty"`

	// An optional expiration date for the request. The request will become inaccessible and get deleted after the date, regardless of its status.  On an HTTP POST, if expiration date is null/missing, the server will assign a default expiration data (30 days unless overwridden in the registry at the account level). On PATCH, if expiration date is null/missing, the behavior is to not change whatever the request's current expiration date is.
	// Format: date-time
	ExpirationDate strfmt.DateTime `json:"expirationDate,omitempty"`

	// A human-facing name for the request. Required on POST, ignored on PATCH.
	Name string `json:"name,omitempty"`

	// The status for this request.
	// Enum: [none created sealed unavailable]
	Status interface{} `json:"status,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Request) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceBase = aO0

	// now for regular properties
	var propsRequest struct {
		Description string `json:"description,omitempty"`

		ExpirationDate strfmt.DateTime `json:"expirationDate,omitempty"`

		Name string `json:"name,omitempty"`

		Status interface{} `json:"status,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsRequest); err != nil {
		return err
	}
	m.Description = propsRequest.Description

	m.ExpirationDate = propsRequest.ExpirationDate

	m.Name = propsRequest.Name

	m.Status = propsRequest.Status

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Request) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ResourceBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsRequest struct {
		Description string `json:"description,omitempty"`

		ExpirationDate strfmt.DateTime `json:"expirationDate,omitempty"`

		Name string `json:"name,omitempty"`

		Status interface{} `json:"status,omitempty"`
	}
	propsRequest.Description = m.Description

	propsRequest.ExpirationDate = m.ExpirationDate

	propsRequest.Name = m.Name

	propsRequest.Status = m.Status

	jsonDataPropsRequest, errRequest := swag.WriteJSON(propsRequest)
	if errRequest != nil {
		return nil, errRequest
	}
	_parts = append(_parts, jsonDataPropsRequest)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this request
func (m *Request) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceBase
	if err := m.ResourceBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Request) validateExpirationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Request) UnmarshalBinary(b []byte) error {
	var res Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
