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

// WorkItemQueryResult The result of a work item query.
// swagger:model WorkItemQueryResult
type WorkItemQueryResult struct {

	// The date the query was run in the context of.
	// Format: date-time
	AsOf strfmt.DateTime `json:"asOf,omitempty"`

	// The columns of the query.
	Columns []*WorkItemFieldReference `json:"columns"`

	// The result type
	// Enum: [workItem workItemLink]
	QueryResultType interface{} `json:"queryResultType,omitempty"`

	// The type of the query
	// Enum: [flat tree oneHop]
	QueryType interface{} `json:"queryType,omitempty"`

	// The sort columns of the query.
	SortColumns []*WorkItemQuerySortColumn `json:"sortColumns"`

	// The work item links returned by the query.
	WorkItemRelations []*WorkItemLink `json:"workItemRelations"`

	// The work items returned by the query.
	WorkItems []*WorkItemReference `json:"workItems"`
}

// Validate validates this work item query result
func (m *WorkItemQueryResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkItemRelations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkItemQueryResult) validateAsOf(formats strfmt.Registry) error {

	if swag.IsZero(m.AsOf) { // not required
		return nil
	}

	if err := validate.FormatOf("asOf", "body", "date-time", m.AsOf.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkItemQueryResult) validateColumns(formats strfmt.Registry) error {

	if swag.IsZero(m.Columns) { // not required
		return nil
	}

	for i := 0; i < len(m.Columns); i++ {
		if swag.IsZero(m.Columns[i]) { // not required
			continue
		}

		if m.Columns[i] != nil {
			if err := m.Columns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkItemQueryResult) validateSortColumns(formats strfmt.Registry) error {

	if swag.IsZero(m.SortColumns) { // not required
		return nil
	}

	for i := 0; i < len(m.SortColumns); i++ {
		if swag.IsZero(m.SortColumns[i]) { // not required
			continue
		}

		if m.SortColumns[i] != nil {
			if err := m.SortColumns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sortColumns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkItemQueryResult) validateWorkItemRelations(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkItemRelations) { // not required
		return nil
	}

	for i := 0; i < len(m.WorkItemRelations); i++ {
		if swag.IsZero(m.WorkItemRelations[i]) { // not required
			continue
		}

		if m.WorkItemRelations[i] != nil {
			if err := m.WorkItemRelations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workItemRelations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkItemQueryResult) validateWorkItems(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkItems) { // not required
		return nil
	}

	for i := 0; i < len(m.WorkItems); i++ {
		if swag.IsZero(m.WorkItems[i]) { // not required
			continue
		}

		if m.WorkItems[i] != nil {
			if err := m.WorkItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemQueryResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemQueryResult) UnmarshalBinary(b []byte) error {
	var res WorkItemQueryResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
