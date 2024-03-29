// Code generated by go-swagger; DO NOT EDIT.

package work_item_revisions_discussions

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
)

// NewReadReportingDiscussionsParams creates a new ReadReportingDiscussionsParams object
// with the default values initialized.
func NewReadReportingDiscussionsParams() *ReadReportingDiscussionsParams {
	var ()
	return &ReadReportingDiscussionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadReportingDiscussionsParamsWithTimeout creates a new ReadReportingDiscussionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadReportingDiscussionsParamsWithTimeout(timeout time.Duration) *ReadReportingDiscussionsParams {
	var ()
	return &ReadReportingDiscussionsParams{

		timeout: timeout,
	}
}

// NewReadReportingDiscussionsParamsWithContext creates a new ReadReportingDiscussionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadReportingDiscussionsParamsWithContext(ctx context.Context) *ReadReportingDiscussionsParams {
	var ()
	return &ReadReportingDiscussionsParams{

		Context: ctx,
	}
}

// NewReadReportingDiscussionsParamsWithHTTPClient creates a new ReadReportingDiscussionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadReportingDiscussionsParamsWithHTTPClient(client *http.Client) *ReadReportingDiscussionsParams {
	var ()
	return &ReadReportingDiscussionsParams{
		HTTPClient: client,
	}
}

/*ReadReportingDiscussionsParams contains all the parameters to send to the API endpoint
for the read reporting discussions operation typically these are written to a http.Request
*/
type ReadReportingDiscussionsParams struct {

	/*NrDollarMaxPageSize*/
	DollarMaxPageSize *int32
	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*ContinuationToken*/
	ContinuationToken *string
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

// WithTimeout adds the timeout to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) WithTimeout(timeout time.Duration) *ReadReportingDiscussionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) WithContext(ctx context.Context) *ReadReportingDiscussionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) WithHTTPClient(client *http.Client) *ReadReportingDiscussionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarMaxPageSize adds the dollarMaxPageSize to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) WithDollarMaxPageSize(dollarMaxPageSize *int32) *ReadReportingDiscussionsParams {
	o.SetDollarMaxPageSize(dollarMaxPageSize)
	return o
}

// SetDollarMaxPageSize adds the dollarMaxPageSize to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) SetDollarMaxPageSize(dollarMaxPageSize *int32) {
	o.DollarMaxPageSize = dollarMaxPageSize
}

// WithAPIVersion adds the aPIVersion to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) WithAPIVersion(aPIVersion string) *ReadReportingDiscussionsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithContinuationToken adds the continuationToken to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) WithContinuationToken(continuationToken *string) *ReadReportingDiscussionsParams {
	o.SetContinuationToken(continuationToken)
	return o
}

// SetContinuationToken adds the continuationToken to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) SetContinuationToken(continuationToken *string) {
	o.ContinuationToken = continuationToken
}

// WithOrganization adds the organization to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) WithOrganization(organization string) *ReadReportingDiscussionsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) WithProject(project string) *ReadReportingDiscussionsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the read reporting discussions params
func (o *ReadReportingDiscussionsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *ReadReportingDiscussionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarMaxPageSize != nil {

		// query param $maxPageSize
		var qrNrDollarMaxPageSize int32
		if o.DollarMaxPageSize != nil {
			qrNrDollarMaxPageSize = *o.DollarMaxPageSize
		}
		qNrDollarMaxPageSize := swag.FormatInt32(qrNrDollarMaxPageSize)
		if qNrDollarMaxPageSize != "" {
			if err := r.SetQueryParam("$maxPageSize", qNrDollarMaxPageSize); err != nil {
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

	if o.ContinuationToken != nil {

		// query param continuationToken
		var qrContinuationToken string
		if o.ContinuationToken != nil {
			qrContinuationToken = *o.ContinuationToken
		}
		qContinuationToken := qrContinuationToken
		if qContinuationToken != "" {
			if err := r.SetQueryParam("continuationToken", qContinuationToken); err != nil {
				return err
			}
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
