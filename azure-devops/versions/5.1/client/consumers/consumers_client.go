// Code generated by go-swagger; DO NOT EDIT.

package consumers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new consumers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for consumers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetConsumerAction Get details about a specific consumer action.
*/
func (a *Client) GetConsumerAction(params *GetConsumerActionParams, authInfo runtime.ClientAuthInfoWriter) (*GetConsumerActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConsumerActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Consumer Action",
		Method:             "GET",
		PathPattern:        "/{organization}/_apis/hooks/consumers/{consumerId}/actions/{consumerActionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConsumerActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConsumerActionOK), nil

}

/*
ListConsumerActions Get a list of consumer actions for a specific consumer.
*/
func (a *Client) ListConsumerActions(params *ListConsumerActionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListConsumerActionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConsumerActionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "List Consumer Actions",
		Method:             "GET",
		PathPattern:        "/{organization}/_apis/hooks/consumers/{consumerId}/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListConsumerActionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListConsumerActionsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}