// Code generated by go-swagger; DO NOT EDIT.

package reporting_work_item_revisions

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

// NewReadReportingRevisionsPostParams creates a new ReadReportingRevisionsPostParams object
// with the default values initialized.
func NewReadReportingRevisionsPostParams() *ReadReportingRevisionsPostParams {
	var ()
	return &ReadReportingRevisionsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadReportingRevisionsPostParamsWithTimeout creates a new ReadReportingRevisionsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadReportingRevisionsPostParamsWithTimeout(timeout time.Duration) *ReadReportingRevisionsPostParams {
	var ()
	return &ReadReportingRevisionsPostParams{

		timeout: timeout,
	}
}

// NewReadReportingRevisionsPostParamsWithContext creates a new ReadReportingRevisionsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadReportingRevisionsPostParamsWithContext(ctx context.Context) *ReadReportingRevisionsPostParams {
	var ()
	return &ReadReportingRevisionsPostParams{

		Context: ctx,
	}
}

// NewReadReportingRevisionsPostParamsWithHTTPClient creates a new ReadReportingRevisionsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadReportingRevisionsPostParamsWithHTTPClient(client *http.Client) *ReadReportingRevisionsPostParams {
	var ()
	return &ReadReportingRevisionsPostParams{
		HTTPClient: client,
	}
}

/*ReadReportingRevisionsPostParams contains all the parameters to send to the API endpoint
for the read reporting revisions post operation typically these are written to a http.Request
*/
type ReadReportingRevisionsPostParams struct {

	/*NrDollarExpand*/
	DollarExpand *string
	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  An object that contains request settings: field filter, type filter, identity format

	*/
	Body *models.ReportingWorkItemRevisionsFilter
	/*ContinuationToken
	  Specifies the watermark to start the batch from. Omit this parameter to get the first batch of revisions.

	*/
	ContinuationToken *string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*StartDateTime
	  Date/time to use as a starting point for revisions, all revisions will occur after this date/time. Cannot be used in conjunction with 'watermark' parameter.

	*/
	StartDateTime *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithTimeout(timeout time.Duration) *ReadReportingRevisionsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithContext(ctx context.Context) *ReadReportingRevisionsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithHTTPClient(client *http.Client) *ReadReportingRevisionsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarExpand adds the dollarExpand to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithDollarExpand(dollarExpand *string) *ReadReportingRevisionsPostParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithAPIVersion adds the aPIVersion to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithAPIVersion(aPIVersion string) *ReadReportingRevisionsPostParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithBody(body *models.ReportingWorkItemRevisionsFilter) *ReadReportingRevisionsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetBody(body *models.ReportingWorkItemRevisionsFilter) {
	o.Body = body
}

// WithContinuationToken adds the continuationToken to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithContinuationToken(continuationToken *string) *ReadReportingRevisionsPostParams {
	o.SetContinuationToken(continuationToken)
	return o
}

// SetContinuationToken adds the continuationToken to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetContinuationToken(continuationToken *string) {
	o.ContinuationToken = continuationToken
}

// WithOrganization adds the organization to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithOrganization(organization string) *ReadReportingRevisionsPostParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithProject(project string) *ReadReportingRevisionsPostParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetProject(project string) {
	o.Project = project
}

// WithStartDateTime adds the startDateTime to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) WithStartDateTime(startDateTime *strfmt.DateTime) *ReadReportingRevisionsPostParams {
	o.SetStartDateTime(startDateTime)
	return o
}

// SetStartDateTime adds the startDateTime to the read reporting revisions post params
func (o *ReadReportingRevisionsPostParams) SetStartDateTime(startDateTime *strfmt.DateTime) {
	o.StartDateTime = startDateTime
}

// WriteToRequest writes these params to a swagger request
func (o *ReadReportingRevisionsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarExpand != nil {

		// query param $expand
		var qrNrDollarExpand string
		if o.DollarExpand != nil {
			qrNrDollarExpand = *o.DollarExpand
		}
		qNrDollarExpand := qrNrDollarExpand
		if qNrDollarExpand != "" {
			if err := r.SetQueryParam("$expand", qNrDollarExpand); err != nil {
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

	if o.StartDateTime != nil {

		// query param startDateTime
		var qrStartDateTime strfmt.DateTime
		if o.StartDateTime != nil {
			qrStartDateTime = *o.StartDateTime
		}
		qStartDateTime := qrStartDateTime.String()
		if qStartDateTime != "" {
			if err := r.SetQueryParam("startDateTime", qStartDateTime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
