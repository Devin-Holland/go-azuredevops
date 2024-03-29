// Code generated by go-swagger; DO NOT EDIT.

package queries

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

	models "azure-devops/5.1/models"
)

// NewGetQueriesBatchParams creates a new GetQueriesBatchParams object
// with the default values initialized.
func NewGetQueriesBatchParams() *GetQueriesBatchParams {
	var ()
	return &GetQueriesBatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetQueriesBatchParamsWithTimeout creates a new GetQueriesBatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetQueriesBatchParamsWithTimeout(timeout time.Duration) *GetQueriesBatchParams {
	var ()
	return &GetQueriesBatchParams{

		timeout: timeout,
	}
}

// NewGetQueriesBatchParamsWithContext creates a new GetQueriesBatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetQueriesBatchParamsWithContext(ctx context.Context) *GetQueriesBatchParams {
	var ()
	return &GetQueriesBatchParams{

		Context: ctx,
	}
}

// NewGetQueriesBatchParamsWithHTTPClient creates a new GetQueriesBatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetQueriesBatchParamsWithHTTPClient(client *http.Client) *GetQueriesBatchParams {
	var ()
	return &GetQueriesBatchParams{
		HTTPClient: client,
	}
}

/*GetQueriesBatchParams contains all the parameters to send to the API endpoint
for the get queries batch operation typically these are written to a http.Request
*/
type GetQueriesBatchParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body*/
	Body *models.QueryBatchGetRequest
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get queries batch params
func (o *GetQueriesBatchParams) WithTimeout(timeout time.Duration) *GetQueriesBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get queries batch params
func (o *GetQueriesBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get queries batch params
func (o *GetQueriesBatchParams) WithContext(ctx context.Context) *GetQueriesBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get queries batch params
func (o *GetQueriesBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get queries batch params
func (o *GetQueriesBatchParams) WithHTTPClient(client *http.Client) *GetQueriesBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get queries batch params
func (o *GetQueriesBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get queries batch params
func (o *GetQueriesBatchParams) WithAPIVersion(aPIVersion string) *GetQueriesBatchParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get queries batch params
func (o *GetQueriesBatchParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the get queries batch params
func (o *GetQueriesBatchParams) WithBody(body *models.QueryBatchGetRequest) *GetQueriesBatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get queries batch params
func (o *GetQueriesBatchParams) SetBody(body *models.QueryBatchGetRequest) {
	o.Body = body
}

// WithOrganization adds the organization to the get queries batch params
func (o *GetQueriesBatchParams) WithOrganization(organization string) *GetQueriesBatchParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get queries batch params
func (o *GetQueriesBatchParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get queries batch params
func (o *GetQueriesBatchParams) WithProject(project string) *GetQueriesBatchParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get queries batch params
func (o *GetQueriesBatchParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetQueriesBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
