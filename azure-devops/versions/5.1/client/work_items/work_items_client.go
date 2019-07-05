// Code generated by go-swagger; DO NOT EDIT.

package work_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new work items API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for work items API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetWorkItem Returns a single work item.
*/
func (a *Client) GetWorkItem(params *GetWorkItemParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Work Item",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/wit/workitems/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkItemOK), nil

}

/*
GetWorkItemTemplate Returns a single work item from a template.
*/
func (a *Client) GetWorkItemTemplate(params *GetWorkItemTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkItemTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkItemTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Work Item Template",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/wit/workitems/${type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkItemTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkItemTemplateOK), nil

}

/*
GetWorkItemsBatch Gets work items for a list of work item ids (Maximum 200)
*/
func (a *Client) GetWorkItemsBatch(params *GetWorkItemsBatchParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkItemsBatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkItemsBatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Work Items Batch",
		Method:             "POST",
		PathPattern:        "/{organization}/{project}/_apis/wit/workitemsbatch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkItemsBatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkItemsBatchOK), nil

}

/*
Update Updates a single work item.
*/
func (a *Client) Update(params *UpdateParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Update",
		Method:             "PATCH",
		PathPattern:        "/{organization}/{project}/_apis/wit/workitems/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json-patch+json"},
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