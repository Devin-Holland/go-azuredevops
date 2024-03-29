// Code generated by go-swagger; DO NOT EDIT.

package download_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new download log API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for download log API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DownloadLog Downloads audit log entries.
*/
func (a *Client) DownloadLog(params *DownloadLogParams, authInfo runtime.ClientAuthInfoWriter) (*DownloadLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Download Log",
		Method:             "GET",
		PathPattern:        "/{organization}/_apis/audit/downloadlog",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownloadLogOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
