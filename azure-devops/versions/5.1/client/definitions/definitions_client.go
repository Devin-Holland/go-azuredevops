// Code generated by go-swagger; DO NOT EDIT.

package definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new definitions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for definitions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDefinitionRevision Get release definition for a given definitionId and revision
*/
func (a *Client) GetDefinitionRevision(params *GetDefinitionRevisionParams, authInfo runtime.ClientAuthInfoWriter) (*GetDefinitionRevisionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefinitionRevisionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Definition Revision",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/Release/definitions/{definitionId}/revisions/{revision}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDefinitionRevisionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDefinitionRevisionOK), nil

}

/*
GetReleaseDefinitionHistory Get revision history for a release definition
*/
func (a *Client) GetReleaseDefinitionHistory(params *GetReleaseDefinitionHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetReleaseDefinitionHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReleaseDefinitionHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Release Definition History",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/Release/definitions/{definitionId}/revisions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReleaseDefinitionHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReleaseDefinitionHistoryOK), nil

}

/*
PutOrganizationProjectApisReleaseDefinitions Update a release definition.
*/
func (a *Client) PutOrganizationProjectApisReleaseDefinitions(params *PutOrganizationProjectApisReleaseDefinitionsParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrganizationProjectApisReleaseDefinitionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrganizationProjectApisReleaseDefinitionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrganizationProjectApisReleaseDefinitions",
		Method:             "PUT",
		PathPattern:        "/{organization}/{project}/_apis/release/definitions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrganizationProjectApisReleaseDefinitionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrganizationProjectApisReleaseDefinitionsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}