// Code generated by go-swagger; DO NOT EDIT.

package properties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new properties API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for properties API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetBuildProperties Gets properties for a build.
*/
func (a *Client) GetBuildProperties(params *GetBuildPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Build Properties",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/build/builds/{buildId}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBuildPropertiesOK), nil

}

/*
GetDefinitionProperties Gets properties for a definition.
*/
func (a *Client) GetDefinitionProperties(params *GetDefinitionPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDefinitionPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefinitionPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Definition Properties",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/build/definitions/{definitionId}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDefinitionPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDefinitionPropertiesOK), nil

}

/*
UpdateBuildProperties Updates properties for a build.
*/
func (a *Client) UpdateBuildProperties(params *UpdateBuildPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBuildPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBuildPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Update Build Properties",
		Method:             "PATCH",
		PathPattern:        "/{organization}/{project}/_apis/build/builds/{buildId}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json-patch+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBuildPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateBuildPropertiesOK), nil

}

/*
UpdateDefinitionProperties Updates properties for a definition.
*/
func (a *Client) UpdateDefinitionProperties(params *UpdateDefinitionPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDefinitionPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDefinitionPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Update Definition Properties",
		Method:             "PATCH",
		PathPattern:        "/{organization}/{project}/_apis/build/definitions/{definitionId}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json-patch+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDefinitionPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDefinitionPropertiesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
