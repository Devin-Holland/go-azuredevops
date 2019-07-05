// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TaskDefinitionEndpoint task definition endpoint
// swagger:model TaskDefinitionEndpoint
type TaskDefinitionEndpoint struct {

	// An ID that identifies a service connection to be used for authenticating endpoint requests.
	ConnectionID string `json:"connectionId,omitempty"`

	// An Json based keyselector to filter response returned by fetching the endpoint <c>Url</c>.A Json based keyselector must be prefixed with "jsonpath:". KeySelector can be used to specify the filter to get the keys for the values specified with Selector. <example> The following keyselector defines an Json for extracting nodes named 'ServiceName'. <code> endpoint.KeySelector = "jsonpath://ServiceName"; </code></example>
	KeySelector string `json:"keySelector,omitempty"`

	// The scope as understood by Connected Services. Essentialy, a project-id for now.
	Scope string `json:"scope,omitempty"`

	// An XPath/Json based selector to filter response returned by fetching the endpoint <c>Url</c>. An XPath based selector must be prefixed with the string "xpath:". A Json based selector must be prefixed with "jsonpath:". <example> The following selector defines an XPath for extracting nodes named 'ServiceName'. <code> endpoint.Selector = "xpath://ServiceName"; </code></example>
	Selector string `json:"selector,omitempty"`

	// TaskId that this endpoint belongs to.
	TaskID string `json:"taskId,omitempty"`

	// URL to GET.
	URL string `json:"url,omitempty"`
}

// Validate validates this task definition endpoint
func (m *TaskDefinitionEndpoint) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskDefinitionEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskDefinitionEndpoint) UnmarshalBinary(b []byte) error {
	var res TaskDefinitionEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
