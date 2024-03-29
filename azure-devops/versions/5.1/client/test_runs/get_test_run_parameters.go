// Code generated by go-swagger; DO NOT EDIT.

package test_runs

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

// NewGetTestRunParams creates a new GetTestRunParams object
// with the default values initialized.
func NewGetTestRunParams() *GetTestRunParams {
	var ()
	return &GetTestRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestRunParamsWithTimeout creates a new GetTestRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestRunParamsWithTimeout(timeout time.Duration) *GetTestRunParams {
	var ()
	return &GetTestRunParams{

		timeout: timeout,
	}
}

// NewGetTestRunParamsWithContext creates a new GetTestRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestRunParamsWithContext(ctx context.Context) *GetTestRunParams {
	var ()
	return &GetTestRunParams{

		Context: ctx,
	}
}

// NewGetTestRunParamsWithHTTPClient creates a new GetTestRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestRunParamsWithHTTPClient(client *http.Client) *GetTestRunParams {
	var ()
	return &GetTestRunParams{
		HTTPClient: client,
	}
}

/*GetTestRunParams contains all the parameters to send to the API endpoint
for the get test run operation typically these are written to a http.Request
*/
type GetTestRunParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*TestRunID
	  Unique ID of the test run

	*/
	TestRunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get test run params
func (o *GetTestRunParams) WithTimeout(timeout time.Duration) *GetTestRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test run params
func (o *GetTestRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test run params
func (o *GetTestRunParams) WithContext(ctx context.Context) *GetTestRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test run params
func (o *GetTestRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test run params
func (o *GetTestRunParams) WithHTTPClient(client *http.Client) *GetTestRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test run params
func (o *GetTestRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get test run params
func (o *GetTestRunParams) WithAPIVersion(aPIVersion string) *GetTestRunParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get test run params
func (o *GetTestRunParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the get test run params
func (o *GetTestRunParams) WithOrganization(organization string) *GetTestRunParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get test run params
func (o *GetTestRunParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithTestRunID adds the testRunID to the get test run params
func (o *GetTestRunParams) WithTestRunID(testRunID string) *GetTestRunParams {
	o.SetTestRunID(testRunID)
	return o
}

// SetTestRunID adds the testRunId to the get test run params
func (o *GetTestRunParams) SetTestRunID(testRunID string) {
	o.TestRunID = testRunID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param testRunId
	if err := r.SetPathParam("testRunId", o.TestRunID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
