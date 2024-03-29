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

// PublisherEvent Wrapper around an event which is being published
// swagger:model PublisherEvent
type PublisherEvent struct {

	// Add key/value pairs which will be stored with a published notification in the SH service DB.  This key/value pairs are for diagnostic purposes only and will have not effect on the delivery of a notificaton.
	Diagnostics map[string]string `json:"diagnostics,omitempty"`

	// The event being published
	Event *Event `json:"event,omitempty"`

	// Gets or sets flag for filtered events
	IsFilteredEvent bool `json:"isFilteredEvent,omitempty"`

	// Additional data that needs to be sent as part of notification to complement the Resource data in the Event
	NotificationData map[string]string `json:"notificationData,omitempty"`

	// Gets or sets the array of older supported resource versions.
	OtherResourceVersions []*VersionedResource `json:"otherResourceVersions"`

	// Optional publisher-input filters which restricts the set of subscriptions which are triggered by the event
	PublisherInputFilters []*InputFilter `json:"publisherInputFilters"`

	// Gets or sets matchd hooks subscription which caused this event.
	Subscription *Subscription `json:"subscription,omitempty"`
}

// Validate validates this publisher event
func (m *PublisherEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherResourceVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublisherInputFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublisherEvent) validateEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {
		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

func (m *PublisherEvent) validateOtherResourceVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.OtherResourceVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.OtherResourceVersions); i++ {
		if swag.IsZero(m.OtherResourceVersions[i]) { // not required
			continue
		}

		if m.OtherResourceVersions[i] != nil {
			if err := m.OtherResourceVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("otherResourceVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PublisherEvent) validatePublisherInputFilters(formats strfmt.Registry) error {

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

func (m *PublisherEvent) validateSubscription(formats strfmt.Registry) error {

	if swag.IsZero(m.Subscription) { // not required
		return nil
	}

	if m.Subscription != nil {
		if err := m.Subscription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscription")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublisherEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublisherEvent) UnmarshalBinary(b []byte) error {
	var res PublisherEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
