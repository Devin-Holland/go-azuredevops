// Code generated by go-swagger; DO NOT EDIT.

package wiql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new wiql API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for wiql API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
HeadOrganizationProjectTeamApisWitWiqlID Gets the results of the query given the query ID.
*/
func (a *Client) HeadOrganizationProjectTeamApisWitWiqlID(params *HeadOrganizationProjectTeamApisWitWiqlIDParams, authInfo runtime.ClientAuthInfoWriter) (*HeadOrganizationProjectTeamApisWitWiqlIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadOrganizationProjectTeamApisWitWiqlIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "HeadOrganizationProjectTeamApisWitWiqlID",
		Method:             "HEAD",
		PathPattern:        "/{organization}/{project}/{team}/_apis/wit/wiql/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HeadOrganizationProjectTeamApisWitWiqlIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HeadOrganizationProjectTeamApisWitWiqlIDOK), nil

}

/*
QueryByID Gets the results of the query given the query ID.
*/
func (a *Client) QueryByID(params *QueryByIDParams, authInfo runtime.ClientAuthInfoWriter) (*QueryByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Query By Id",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/{team}/_apis/wit/wiql/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryByIDOK), nil

}

/*
QueryByWiql Gets the results of the query given its WIQL.
*/
func (a *Client) QueryByWiql(params *QueryByWiqlParams, authInfo runtime.ClientAuthInfoWriter) (*QueryByWiqlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryByWiqlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Query By Wiql",
		Method:             "POST",
		PathPattern:        "/{organization}/{project}/{team}/_apis/wit/wiql",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryByWiqlReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryByWiqlOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}