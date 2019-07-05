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

// SourceRepositories A set of repositories returned from the source provider.
// swagger:model SourceRepositories
type SourceRepositories struct {

	// A token used to continue this paged request; 'null' if the request is complete
	ContinuationToken string `json:"continuationToken,omitempty"`

	// The number of repositories requested for each page
	PageLength int32 `json:"pageLength,omitempty"`

	// A list of repositories
	Repositories []*SourceRepository `json:"repositories"`

	// The total number of pages, or '-1' if unknown
	TotalPageCount int32 `json:"totalPageCount,omitempty"`
}

// Validate validates this source repositories
func (m *SourceRepositories) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepositories(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceRepositories) validateRepositories(formats strfmt.Registry) error {

	if swag.IsZero(m.Repositories) { // not required
		return nil
	}

	for i := 0; i < len(m.Repositories); i++ {
		if swag.IsZero(m.Repositories[i]) { // not required
			continue
		}

		if m.Repositories[i] != nil {
			if err := m.Repositories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourceRepositories) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceRepositories) UnmarshalBinary(b []byte) error {
	var res SourceRepositories
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}