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

// AccountMyWorkResult account my work result
// swagger:model AccountMyWorkResult
type AccountMyWorkResult struct {

	// True, when length of WorkItemDetails is same as the limit
	QuerySizeLimitExceeded bool `json:"querySizeLimitExceeded,omitempty"`

	// WorkItem Details
	WorkItemDetails []*AccountWorkWorkItemModel `json:"workItemDetails"`
}

// Validate validates this account my work result
func (m *AccountMyWorkResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkItemDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountMyWorkResult) validateWorkItemDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkItemDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.WorkItemDetails); i++ {
		if swag.IsZero(m.WorkItemDetails[i]) { // not required
			continue
		}

		if m.WorkItemDetails[i] != nil {
			if err := m.WorkItemDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workItemDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountMyWorkResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountMyWorkResult) UnmarshalBinary(b []byte) error {
	var res AccountMyWorkResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
