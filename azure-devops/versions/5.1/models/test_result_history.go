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

// TestResultHistory test result history
// swagger:model TestResultHistory
type TestResultHistory struct {

	// group by field
	GroupByField string `json:"groupByField,omitempty"`

	// results for group
	ResultsForGroup []*TestResultHistoryDetailsForGroup `json:"resultsForGroup"`
}

// Validate validates this test result history
func (m *TestResultHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResultsForGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestResultHistory) validateResultsForGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultsForGroup) { // not required
		return nil
	}

	for i := 0; i < len(m.ResultsForGroup); i++ {
		if swag.IsZero(m.ResultsForGroup[i]) { // not required
			continue
		}

		if m.ResultsForGroup[i] != nil {
			if err := m.ResultsForGroup[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resultsForGroup" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestResultHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultHistory) UnmarshalBinary(b []byte) error {
	var res TestResultHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}