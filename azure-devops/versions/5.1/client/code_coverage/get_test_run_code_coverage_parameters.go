// Code generated by go-swagger; DO NOT EDIT.

package code_coverage

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

// NewGetTestRunCodeCoverageParams creates a new GetTestRunCodeCoverageParams object
// with the default values initialized.
func NewGetTestRunCodeCoverageParams() *GetTestRunCodeCoverageParams {
	var ()
	return &GetTestRunCodeCoverageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestRunCodeCoverageParamsWithTimeout creates a new GetTestRunCodeCoverageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestRunCodeCoverageParamsWithTimeout(timeout time.Duration) *GetTestRunCodeCoverageParams {
	var ()
	return &GetTestRunCodeCoverageParams{

		timeout: timeout,
	}
}

// NewGetTestRunCodeCoverageParamsWithContext creates a new GetTestRunCodeCoverageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestRunCodeCoverageParamsWithContext(ctx context.Context) *GetTestRunCodeCoverageParams {
	var ()
	return &GetTestRunCodeCoverageParams{

		Context: ctx,
	}
}

// NewGetTestRunCodeCoverageParamsWithHTTPClient creates a new GetTestRunCodeCoverageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestRunCodeCoverageParamsWithHTTPClient(client *http.Client) *GetTestRunCodeCoverageParams {
	var ()
	return &GetTestRunCodeCoverageParams{
		HTTPClient: client,
	}
}

/*GetTestRunCodeCoverageParams contains all the parameters to send to the API endpoint
for the get test run code coverage operation typically these are written to a http.Request
*/
type GetTestRunCodeCoverageParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Flags
	  Value of flags determine the level of code coverage details to be fetched. Flags are additive. Expected Values are 1 for Modules, 2 for Functions, 4 for BlockData.

	*/
	Flags int32
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*RunID
	  ID of the test run for which code coverage data needs to be fetched.

	*/
	RunID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) WithTimeout(timeout time.Duration) *GetTestRunCodeCoverageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) WithContext(ctx context.Context) *GetTestRunCodeCoverageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) WithHTTPClient(client *http.Client) *GetTestRunCodeCoverageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) WithAPIVersion(aPIVersion string) *GetTestRunCodeCoverageParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithFlags adds the flags to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) WithFlags(flags int32) *GetTestRunCodeCoverageParams {
	o.SetFlags(flags)
	return o
}

// SetFlags adds the flags to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) SetFlags(flags int32) {
	o.Flags = flags
}

// WithOrganization adds the organization to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) WithOrganization(organization string) *GetTestRunCodeCoverageParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) WithProject(project string) *GetTestRunCodeCoverageParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) SetProject(project string) {
	o.Project = project
}

// WithRunID adds the runID to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) WithRunID(runID int32) *GetTestRunCodeCoverageParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get test run code coverage params
func (o *GetTestRunCodeCoverageParams) SetRunID(runID int32) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestRunCodeCoverageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param flags
	qrFlags := o.Flags
	qFlags := swag.FormatInt32(qrFlags)
	if qFlags != "" {
		if err := r.SetQueryParam("flags", qFlags); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
