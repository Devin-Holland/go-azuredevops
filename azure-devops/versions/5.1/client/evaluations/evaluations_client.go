// Code generated by go-swagger; DO NOT EDIT.

package evaluations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new evaluations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for evaluations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RequeuePolicyEvaluation Requeue the policy evaluation.

Some policies define a "requeue" action which performs some policy-specific operation.
You can trigger this operation by updating an existing policy evaluation and setting the
PolicyEvaluationRecord.Status field to Queued.
Although any policy evaluation can be requeued, at present only build policies perform any action
in response. Requeueing a build policy will queue a new build to run (cancelling any existing build which
is running).
*/
func (a *Client) RequeuePolicyEvaluation(params *RequeuePolicyEvaluationParams, authInfo runtime.ClientAuthInfoWriter) (*RequeuePolicyEvaluationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequeuePolicyEvaluationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Requeue Policy Evaluation",
		Method:             "PATCH",
		PathPattern:        "/{organization}/{project}/_apis/policy/evaluations/{evaluationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequeuePolicyEvaluationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequeuePolicyEvaluationOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
