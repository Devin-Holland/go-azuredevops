// Code generated by go-swagger; DO NOT EDIT.

package package_search_results

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

// NewFetchPackageSearchResultsParams creates a new FetchPackageSearchResultsParams object
// with the default values initialized.
func NewFetchPackageSearchResultsParams() *FetchPackageSearchResultsParams {
	var ()
	return &FetchPackageSearchResultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFetchPackageSearchResultsParamsWithTimeout creates a new FetchPackageSearchResultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFetchPackageSearchResultsParamsWithTimeout(timeout time.Duration) *FetchPackageSearchResultsParams {
	var ()
	return &FetchPackageSearchResultsParams{

		timeout: timeout,
	}
}

// NewFetchPackageSearchResultsParamsWithContext creates a new FetchPackageSearchResultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFetchPackageSearchResultsParamsWithContext(ctx context.Context) *FetchPackageSearchResultsParams {
	var ()
	return &FetchPackageSearchResultsParams{

		Context: ctx,
	}
}

// NewFetchPackageSearchResultsParamsWithHTTPClient creates a new FetchPackageSearchResultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFetchPackageSearchResultsParamsWithHTTPClient(client *http.Client) *FetchPackageSearchResultsParams {
	var ()
	return &FetchPackageSearchResultsParams{
		HTTPClient: client,
	}
}

/*FetchPackageSearchResultsParams contains all the parameters to send to the API endpoint
for the fetch package search results operation typically these are written to a http.Request
*/
type FetchPackageSearchResultsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  The Package Search Request.

	*/
	Body *models.PackageSearchRequest
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the fetch package search results params
func (o *FetchPackageSearchResultsParams) WithTimeout(timeout time.Duration) *FetchPackageSearchResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fetch package search results params
func (o *FetchPackageSearchResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fetch package search results params
func (o *FetchPackageSearchResultsParams) WithContext(ctx context.Context) *FetchPackageSearchResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fetch package search results params
func (o *FetchPackageSearchResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fetch package search results params
func (o *FetchPackageSearchResultsParams) WithHTTPClient(client *http.Client) *FetchPackageSearchResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fetch package search results params
func (o *FetchPackageSearchResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the fetch package search results params
func (o *FetchPackageSearchResultsParams) WithAPIVersion(aPIVersion string) *FetchPackageSearchResultsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the fetch package search results params
func (o *FetchPackageSearchResultsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the fetch package search results params
func (o *FetchPackageSearchResultsParams) WithBody(body *models.PackageSearchRequest) *FetchPackageSearchResultsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the fetch package search results params
func (o *FetchPackageSearchResultsParams) SetBody(body *models.PackageSearchRequest) {
	o.Body = body
}

// WithOrganization adds the organization to the fetch package search results params
func (o *FetchPackageSearchResultsParams) WithOrganization(organization string) *FetchPackageSearchResultsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the fetch package search results params
func (o *FetchPackageSearchResultsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *FetchPackageSearchResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
