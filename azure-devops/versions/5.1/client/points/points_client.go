// Code generated by go-swagger; DO NOT EDIT.

package points

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new points API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for points API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPoint Get a test point.
*/
func (a *Client) GetPoint(params *GetPointParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Point",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/test/Plans/{planId}/Suites/{suiteId}/points/{pointIds}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPointOK), nil

}

/*
GetPointsByQuery Get test points using query.
*/
func (a *Client) GetPointsByQuery(params *GetPointsByQueryParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointsByQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPointsByQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Points By Query",
		Method:             "POST",
		PathPattern:        "/{organization}/{project}/_apis/test/points",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPointsByQueryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPointsByQueryOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
