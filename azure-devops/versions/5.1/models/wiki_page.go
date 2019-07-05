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

// WikiPage Defines a page in a wiki.
// swagger:model WikiPage
type WikiPage struct {
	WikiPageCreateOrUpdateParameters

	// Path of the git item corresponding to the wiki page stored in the backing Git repository.
	GitItemPath string `json:"gitItemPath,omitempty"`

	// When present, permanent identifier for the wiki page
	ID int32 `json:"id,omitempty"`

	// True if a page is non-conforming, i.e. 1) if the name doesn't match page naming standards. 2) if the page does not have a valid entry in the appropriate order file.
	IsNonConformant bool `json:"isNonConformant,omitempty"`

	// True if this page has subpages under its path.
	IsParentPage bool `json:"isParentPage,omitempty"`

	// Order of the wiki page, relative to other pages in the same hierarchy level.
	Order int32 `json:"order,omitempty"`

	// Path of the wiki page.
	Path string `json:"path,omitempty"`

	// Remote web url to the wiki page.
	RemoteURL string `json:"remoteUrl,omitempty"`

	// List of subpages of the current page.
	SubPages []*WikiPage `json:"subPages"`

	// REST url for this wiki page.
	URL string `json:"url,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WikiPage) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WikiPageCreateOrUpdateParameters
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WikiPageCreateOrUpdateParameters = aO0

	// now for regular properties
	var propsWikiPage struct {
		GitItemPath string `json:"gitItemPath,omitempty"`

		ID int32 `json:"id,omitempty"`

		IsNonConformant bool `json:"isNonConformant,omitempty"`

		IsParentPage bool `json:"isParentPage,omitempty"`

		Order int32 `json:"order,omitempty"`

		Path string `json:"path,omitempty"`

		RemoteURL string `json:"remoteUrl,omitempty"`

		SubPages []*WikiPage `json:"subPages"`

		URL string `json:"url,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsWikiPage); err != nil {
		return err
	}
	m.GitItemPath = propsWikiPage.GitItemPath

	m.ID = propsWikiPage.ID

	m.IsNonConformant = propsWikiPage.IsNonConformant

	m.IsParentPage = propsWikiPage.IsParentPage

	m.Order = propsWikiPage.Order

	m.Path = propsWikiPage.Path

	m.RemoteURL = propsWikiPage.RemoteURL

	m.SubPages = propsWikiPage.SubPages

	m.URL = propsWikiPage.URL

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WikiPage) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.WikiPageCreateOrUpdateParameters)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsWikiPage struct {
		GitItemPath string `json:"gitItemPath,omitempty"`

		ID int32 `json:"id,omitempty"`

		IsNonConformant bool `json:"isNonConformant,omitempty"`

		IsParentPage bool `json:"isParentPage,omitempty"`

		Order int32 `json:"order,omitempty"`

		Path string `json:"path,omitempty"`

		RemoteURL string `json:"remoteUrl,omitempty"`

		SubPages []*WikiPage `json:"subPages"`

		URL string `json:"url,omitempty"`
	}
	propsWikiPage.GitItemPath = m.GitItemPath

	propsWikiPage.ID = m.ID

	propsWikiPage.IsNonConformant = m.IsNonConformant

	propsWikiPage.IsParentPage = m.IsParentPage

	propsWikiPage.Order = m.Order

	propsWikiPage.Path = m.Path

	propsWikiPage.RemoteURL = m.RemoteURL

	propsWikiPage.SubPages = m.SubPages

	propsWikiPage.URL = m.URL

	jsonDataPropsWikiPage, errWikiPage := swag.WriteJSON(propsWikiPage)
	if errWikiPage != nil {
		return nil, errWikiPage
	}
	_parts = append(_parts, jsonDataPropsWikiPage)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this wiki page
func (m *WikiPage) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WikiPageCreateOrUpdateParameters
	if err := m.WikiPageCreateOrUpdateParameters.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubPages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WikiPage) validateSubPages(formats strfmt.Registry) error {

	if swag.IsZero(m.SubPages) { // not required
		return nil
	}

	for i := 0; i < len(m.SubPages); i++ {
		if swag.IsZero(m.SubPages[i]) { // not required
			continue
		}

		if m.SubPages[i] != nil {
			if err := m.SubPages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subPages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WikiPage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WikiPage) UnmarshalBinary(b []byte) error {
	var res WikiPage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}