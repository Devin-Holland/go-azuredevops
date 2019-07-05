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

// WorkItemQueryClause Represents a clause in a work item query. This shows the structure of a work item query.
// swagger:model WorkItemQueryClause
type WorkItemQueryClause struct {

	// Child clauses if the current clause is a logical operator
	Clauses []*WorkItemQueryClause `json:"clauses"`

	// Field associated with condition
	Field *WorkItemFieldReference `json:"field,omitempty"`

	// Right side of the condition when a field to field comparison
	FieldValue *WorkItemFieldReference `json:"fieldValue,omitempty"`

	// Determines if this is a field to field comparison
	IsFieldValue bool `json:"isFieldValue,omitempty"`

	// Logical operator separating the condition clause
	// Enum: [none and or]
	LogicalOperator interface{} `json:"logicalOperator,omitempty"`

	// The field operator
	Operator *WorkItemFieldOperation `json:"operator,omitempty"`

	// Right side of the condition when a field to value comparison
	Value string `json:"value,omitempty"`
}

// Validate validates this work item query clause
func (m *WorkItemQueryClause) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClauses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkItemQueryClause) validateClauses(formats strfmt.Registry) error {

	if swag.IsZero(m.Clauses) { // not required
		return nil
	}

	for i := 0; i < len(m.Clauses); i++ {
		if swag.IsZero(m.Clauses[i]) { // not required
			continue
		}

		if m.Clauses[i] != nil {
			if err := m.Clauses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clauses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkItemQueryClause) validateField(formats strfmt.Registry) error {

	if swag.IsZero(m.Field) { // not required
		return nil
	}

	if m.Field != nil {
		if err := m.Field.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("field")
			}
			return err
		}
	}

	return nil
}

func (m *WorkItemQueryClause) validateFieldValue(formats strfmt.Registry) error {

	if swag.IsZero(m.FieldValue) { // not required
		return nil
	}

	if m.FieldValue != nil {
		if err := m.FieldValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldValue")
			}
			return err
		}
	}

	return nil
}

func (m *WorkItemQueryClause) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	if m.Operator != nil {
		if err := m.Operator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemQueryClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemQueryClause) UnmarshalBinary(b []byte) error {
	var res WorkItemQueryClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}