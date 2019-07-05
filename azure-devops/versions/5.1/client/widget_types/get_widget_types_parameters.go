// Code generated by go-swagger; DO NOT EDIT.

package widget_types

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

// NewGetWidgetTypesParams creates a new GetWidgetTypesParams object
// with the default values initialized.
func NewGetWidgetTypesParams() *GetWidgetTypesParams {
	var ()
	return &GetWidgetTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWidgetTypesParamsWithTimeout creates a new GetWidgetTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWidgetTypesParamsWithTimeout(timeout time.Duration) *GetWidgetTypesParams {
	var ()
	return &GetWidgetTypesParams{

		timeout: timeout,
	}
}

// NewGetWidgetTypesParamsWithContext creates a new GetWidgetTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWidgetTypesParamsWithContext(ctx context.Context) *GetWidgetTypesParams {
	var ()
	return &GetWidgetTypesParams{

		Context: ctx,
	}
}

// NewGetWidgetTypesParamsWithHTTPClient creates a new GetWidgetTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWidgetTypesParamsWithHTTPClient(client *http.Client) *GetWidgetTypesParams {
	var ()
	return &GetWidgetTypesParams{
		HTTPClient: client,
	}
}

/*GetWidgetTypesParams contains all the parameters to send to the API endpoint
for the get widget types operation typically these are written to a http.Request
*/
type GetWidgetTypesParams struct {

	/*NrDollarScope*/
	DollarScope string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get widget types params
func (o *GetWidgetTypesParams) WithTimeout(timeout time.Duration) *GetWidgetTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get widget types params
func (o *GetWidgetTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get widget types params
func (o *GetWidgetTypesParams) WithContext(ctx context.Context) *GetWidgetTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get widget types params
func (o *GetWidgetTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get widget types params
func (o *GetWidgetTypesParams) WithHTTPClient(client *http.Client) *GetWidgetTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get widget types params
func (o *GetWidgetTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarScope adds the dollarScope to the get widget types params
func (o *GetWidgetTypesParams) WithDollarScope(dollarScope string) *GetWidgetTypesParams {
	o.SetDollarScope(dollarScope)
	return o
}

// SetDollarScope adds the dollarScope to the get widget types params
func (o *GetWidgetTypesParams) SetDollarScope(dollarScope string) {
	o.DollarScope = dollarScope
}

// WithAPIVersion adds the aPIVersion to the get widget types params
func (o *GetWidgetTypesParams) WithAPIVersion(aPIVersion string) *GetWidgetTypesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get widget types params
func (o *GetWidgetTypesParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the get widget types params
func (o *GetWidgetTypesParams) WithOrganization(organization string) *GetWidgetTypesParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get widget types params
func (o *GetWidgetTypesParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get widget types params
func (o *GetWidgetTypesParams) WithProject(project string) *GetWidgetTypesParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get widget types params
func (o *GetWidgetTypesParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetWidgetTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param $scope
	qrNrDollarScope := o.DollarScope
	qNrDollarScope := qrNrDollarScope
	if qNrDollarScope != "" {
		if err := r.SetQueryParam("$scope", qNrDollarScope); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
