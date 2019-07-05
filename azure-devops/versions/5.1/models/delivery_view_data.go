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

// DeliveryViewData Data contract for Data of Delivery View
// swagger:model DeliveryViewData
type DeliveryViewData struct {
	PlanViewData

	// Work item child id to parenet id map
	ChildIDToParentIDMap map[string]int32 `json:"childIdToParentIdMap,omitempty"`

	// Filter criteria status of the timeline
	CriteriaStatus *TimelineCriteriaStatus `json:"criteriaStatus,omitempty"`

	// The end date of the delivery view data
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// Max number of teams can be configured for a delivery plan.
	MaxExpandedTeams int32 `json:"maxExpandedTeams,omitempty"`

	// The start date for the delivery view data
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// All the team data
	Teams []*TimelineTeamData `json:"teams"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DeliveryViewData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PlanViewData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PlanViewData = aO0

	// now for regular properties
	var propsDeliveryViewData struct {
		ChildIDToParentIDMap map[string]int32 `json:"childIdToParentIdMap,omitempty"`

		CriteriaStatus *TimelineCriteriaStatus `json:"criteriaStatus,omitempty"`

		EndDate strfmt.DateTime `json:"endDate,omitempty"`

		MaxExpandedTeams int32 `json:"maxExpandedTeams,omitempty"`

		StartDate strfmt.DateTime `json:"startDate,omitempty"`

		Teams []*TimelineTeamData `json:"teams"`
	}
	if err := swag.ReadJSON(raw, &propsDeliveryViewData); err != nil {
		return err
	}
	m.ChildIDToParentIDMap = propsDeliveryViewData.ChildIDToParentIDMap

	m.CriteriaStatus = propsDeliveryViewData.CriteriaStatus

	m.EndDate = propsDeliveryViewData.EndDate

	m.MaxExpandedTeams = propsDeliveryViewData.MaxExpandedTeams

	m.StartDate = propsDeliveryViewData.StartDate

	m.Teams = propsDeliveryViewData.Teams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DeliveryViewData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.PlanViewData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsDeliveryViewData struct {
		ChildIDToParentIDMap map[string]int32 `json:"childIdToParentIdMap,omitempty"`

		CriteriaStatus *TimelineCriteriaStatus `json:"criteriaStatus,omitempty"`

		EndDate strfmt.DateTime `json:"endDate,omitempty"`

		MaxExpandedTeams int32 `json:"maxExpandedTeams,omitempty"`

		StartDate strfmt.DateTime `json:"startDate,omitempty"`

		Teams []*TimelineTeamData `json:"teams"`
	}
	propsDeliveryViewData.ChildIDToParentIDMap = m.ChildIDToParentIDMap

	propsDeliveryViewData.CriteriaStatus = m.CriteriaStatus

	propsDeliveryViewData.EndDate = m.EndDate

	propsDeliveryViewData.MaxExpandedTeams = m.MaxExpandedTeams

	propsDeliveryViewData.StartDate = m.StartDate

	propsDeliveryViewData.Teams = m.Teams

	jsonDataPropsDeliveryViewData, errDeliveryViewData := swag.WriteJSON(propsDeliveryViewData)
	if errDeliveryViewData != nil {
		return nil, errDeliveryViewData
	}
	_parts = append(_parts, jsonDataPropsDeliveryViewData)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delivery view data
func (m *DeliveryViewData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PlanViewData
	if err := m.PlanViewData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCriteriaStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeliveryViewData) validateCriteriaStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CriteriaStatus) { // not required
		return nil
	}

	if m.CriteriaStatus != nil {
		if err := m.CriteriaStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("criteriaStatus")
			}
			return err
		}
	}

	return nil
}

func (m *DeliveryViewData) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryViewData) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryViewData) validateTeams(formats strfmt.Registry) error {

	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeliveryViewData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeliveryViewData) UnmarshalBinary(b []byte) error {
	var res DeliveryViewData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
