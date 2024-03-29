// Code generated by go-swagger; DO NOT EDIT.

package comments_reactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new comments reactions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for comments reactions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PutOrganizationProjectApisWitWorkItemsWorkItemIDCommentsCommentIDReactionsReactionType Adds a new reaction to a comment.
*/
func (a *Client) PutOrganizationProjectApisWitWorkItemsWorkItemIDCommentsCommentIDReactionsReactionType(params *PutOrganizationProjectApisWitWorkItemsWorkItemIDCommentsCommentIDReactionsReactionTypeParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrganizationProjectApisWitWorkItemsWorkItemIDCommentsCommentIDReactionsReactionTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrganizationProjectApisWitWorkItemsWorkItemIDCommentsCommentIDReactionsReactionTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrganizationProjectApisWitWorkItemsWorkItemIDCommentsCommentIDReactionsReactionType",
		Method:             "PUT",
		PathPattern:        "/{organization}/{project}/_apis/wit/workItems/{workItemId}/comments/{commentId}/reactions/{reactionType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrganizationProjectApisWitWorkItemsWorkItemIDCommentsCommentIDReactionsReactionTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrganizationProjectApisWitWorkItemsWorkItemIDCommentsCommentIDReactionsReactionTypeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
