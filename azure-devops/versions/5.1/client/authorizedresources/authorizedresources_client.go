// Code generated by go-swagger; DO NOT EDIT.

package authorizedresources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new authorizedresources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authorizedresources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AuthorizeProjectResources authorize project resources API
*/
func (a *Client) AuthorizeProjectResources(params *AuthorizeProjectResourcesParams, authInfo runtime.ClientAuthInfoWriter) (*AuthorizeProjectResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthorizeProjectResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Authorize Project Resources",
		Method:             "PATCH",
		PathPattern:        "/{organization}/{project}/_apis/build/authorizedresources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AuthorizeProjectResourcesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthorizeProjectResourcesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
