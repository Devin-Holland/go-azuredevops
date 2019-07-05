// Code generated by go-swagger; DO NOT EDIT.

package code_search_results

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

// NewFetchCodeSearchResultsParams creates a new FetchCodeSearchResultsParams object
// with the default values initialized.
func NewFetchCodeSearchResultsParams() *FetchCodeSearchResultsParams {
	var ()
	return &FetchCodeSearchResultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFetchCodeSearchResultsParamsWithTimeout creates a new FetchCodeSearchResultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFetchCodeSearchResultsParamsWithTimeout(timeout time.Duration) *FetchCodeSearchResultsParams {
	var ()
	return &FetchCodeSearchResultsParams{

		timeout: timeout,
	}
}

// NewFetchCodeSearchResultsParamsWithContext creates a new FetchCodeSearchResultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFetchCodeSearchResultsParamsWithContext(ctx context.Context) *FetchCodeSearchResultsParams {
	var ()
	return &FetchCodeSearchResultsParams{

		Context: ctx,
	}
}

// NewFetchCodeSearchResultsParamsWithHTTPClient creates a new FetchCodeSearchResultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFetchCodeSearchResultsParamsWithHTTPClient(client *http.Client) *FetchCodeSearchResultsParams {
	var ()
	return &FetchCodeSearchResultsParams{
		HTTPClient: client,
	}
}

/*FetchCodeSearchResultsParams contains all the parameters to send to the API endpoint
for the fetch code search results operation typically these are written to a http.Request
*/
type FetchCodeSearchResultsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  The Code Search Request.

	*/
	Body *models.CodeSearchRequest
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

// WithTimeout adds the timeout to the fetch code search results params
func (o *FetchCodeSearchResultsParams) WithTimeout(timeout time.Duration) *FetchCodeSearchResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fetch code search results params
func (o *FetchCodeSearchResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fetch code search results params
func (o *FetchCodeSearchResultsParams) WithContext(ctx context.Context) *FetchCodeSearchResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fetch code search results params
func (o *FetchCodeSearchResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fetch code search results params
func (o *FetchCodeSearchResultsParams) WithHTTPClient(client *http.Client) *FetchCodeSearchResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fetch code search results params
func (o *FetchCodeSearchResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the fetch code search results params
func (o *FetchCodeSearchResultsParams) WithAPIVersion(aPIVersion string) *FetchCodeSearchResultsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the fetch code search results params
func (o *FetchCodeSearchResultsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the fetch code search results params
func (o *FetchCodeSearchResultsParams) WithBody(body *models.CodeSearchRequest) *FetchCodeSearchResultsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the fetch code search results params
func (o *FetchCodeSearchResultsParams) SetBody(body *models.CodeSearchRequest) {
	o.Body = body
}

// WithOrganization adds the organization to the fetch code search results params
func (o *FetchCodeSearchResultsParams) WithOrganization(organization string) *FetchCodeSearchResultsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the fetch code search results params
func (o *FetchCodeSearchResultsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the fetch code search results params
func (o *FetchCodeSearchResultsParams) WithProject(project string) *FetchCodeSearchResultsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the fetch code search results params
func (o *FetchCodeSearchResultsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *FetchCodeSearchResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
