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

// IdentityBase Base Identity class to allow "trimmed" identity class in the GetConnectionData API Makes sure that on-the-wire representations of the derived classes are compatible with each other (e.g. Server responds with PublicIdentity object while client deserializes it as Identity object) Derived classes should not have additional [DataMember] properties
// swagger:model IdentityBase
type IdentityBase struct {

	// The custom display name for the identity (if any). Setting this property to an empty string will clear the existing custom display name. Setting this property to null will not affect the existing persisted value (since null values do not get sent over the wire or to the database)
	CustomDisplayName string `json:"customDisplayName,omitempty"`

	// descriptor
	Descriptor *IdentityDescriptor `json:"descriptor,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// is active
	IsActive bool `json:"isActive,omitempty"`

	// is container
	IsContainer bool `json:"isContainer,omitempty"`

	// master Id
	// Format: uuid
	MasterID strfmt.UUID `json:"masterId,omitempty"`

	// member ids
	MemberIds []strfmt.UUID `json:"memberIds"`

	// member of
	MemberOf []*IdentityDescriptor `json:"memberOf"`

	// members
	Members []*IdentityDescriptor `json:"members"`

	// meta type Id
	MetaTypeID int32 `json:"metaTypeId,omitempty"`

	// properties
	Properties *PropertiesCollection `json:"properties,omitempty"`

	// The display name for the identity as specified by the source identity provider.
	ProviderDisplayName string `json:"providerDisplayName,omitempty"`

	// resource version
	ResourceVersion int32 `json:"resourceVersion,omitempty"`

	// social descriptor
	SocialDescriptor string `json:"socialDescriptor,omitempty"`

	// subject descriptor
	SubjectDescriptor string `json:"subjectDescriptor,omitempty"`

	// unique user Id
	UniqueUserID int32 `json:"uniqueUserId,omitempty"`
}

// Validate validates this identity base
func (m *IdentityBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityBase) validateDescriptor(formats strfmt.Registry) error {

	if swag.IsZero(m.Descriptor) { // not required
		return nil
	}

	if m.Descriptor != nil {
		if err := m.Descriptor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("descriptor")
			}
			return err
		}
	}

	return nil
}

func (m *IdentityBase) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IdentityBase) validateMasterID(formats strfmt.Registry) error {

	if swag.IsZero(m.MasterID) { // not required
		return nil
	}

	if err := validate.FormatOf("masterId", "body", "uuid", m.MasterID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IdentityBase) validateMemberIds(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberIds) { // not required
		return nil
	}

	for i := 0; i < len(m.MemberIds); i++ {

		if err := validate.FormatOf("memberIds"+"."+strconv.Itoa(i), "body", "uuid", m.MemberIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *IdentityBase) validateMemberOf(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberOf) { // not required
		return nil
	}

	for i := 0; i < len(m.MemberOf); i++ {
		if swag.IsZero(m.MemberOf[i]) { // not required
			continue
		}

		if m.MemberOf[i] != nil {
			if err := m.MemberOf[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memberOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IdentityBase) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IdentityBase) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentityBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityBase) UnmarshalBinary(b []byte) error {
	var res IdentityBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
