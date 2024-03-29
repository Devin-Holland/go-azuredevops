// Code generated by go-swagger; DO NOT EDIT.

package test_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new test history API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for test history API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Query Get history of a test method using TestHistoryQuery
*/
func (a *Client) Query(params *QueryParams, authInfo runtime.ClientAuthInfoWriter) (*QueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Query",
		Method:             "POST",
		PathPattern:        "/{organization}/{project}/_apis/test/Results/testhistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
