// Code generated by go-swagger; DO NOT EDIT.

package test_suite_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new test suite entry API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for test suite entry API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ReorderSuiteEntries Reorder test suite entries in the test suite.
*/
func (a *Client) ReorderSuiteEntries(params *ReorderSuiteEntriesParams, authInfo runtime.ClientAuthInfoWriter) (*ReorderSuiteEntriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReorderSuiteEntriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Reorder Suite Entries",
		Method:             "PATCH",
		PathPattern:        "/{organization}/{project}/_apis/testplan/suiteentry/{suiteId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReorderSuiteEntriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReorderSuiteEntriesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
