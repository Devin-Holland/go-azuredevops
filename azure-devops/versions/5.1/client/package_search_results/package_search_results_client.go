// Code generated by go-swagger; DO NOT EDIT.

package package_search_results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new package search results API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for package search results API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
FetchPackageSearchResults Provides a set of results for the search text.
*/
func (a *Client) FetchPackageSearchResults(params *FetchPackageSearchResultsParams, authInfo runtime.ClientAuthInfoWriter) (*FetchPackageSearchResultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchPackageSearchResultsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Fetch Package Search Results",
		Method:             "POST",
		PathPattern:        "/{organization}/_apis/search/packagesearchresults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchPackageSearchResultsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FetchPackageSearchResultsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
