// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new endpoints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Create Create a service endpoint.
*/
func (a *Client) Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Create",
		Method:             "POST",
		PathPattern:        "/{organization}/{project}/_apis/serviceendpoint/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOK), nil

}

/*
Delete Delete a service endpoint.
*/
func (a *Client) Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Delete",
		Method:             "DELETE",
		PathPattern:        "/{organization}/{project}/_apis/serviceendpoint/endpoints/{endpointId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOK), nil

}

/*
Get Get the service endpoint details.
*/
func (a *Client) Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/serviceendpoint/endpoints/{endpointId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOK), nil

}

/*
GetServiceEndpointsByNames Get the service endpoints by name.
*/
func (a *Client) GetServiceEndpointsByNames(params *GetServiceEndpointsByNamesParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceEndpointsByNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceEndpointsByNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Service Endpoints By Names",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/serviceendpoint/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceEndpointsByNamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceEndpointsByNamesOK), nil

}

/*
UpdateServiceEndpoint Update a service endpoint.
*/
func (a *Client) UpdateServiceEndpoint(params *UpdateServiceEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Update Service Endpoint",
		Method:             "PUT",
		PathPattern:        "/{organization}/{project}/_apis/serviceendpoint/endpoints/{endpointId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServiceEndpointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateServiceEndpointOK), nil

}

/*
UpdateServiceEndpoints Update the service endpoints.
*/
func (a *Client) UpdateServiceEndpoints(params *UpdateServiceEndpointsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceEndpointsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Update Service Endpoints",
		Method:             "PUT",
		PathPattern:        "/{organization}/{project}/_apis/serviceendpoint/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServiceEndpointsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateServiceEndpointsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}