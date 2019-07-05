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

// Notification Defines the data contract of the result of processing an event for a subscription.
// swagger:model Notification
type Notification struct {

	// Gets or sets date and time that this result was created.
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// Details about this notification (if available)
	Details *NotificationDetails `json:"details,omitempty"`

	// The event id associated with this notification
	// Format: uuid
	EventID strfmt.UUID `json:"eventId,omitempty"`

	// The notification id
	ID int32 `json:"id,omitempty"`

	// Gets or sets date and time that this result was last modified.
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// Result of the notification
	// Enum: [pending succeeded failed filtered]
	Result interface{} `json:"result,omitempty"`

	// Status of the notification
	// Enum: [queued processing requestInProgress completed]
	Status interface{} `json:"status,omitempty"`

	// The subscriber Id  associated with this notification. This is the last identity who touched in the subscription. In case of test notifications it can be the tester if the subscription is not created yet.
	// Format: uuid
	SubscriberID strfmt.UUID `json:"subscriberId,omitempty"`

	// The subscription id associated with this notification
	// Format: uuid
	SubscriptionID strfmt.UUID `json:"subscriptionId,omitempty"`
}

// Validate validates this notification
func (m *Notification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notification) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *Notification) validateEventID(formats strfmt.Registry) error {

	if swag.IsZero(m.EventID) { // not required
		return nil
	}

	if err := validate.FormatOf("eventId", "body", "uuid", m.EventID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateSubscriberID(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriberID) { // not required
		return nil
	}

	if err := validate.FormatOf("subscriberId", "body", "uuid", m.SubscriberID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateSubscriptionID(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionID) { // not required
		return nil
	}

	if err := validate.FormatOf("subscriptionId", "body", "uuid", m.SubscriptionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notification) UnmarshalBinary(b []byte) error {
	var res Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}