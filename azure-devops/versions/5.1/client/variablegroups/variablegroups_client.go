// Code generated by go-swagger; DO NOT EDIT.

package variablegroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new variablegroups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for variablegroups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Add Add a variable group.
*/
func (a *Client) Add(params *AddParams, authInfo runtime.ClientAuthInfoWriter) (*AddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Add",
		Method:             "POST",
		PathPattern:        "/{organization}/{project}/_apis/distributedtask/variablegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddOK), nil

}

/*
Delete Delete a variable group
*/
func (a *Client) Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Delete",
		Method:             "DELETE",
		PathPattern:        "/{organization}/{project}/_apis/distributedtask/variablegroups/{groupId}",
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
Get Get a variable group.
*/
func (a *Client) Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/distributedtask/variablegroups/{groupId}",
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
GetVariableGroupsByID Get variable groups by ids.
*/
func (a *Client) GetVariableGroupsByID(params *GetVariableGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariableGroupsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariableGroupsByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Variable Groups By Id",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/distributedtask/variablegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVariableGroupsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVariableGroupsByIDOK), nil

}

/*
PutOrganizationProjectApisDistributedtaskVariablegroupsGroupID Update a variable group.
*/
func (a *Client) PutOrganizationProjectApisDistributedtaskVariablegroupsGroupID(params *PutOrganizationProjectApisDistributedtaskVariablegroupsGroupIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrganizationProjectApisDistributedtaskVariablegroupsGroupIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrganizationProjectApisDistributedtaskVariablegroupsGroupIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrganizationProjectApisDistributedtaskVariablegroupsGroupID",
		Method:             "PUT",
		PathPattern:        "/{organization}/{project}/_apis/distributedtask/variablegroups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrganizationProjectApisDistributedtaskVariablegroupsGroupIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrganizationProjectApisDistributedtaskVariablegroupsGroupIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
