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

// TaskGroup task group
// swagger:model TaskGroup
type TaskGroup struct {
	TaskDefinition

	// Gets or sets comment.
	Comment string `json:"comment,omitempty"`

	// Gets or sets the identity who created.
	CreatedBy *IdentityRef `json:"createdBy,omitempty"`

	// Gets or sets date on which it got created.
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

	// Gets or sets as 'true' to indicate as deleted, 'false' otherwise.
	Deleted bool `json:"deleted,omitempty"`

	// Gets or sets the identity who modified.
	ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`

	// Gets or sets date on which it got modified.
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modifiedOn,omitempty"`

	// Gets or sets the owner.
	Owner string `json:"owner,omitempty"`

	// Gets or sets parent task group Id. This is used while creating a draft task group.
	// Format: uuid
	ParentDefinitionID strfmt.UUID `json:"parentDefinitionId,omitempty"`

	// Gets or sets revision.
	Revision int32 `json:"revision,omitempty"`

	// Gets or sets the tasks.
	Tasks []*TaskGroupStep `json:"tasks"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TaskGroup) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TaskDefinition
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TaskDefinition = aO0

	// now for regular properties
	var propsTaskGroup struct {
		Comment string `json:"comment,omitempty"`

		CreatedBy *IdentityRef `json:"createdBy,omitempty"`

		CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

		Deleted bool `json:"deleted,omitempty"`

		ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`

		ModifiedOn strfmt.DateTime `json:"modifiedOn,omitempty"`

		Owner string `json:"owner,omitempty"`

		ParentDefinitionID strfmt.UUID `json:"parentDefinitionId,omitempty"`

		Revision int32 `json:"revision,omitempty"`

		Tasks []*TaskGroupStep `json:"tasks"`
	}
	if err := swag.ReadJSON(raw, &propsTaskGroup); err != nil {
		return err
	}
	m.Comment = propsTaskGroup.Comment

	m.CreatedBy = propsTaskGroup.CreatedBy

	m.CreatedOn = propsTaskGroup.CreatedOn

	m.Deleted = propsTaskGroup.Deleted

	m.ModifiedBy = propsTaskGroup.ModifiedBy

	m.ModifiedOn = propsTaskGroup.ModifiedOn

	m.Owner = propsTaskGroup.Owner

	m.ParentDefinitionID = propsTaskGroup.ParentDefinitionID

	m.Revision = propsTaskGroup.Revision

	m.Tasks = propsTaskGroup.Tasks

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TaskGroup) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TaskDefinition)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsTaskGroup struct {
		Comment string `json:"comment,omitempty"`

		CreatedBy *IdentityRef `json:"createdBy,omitempty"`

		CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

		Deleted bool `json:"deleted,omitempty"`

		ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`

		ModifiedOn strfmt.DateTime `json:"modifiedOn,omitempty"`

		Owner string `json:"owner,omitempty"`

		ParentDefinitionID strfmt.UUID `json:"parentDefinitionId,omitempty"`

		Revision int32 `json:"revision,omitempty"`

		Tasks []*TaskGroupStep `json:"tasks"`
	}
	propsTaskGroup.Comment = m.Comment

	propsTaskGroup.CreatedBy = m.CreatedBy

	propsTaskGroup.CreatedOn = m.CreatedOn

	propsTaskGroup.Deleted = m.Deleted

	propsTaskGroup.ModifiedBy = m.ModifiedBy

	propsTaskGroup.ModifiedOn = m.ModifiedOn

	propsTaskGroup.Owner = m.Owner

	propsTaskGroup.ParentDefinitionID = m.ParentDefinitionID

	propsTaskGroup.Revision = m.Revision

	propsTaskGroup.Tasks = m.Tasks

	jsonDataPropsTaskGroup, errTaskGroup := swag.WriteJSON(propsTaskGroup)
	if errTaskGroup != nil {
		return nil, errTaskGroup
	}
	_parts = append(_parts, jsonDataPropsTaskGroup)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this task group
func (m *TaskGroup) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TaskDefinition
	if err := m.TaskDefinition.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDefinitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskGroup) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *TaskGroup) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("createdOn", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskGroup) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TaskGroup) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedOn", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskGroup) validateParentDefinitionID(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentDefinitionID) { // not required
		return nil
	}

	if err := validate.FormatOf("parentDefinitionId", "body", "uuid", m.ParentDefinitionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskGroup) validateTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {
		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskGroup) UnmarshalBinary(b []byte) error {
	var res TaskGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
