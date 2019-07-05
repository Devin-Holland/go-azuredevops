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

// TeamIterationAttributes team iteration attributes
// swagger:model TeamIterationAttributes
type TeamIterationAttributes struct {

	// Finish date of the iteration. Date-only, correct unadjusted at midnight in UTC.
	// Format: date-time
	FinishDate strfmt.DateTime `json:"finishDate,omitempty"`

	// Start date of the iteration. Date-only, correct unadjusted at midnight in UTC.
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Time frame of the iteration, such as past, current or future.
	// Enum: [past current future]
	TimeFrame interface{} `json:"timeFrame,omitempty"`
}

// Validate validates this team iteration attributes
func (m *TeamIterationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamIterationAttributes) validateFinishDate(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishDate) { // not required
		return nil
	}

	if err := validate.FormatOf("finishDate", "body", "date-time", m.FinishDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TeamIterationAttributes) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamIterationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamIterationAttributes) UnmarshalBinary(b []byte) error {
	var res TeamIterationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
