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

// SubscriptionsQuery Defines a query for service hook subscriptions.
// swagger:model SubscriptionsQuery
type SubscriptionsQuery struct {

	// Optional consumer action id to restrict the results to (null for any)
	ConsumerActionID string `json:"consumerActionId,omitempty"`

	// Optional consumer id to restrict the results to (null for any)
	ConsumerID string `json:"consumerId,omitempty"`

	// Filter for subscription consumer inputs
	ConsumerInputFilters []*InputFilter `json:"consumerInputFilters"`

	// Optional event type id to restrict the results to (null for any)
	EventType string `json:"eventType,omitempty"`

	// Optional publisher id to restrict the results to (null for any)
	PublisherID string `json:"publisherId,omitempty"`

	// Filter for subscription publisher inputs
	PublisherInputFilters []*InputFilter `json:"publisherInputFilters"`

	// Results from the query
	Results []*Subscription `json:"results"`

	// Optional subscriber filter.
	// Format: uuid
	SubscriberID strfmt.UUID `json:"subscriberId,omitempty"`
}

// Validate validates this subscriptions query
func (m *SubscriptionsQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsumerInputFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublisherInputFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionsQuery) validateConsumerInputFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.ConsumerInputFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsumerInputFilters); i++ {
		if swag.IsZero(m.ConsumerInputFilters[i]) { // not required
			continue
		}

		if m.ConsumerInputFilters[i] != nil {
			if err := m.ConsumerInputFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consumerInputFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubscriptionsQuery) validatePublisherInputFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.PublisherInputFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.PublisherInputFilters); i++ {
		if swag.IsZero(m.PublisherInputFilters[i]) { // not required
			continue
		}

		if m.PublisherInputFilters[i] != nil {
			if err := m.PublisherInputFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("publisherInputFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubscriptionsQuery) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubscriptionsQuery) validateSubscriberID(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriberID) { // not required
		return nil
	}

	if err := validate.FormatOf("subscriberId", "body", "uuid", m.SubscriberID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionsQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionsQuery) UnmarshalBinary(b []byte) error {
	var res SubscriptionsQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}