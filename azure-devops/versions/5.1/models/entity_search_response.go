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

// EntitySearchResponse Defines the base contract for search response.
// swagger:model EntitySearchResponse
type EntitySearchResponse struct {

	// A dictionary storing an array of <code>Filter</code> object against each facet.
	Facets map[string][]Filter `json:"facets,omitempty"`

	// Numeric code indicating any additional information: 0 - Ok, 1 - Account is being reindexed, 2 - Account indexing has not started, 3 - Invalid Request, 4 - Prefix wildcard query not supported, 5 - MultiWords with code facet not supported, 6 - Account is being onboarded, 7 - Account is being onboarded or reindexed, 8 - Top value trimmed to maxresult allowed 9 - Branches are being indexed, 10 - Faceting not enabled, 11 - Work items not accessible, 19 - Phrase queries with code type filters not supported, 20 - Wildcard queries with code type filters not supported. Any other info code is used for internal purpose.
	InfoCode int32 `json:"infoCode,omitempty"`
}

// Validate validates this entity search response
func (m *EntitySearchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntitySearchResponse) validateFacets(formats strfmt.Registry) error {

	if swag.IsZero(m.Facets) { // not required
		return nil
	}

	for k := range m.Facets {

		if err := validate.Required("facets"+"."+k, "body", m.Facets[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.Facets[k]); i++ {

			if err := m.Facets[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("facets" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntitySearchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntitySearchResponse) UnmarshalBinary(b []byte) error {
	var res EntitySearchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
