// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PutOrganizationApisWorkProcessesProcessIDWorkItemTypesWitRefNameRulesRuleID Updates a rule in the work item type of the process.
*/
func (a *Client) PutOrganizationApisWorkProcessesProcessIDWorkItemTypesWitRefNameRulesRuleID(params *PutOrganizationApisWorkProcessesProcessIDWorkItemTypesWitRefNameRulesRuleIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrganizationApisWorkProcessesProcessIDWorkItemTypesWitRefNameRulesRuleIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrganizationApisWorkProcessesProcessIDWorkItemTypesWitRefNameRulesRuleIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrganizationApisWorkProcessesProcessIDWorkItemTypesWitRefNameRulesRuleID",
		Method:             "PUT",
		PathPattern:        "/{organization}/_apis/work/processes/{processId}/workItemTypes/{witRefName}/rules/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrganizationApisWorkProcessesProcessIDWorkItemTypesWitRefNameRulesRuleIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrganizationApisWorkProcessesProcessIDWorkItemTypesWitRefNameRulesRuleIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}