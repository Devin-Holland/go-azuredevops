// Code generated by go-swagger; DO NOT EDIT.

package points

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// NewGetPointsByQueryParams creates a new GetPointsByQueryParams object
// with the default values initialized.
func NewGetPointsByQueryParams() *GetPointsByQueryParams {
	var ()
	return &GetPointsByQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPointsByQueryParamsWithTimeout creates a new GetPointsByQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPointsByQueryParamsWithTimeout(timeout time.Duration) *GetPointsByQueryParams {
	var ()
	return &GetPointsByQueryParams{

		timeout: timeout,
	}
}

// NewGetPointsByQueryParamsWithContext creates a new GetPointsByQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPointsByQueryParamsWithContext(ctx context.Context) *GetPointsByQueryParams {
	var ()
	return &GetPointsByQueryParams{

		Context: ctx,
	}
}

// NewGetPointsByQueryParamsWithHTTPClient creates a new GetPointsByQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPointsByQueryParamsWithHTTPClient(client *http.Client) *GetPointsByQueryParams {
	var ()
	return &GetPointsByQueryParams{
		HTTPClient: client,
	}
}

/*GetPointsByQueryParams contains all the parameters to send to the API endpoint
for the get points by query operation typically these are written to a http.Request
*/
type GetPointsByQueryParams struct {

	/*NrDollarSkip
	  Number of test points to skip..

	*/
	DollarSkip *int32
	/*NrDollarTop
	  Number of test points to return.

	*/
	DollarTop *int32
	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  TestPointsQuery to get test points.

	*/
	Body *models.TestPointsQuery
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

// WithTimeout adds the timeout to the get points by query params
func (o *GetPointsByQueryParams) WithTimeout(timeout time.Duration) *GetPointsByQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get points by query params
func (o *GetPointsByQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get points by query params
func (o *GetPointsByQueryParams) WithContext(ctx context.Context) *GetPointsByQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get points by query params
func (o *GetPointsByQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get points by query params
func (o *GetPointsByQueryParams) WithHTTPClient(client *http.Client) *GetPointsByQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get points by query params
func (o *GetPointsByQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarSkip adds the dollarSkip to the get points by query params
func (o *GetPointsByQueryParams) WithDollarSkip(dollarSkip *int32) *GetPointsByQueryParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get points by query params
func (o *GetPointsByQueryParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get points by query params
func (o *GetPointsByQueryParams) WithDollarTop(dollarTop *int32) *GetPointsByQueryParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get points by query params
func (o *GetPointsByQueryParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get points by query params
func (o *GetPointsByQueryParams) WithAPIVersion(aPIVersion string) *GetPointsByQueryParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get points by query params
func (o *GetPointsByQueryParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the get points by query params
func (o *GetPointsByQueryParams) WithBody(body *models.TestPointsQuery) *GetPointsByQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get points by query params
func (o *GetPointsByQueryParams) SetBody(body *models.TestPointsQuery) {
	o.Body = body
}

// WithOrganization adds the organization to the get points by query params
func (o *GetPointsByQueryParams) WithOrganization(organization string) *GetPointsByQueryParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get points by query params
func (o *GetPointsByQueryParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get points by query params
func (o *GetPointsByQueryParams) WithProject(project string) *GetPointsByQueryParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get points by query params
func (o *GetPointsByQueryParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetPointsByQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarSkip != nil {

		// query param $skip
		var qrNrDollarSkip int32
		if o.DollarSkip != nil {
			qrNrDollarSkip = *o.DollarSkip
		}
		qNrDollarSkip := swag.FormatInt32(qrNrDollarSkip)
		if qNrDollarSkip != "" {
			if err := r.SetQueryParam("$skip", qNrDollarSkip); err != nil {
				return err
			}
		}

	}

	if o.DollarTop != nil {

		// query param $top
		var qrNrDollarTop int32
		if o.DollarTop != nil {
			qrNrDollarTop = *o.DollarTop
		}
		qNrDollarTop := swag.FormatInt32(qrNrDollarTop)
		if qNrDollarTop != "" {
			if err := r.SetQueryParam("$top", qNrDollarTop); err != nil {
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
