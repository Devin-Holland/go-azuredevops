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

// WikiResult Defines the wiki result that matched a wiki search request.
// swagger:model WikiResult
type WikiResult struct {

	// Collection of the result file.
	Collection *Collection `json:"collection,omitempty"`

	// ContentId of the result file.
	ContentID string `json:"contentId,omitempty"`

	// Name of the result file.
	FileName string `json:"fileName,omitempty"`

	// Highlighted snippets of fields that match the search request. The list is sorted by relevance of the snippets.
	Hits []*WikiHit `json:"hits"`

	// Path at which result file is present.
	Path string `json:"path,omitempty"`

	// Project details of the wiki document.
	Project *ProjectReference `json:"project,omitempty"`

	// Wiki information for the result.
	Wiki *Wiki `json:"wiki,omitempty"`
}

// Validate validates this wiki result
func (m *WikiResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWiki(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WikiResult) validateCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.Collection) { // not required
		return nil
	}

	if m.Collection != nil {
		if err := m.Collection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collection")
			}
			return err
		}
	}

	return nil
}

func (m *WikiResult) validateHits(formats strfmt.Registry) error {

	if swag.IsZero(m.Hits) { // not required
		return nil
	}

	for i := 0; i < len(m.Hits); i++ {
		if swag.IsZero(m.Hits[i]) { // not required
			continue
		}

		if m.Hits[i] != nil {
			if err := m.Hits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WikiResult) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *WikiResult) validateWiki(formats strfmt.Registry) error {

	if swag.IsZero(m.Wiki) { // not required
		return nil
	}

	if m.Wiki != nil {
		if err := m.Wiki.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wiki")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WikiResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WikiResult) UnmarshalBinary(b []byte) error {
	var res WikiResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
