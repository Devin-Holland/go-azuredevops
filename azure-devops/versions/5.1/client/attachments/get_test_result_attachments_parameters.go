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

// NewGetTestResultAttachmentsParams creates a new GetTestResultAttachmentsParams object
// with the default values initialized.
func NewGetTestResultAttachmentsParams() *GetTestResultAttachmentsParams {
	var ()
	return &GetTestResultAttachmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestResultAttachmentsParamsWithTimeout creates a new GetTestResultAttachmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestResultAttachmentsParamsWithTimeout(timeout time.Duration) *GetTestResultAttachmentsParams {
	var ()
	return &GetTestResultAttachmentsParams{

		timeout: timeout,
	}
}

// NewGetTestResultAttachmentsParamsWithContext creates a new GetTestResultAttachmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestResultAttachmentsParamsWithContext(ctx context.Context) *GetTestResultAttachmentsParams {
	var ()
	return &GetTestResultAttachmentsParams{

		Context: ctx,
	}
}

// NewGetTestResultAttachmentsParamsWithHTTPClient creates a new GetTestResultAttachmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestResultAttachmentsParamsWithHTTPClient(client *http.Client) *GetTestResultAttachmentsParams {
	var ()
	return &GetTestResultAttachmentsParams{
		HTTPClient: client,
	}
}

/*GetTestResultAttachmentsParams contains all the parameters to send to the API endpoint
for the get test result attachments operation typically these are written to a http.Request
*/
type GetTestResultAttachmentsParams struct {

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
	/*RunID
	  ID of the test run that contains the result.

	*/
	RunID int32
	/*TestCaseResultID
	  ID of the test result.

	*/
	TestCaseResultID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get test result attachments params
func (o *GetTestResultAttachmentsParams) WithTimeout(timeout time.Duration) *GetTestResultAttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test result attachments params
func (o *GetTestResultAttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test result attachments params
func (o *GetTestResultAttachmentsParams) WithContext(ctx context.Context) *GetTestResultAttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test result attachments params
func (o *GetTestResultAttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test result attachments params
func (o *GetTestResultAttachmentsParams) WithHTTPClient(client *http.Client) *GetTestResultAttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test result attachments params
func (o *GetTestResultAttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get test result attachments params
func (o *GetTestResultAttachmentsParams) WithAPIVersion(aPIVersion string) *GetTestResultAttachmentsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get test result attachments params
func (o *GetTestResultAttachmentsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the get test result attachments params
func (o *GetTestResultAttachmentsParams) WithOrganization(organization string) *GetTestResultAttachmentsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get test result attachments params
func (o *GetTestResultAttachmentsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get test result attachments params
func (o *GetTestResultAttachmentsParams) WithProject(project string) *GetTestResultAttachmentsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get test result attachments params
func (o *GetTestResultAttachmentsParams) SetProject(project string) {
	o.Project = project
}

// WithRunID adds the runID to the get test result attachments params
func (o *GetTestResultAttachmentsParams) WithRunID(runID int32) *GetTestResultAttachmentsParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get test result attachments params
func (o *GetTestResultAttachmentsParams) SetRunID(runID int32) {
	o.RunID = runID
}

// WithTestCaseResultID adds the testCaseResultID to the get test result attachments params
func (o *GetTestResultAttachmentsParams) WithTestCaseResultID(testCaseResultID int32) *GetTestResultAttachmentsParams {
	o.SetTestCaseResultID(testCaseResultID)
	return o
}

// SetTestCaseResultID adds the testCaseResultId to the get test result attachments params
func (o *GetTestResultAttachmentsParams) SetTestCaseResultID(testCaseResultID int32) {
	o.TestCaseResultID = testCaseResultID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestResultAttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param testCaseResultId
	if err := r.SetPathParam("testCaseResultId", swag.FormatInt32(o.TestCaseResultID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
