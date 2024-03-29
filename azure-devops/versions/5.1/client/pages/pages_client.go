// Code generated by go-swagger; DO NOT EDIT.

package pages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateOrUpdate Creates or edits a wiki page.
*/
func (a *Client) CreateOrUpdate(params *CreateOrUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrUpdateOK, *CreateOrUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Create Or Update",
		Method:             "PUT",
		PathPattern:        "/{organization}/{project}/_apis/wiki/wikis/{wikiIdentifier}/pages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateOrUpdateOK:
		return value, nil, nil
	case *CreateOrUpdateCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeletePage Deletes a wiki page.
*/
func (a *Client) DeletePage(params *DeletePageParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Delete Page",
		Method:             "DELETE",
		PathPattern:        "/{organization}/{project}/_apis/wiki/wikis/{wikiIdentifier}/pages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePageOK), nil

}

/*
DeletePageByID Deletes a wiki page.
*/
func (a *Client) DeletePageByID(params *DeletePageByIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePageByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePageByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Delete Page By Id",
		Method:             "DELETE",
		PathPattern:        "/{organization}/{project}/_apis/wiki/wikis/{wikiIdentifier}/pages/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePageByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePageByIDOK), nil

}

/*
GetPage Gets metadata or content of the wiki page for the provided path. Content negotiation is done based on the `Accept` header sent in the request.
*/
func (a *Client) GetPage(params *GetPageParams, authInfo runtime.ClientAuthInfoWriter) (*GetPageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Page",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/wiki/wikis/{wikiIdentifier}/pages",
		ProducesMediaTypes: []string{"application/json", "application/zip", "text/plain"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPageOK), nil

}

/*
GetPageByID Gets metadata or content of the wiki page for the provided page id. Content negotiation is done based on the `Accept` header sent in the request.
*/
func (a *Client) GetPageByID(params *GetPageByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPageByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPageByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get Page By Id",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/wiki/wikis/{wikiIdentifier}/pages/{id}",
		ProducesMediaTypes: []string{"application/json", "application/zip", "text/plain"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPageByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPageByIDOK), nil

}

/*
Update Edits a wiki page.
*/
func (a *Client) Update(params *UpdateParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Update",
		Method:             "PATCH",
		PathPattern:        "/{organization}/{project}/_apis/wiki/wikis/{wikiIdentifier}/pages/{id}",
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
