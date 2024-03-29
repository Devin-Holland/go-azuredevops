// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Create Materialize an existing AAD or MSA user into the VSTS account.

NOTE: Created users are not active in an account unless they have been explicitly assigned a parent group at creation time or have signed in
  and been autolicensed through AAD group memberships.

 Adding a user to an account is required before the user can be added to VSTS groups or assigned an asset.

 The body of the request must be a derived type of GraphUserCreationContext:
  * GraphUserMailAddressCreationContext - Create a new user using the mail address as a reference to an existing user from an external AD or AAD backed provider.
  * GraphUserOriginIdCreationContext - Create a new user using the OriginID as a reference to an existing user from an external AD or AAD backed provider.
  * GraphUserPrincipalNameCreationContext - Create a new user using the principal name as a reference to an existing user from an external AD or AAD backed provider.

 If the user to be added corresponds to a user that was previously deleted, then that user will be restored.

 Optionally, you can add the newly created user as a member of an existing VSTS group and/or specify a custom storage key for the user.
*/
func (a *Client) Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Create",
		Method:             "POST",
		PathPattern:        "/{organization}/_apis/graph/users",
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
Delete Disables a user.

The user will still be visible, but membership checks for the user will return false.”
*/
func (a *Client) Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Delete",
		Method:             "DELETE",
		PathPattern:        "/{organization}/_apis/graph/users/{userDescriptor}",
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
List Get a list of all users in a given scope.

Since the list of users may be large, results are returned in pages of users.  If there are more results
 than can be returned in a single page, the result set will contain a continuation token for retrieval of the
 next set of results.
*/
func (a *Client) List(params *ListParams, authInfo runtime.ClientAuthInfoWriter) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "List",
		Method:             "GET",
		PathPattern:        "/{organization}/_apis/graph/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListOK), nil

}

/*
Update Map an existing user to a different identity
*/
func (a *Client) Update(params *UpdateParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Update",
		Method:             "PATCH",
		PathPattern:        "/{organization}/_apis/graph/users/{userDescriptor}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
