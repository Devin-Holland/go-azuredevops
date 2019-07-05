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

// QueryBatchGetRequest Describes a request to get a list of queries
// swagger:model QueryBatchGetRequest
type QueryBatchGetRequest struct {

	// The expand parameters for queries. Possible options are { None, Wiql, Clauses, All, Minimal }
	// Enum: [none wiql clauses all minimal]
	NrDollarExpand interface{} `json:"$expand,omitempty"`

	// The flag to control error policy in a query batch request. Possible options are { Fail, Omit }.
	// Enum: [fail omit]
	ErrorPolicy interface{} `json:"errorPolicy,omitempty"`

	// The requested query ids
	Ids []strfmt.UUID `json:"ids"`
}

// Validate validates this query batch get request
func (m *QueryBatchGetRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryBatchGetRequest) validateIds(formats strfmt.Registry) error {

	if swag.IsZero(m.Ids) { // not required
		return nil
	}

	for i := 0; i < len(m.Ids); i++ {

		if err := validate.FormatOf("ids"+"."+strconv.Itoa(i), "body", "uuid", m.Ids[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryBatchGetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryBatchGetRequest) UnmarshalBinary(b []byte) error {
	var res QueryBatchGetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
