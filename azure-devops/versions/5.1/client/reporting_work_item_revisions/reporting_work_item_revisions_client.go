// Code generated by go-swagger; DO NOT EDIT.

package reporting_work_item_revisions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new reporting work item revisions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reporting work item revisions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ReadReportingRevisionsGet Get a batch of work item revisions with the option of including deleted items
*/
func (a *Client) ReadReportingRevisionsGet(params *ReadReportingRevisionsGetParams, authInfo runtime.ClientAuthInfoWriter) (*ReadReportingRevisionsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadReportingRevisionsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Read Reporting Revisions Get",
		Method:             "GET",
		PathPattern:        "/{organization}/{project}/_apis/wit/reporting/workitemrevisions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadReportingRevisionsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReadReportingRevisionsGetOK), nil

}

/*
ReadReportingRevisionsPost Get a batch of work item revisions. This request may be used if your list of fields is large enough that it may run the URL over the length limit.
*/
func (a *Client) ReadReportingRevisionsPost(params *ReadReportingRevisionsPostParams, authInfo runtime.ClientAuthInfoWriter) (*ReadReportingRevisionsPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadReportingRevisionsPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Read Reporting Revisions Post",
		Method:             "POST",
		PathPattern:        "/{organization}/{project}/_apis/wit/reporting/workitemrevisions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadReportingRevisionsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReadReportingRevisionsPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}