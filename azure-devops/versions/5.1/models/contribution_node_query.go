// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContributionNodeQuery A query that can be issued for contribution nodes
// swagger:model ContributionNodeQuery
type ContributionNodeQuery struct {

	// The contribution ids of the nodes to find.
	ContributionIds []string `json:"contributionIds"`

	// Contextual information that can be leveraged by contribution constraints
	DataProviderContext *DataProviderContext `json:"dataProviderContext,omitempty"`

	// Indicator if contribution provider details should be included in the result.
	IncludeProviderDetails bool `json:"includeProviderDetails,omitempty"`

	// Query options tpo be used when fetching ContributionNodes
	// Enum: [none includeSelf includeChildren includeSubTree includeAll ignoreConstraints]
	QueryOptions interface{} `json:"queryOptions,omitempty"`
}

// Validate validates this contribution node query
func (m *ContributionNodeQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataProviderContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContributionNodeQuery) validateDataProviderContext(formats strfmt.Registry) error {

	if swag.IsZero(m.DataProviderContext) { // not required
		return nil
	}

	if m.DataProviderContext != nil {
		if err := m.DataProviderContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataProviderContext")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContributionNodeQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContributionNodeQuery) UnmarshalBinary(b []byte) error {
	var res ContributionNodeQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
