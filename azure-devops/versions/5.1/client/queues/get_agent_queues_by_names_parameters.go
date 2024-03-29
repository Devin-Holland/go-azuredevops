// Code generated by go-swagger; DO NOT EDIT.

package queues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAgentQueuesByNamesParams creates a new GetAgentQueuesByNamesParams object
// with the default values initialized.
func NewGetAgentQueuesByNamesParams() *GetAgentQueuesByNamesParams {
	var ()
	return &GetAgentQueuesByNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentQueuesByNamesParamsWithTimeout creates a new GetAgentQueuesByNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAgentQueuesByNamesParamsWithTimeout(timeout time.Duration) *GetAgentQueuesByNamesParams {
	var ()
	return &GetAgentQueuesByNamesParams{

		timeout: timeout,
	}
}

// NewGetAgentQueuesByNamesParamsWithContext creates a new GetAgentQueuesByNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAgentQueuesByNamesParamsWithContext(ctx context.Context) *GetAgentQueuesByNamesParams {
	var ()
	return &GetAgentQueuesByNamesParams{

		Context: ctx,
	}
}

// NewGetAgentQueuesByNamesParamsWithHTTPClient creates a new GetAgentQueuesByNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAgentQueuesByNamesParamsWithHTTPClient(client *http.Client) *GetAgentQueuesByNamesParams {
	var ()
	return &GetAgentQueuesByNamesParams{
		HTTPClient: client,
	}
}

/*GetAgentQueuesByNamesParams contains all the parameters to send to the API endpoint
for the get agent queues by names operation typically these are written to a http.Request
*/
type GetAgentQueuesByNamesParams struct {

	/*ActionFilter
	  Filter by whether the calling user has use or manage permissions

	*/
	ActionFilter *string
	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*QueueNames
	  A comma-separated list of agent names to retrieve

	*/
	QueueNames string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) WithTimeout(timeout time.Duration) *GetAgentQueuesByNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) WithContext(ctx context.Context) *GetAgentQueuesByNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) WithHTTPClient(client *http.Client) *GetAgentQueuesByNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionFilter adds the actionFilter to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) WithActionFilter(actionFilter *string) *GetAgentQueuesByNamesParams {
	o.SetActionFilter(actionFilter)
	return o
}

// SetActionFilter adds the actionFilter to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) SetActionFilter(actionFilter *string) {
	o.ActionFilter = actionFilter
}

// WithAPIVersion adds the aPIVersion to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) WithAPIVersion(aPIVersion string) *GetAgentQueuesByNamesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) WithOrganization(organization string) *GetAgentQueuesByNamesParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) WithProject(project string) *GetAgentQueuesByNamesParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) SetProject(project string) {
	o.Project = project
}

// WithQueueNames adds the queueNames to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) WithQueueNames(queueNames string) *GetAgentQueuesByNamesParams {
	o.SetQueueNames(queueNames)
	return o
}

// SetQueueNames adds the queueNames to the get agent queues by names params
func (o *GetAgentQueuesByNamesParams) SetQueueNames(queueNames string) {
	o.QueueNames = queueNames
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentQueuesByNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionFilter != nil {

		// query param actionFilter
		var qrActionFilter string
		if o.ActionFilter != nil {
			qrActionFilter = *o.ActionFilter
		}
		qActionFilter := qrActionFilter
		if qActionFilter != "" {
			if err := r.SetQueryParam("actionFilter", qActionFilter); err != nil {
				return err
			}
		}

	}

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// query param queueNames
	qrQueueNames := o.QueueNames
	qQueueNames := qrQueueNames
	if qQueueNames != "" {
		if err := r.SetQueryParam("queueNames", qQueueNames); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
