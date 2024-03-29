// Code generated by go-swagger; DO NOT EDIT.

package attachments

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

// NewGetTestRunAttachmentZipParams creates a new GetTestRunAttachmentZipParams object
// with the default values initialized.
func NewGetTestRunAttachmentZipParams() *GetTestRunAttachmentZipParams {
	var ()
	return &GetTestRunAttachmentZipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestRunAttachmentZipParamsWithTimeout creates a new GetTestRunAttachmentZipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestRunAttachmentZipParamsWithTimeout(timeout time.Duration) *GetTestRunAttachmentZipParams {
	var ()
	return &GetTestRunAttachmentZipParams{

		timeout: timeout,
	}
}

// NewGetTestRunAttachmentZipParamsWithContext creates a new GetTestRunAttachmentZipParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestRunAttachmentZipParamsWithContext(ctx context.Context) *GetTestRunAttachmentZipParams {
	var ()
	return &GetTestRunAttachmentZipParams{

		Context: ctx,
	}
}

// NewGetTestRunAttachmentZipParamsWithHTTPClient creates a new GetTestRunAttachmentZipParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestRunAttachmentZipParamsWithHTTPClient(client *http.Client) *GetTestRunAttachmentZipParams {
	var ()
	return &GetTestRunAttachmentZipParams{
		HTTPClient: client,
	}
}

/*GetTestRunAttachmentZipParams contains all the parameters to send to the API endpoint
for the get test run attachment zip operation typically these are written to a http.Request
*/
type GetTestRunAttachmentZipParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*AttachmentID
	  ID of the test run attachment to be downloaded.

	*/
	AttachmentID int32
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*RunID
	  ID of the test run whose attachment has to be downloaded.

	*/
	RunID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) WithTimeout(timeout time.Duration) *GetTestRunAttachmentZipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) WithContext(ctx context.Context) *GetTestRunAttachmentZipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) WithHTTPClient(client *http.Client) *GetTestRunAttachmentZipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) WithAPIVersion(aPIVersion string) *GetTestRunAttachmentZipParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithAttachmentID adds the attachmentID to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) WithAttachmentID(attachmentID int32) *GetTestRunAttachmentZipParams {
	o.SetAttachmentID(attachmentID)
	return o
}

// SetAttachmentID adds the attachmentId to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) SetAttachmentID(attachmentID int32) {
	o.AttachmentID = attachmentID
}

// WithOrganization adds the organization to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) WithOrganization(organization string) *GetTestRunAttachmentZipParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) WithProject(project string) *GetTestRunAttachmentZipParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) SetProject(project string) {
	o.Project = project
}

// WithRunID adds the runID to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) WithRunID(runID int32) *GetTestRunAttachmentZipParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get test run attachment zip params
func (o *GetTestRunAttachmentZipParams) SetRunID(runID int32) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestRunAttachmentZipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param attachmentId
	if err := r.SetPathParam("attachmentId", swag.FormatInt32(o.AttachmentID)); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param runId
	if err := r.SetPathParam("runId", swag.FormatInt32(o.RunID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
