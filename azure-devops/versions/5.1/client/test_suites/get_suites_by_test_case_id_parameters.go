// Code generated by go-swagger; DO NOT EDIT.

package test_suites

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

// NewGetSuitesByTestCaseIDParams creates a new GetSuitesByTestCaseIDParams object
// with the default values initialized.
func NewGetSuitesByTestCaseIDParams() *GetSuitesByTestCaseIDParams {
	var ()
	return &GetSuitesByTestCaseIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSuitesByTestCaseIDParamsWithTimeout creates a new GetSuitesByTestCaseIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSuitesByTestCaseIDParamsWithTimeout(timeout time.Duration) *GetSuitesByTestCaseIDParams {
	var ()
	return &GetSuitesByTestCaseIDParams{

		timeout: timeout,
	}
}

// NewGetSuitesByTestCaseIDParamsWithContext creates a new GetSuitesByTestCaseIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSuitesByTestCaseIDParamsWithContext(ctx context.Context) *GetSuitesByTestCaseIDParams {
	var ()
	return &GetSuitesByTestCaseIDParams{

		Context: ctx,
	}
}

// NewGetSuitesByTestCaseIDParamsWithHTTPClient creates a new GetSuitesByTestCaseIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSuitesByTestCaseIDParamsWithHTTPClient(client *http.Client) *GetSuitesByTestCaseIDParams {
	var ()
	return &GetSuitesByTestCaseIDParams{
		HTTPClient: client,
	}
}

/*GetSuitesByTestCaseIDParams contains all the parameters to send to the API endpoint
for the get suites by test case Id operation typically these are written to a http.Request
*/
type GetSuitesByTestCaseIDParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*TestCaseID
	  ID of the test case for which suites need to be fetched.

	*/
	TestCaseID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) WithTimeout(timeout time.Duration) *GetSuitesByTestCaseIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) WithContext(ctx context.Context) *GetSuitesByTestCaseIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) WithHTTPClient(client *http.Client) *GetSuitesByTestCaseIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) WithAPIVersion(aPIVersion string) *GetSuitesByTestCaseIDParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) WithOrganization(organization string) *GetSuitesByTestCaseIDParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithTestCaseID adds the testCaseID to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) WithTestCaseID(testCaseID int32) *GetSuitesByTestCaseIDParams {
	o.SetTestCaseID(testCaseID)
	return o
}

// SetTestCaseID adds the testCaseId to the get suites by test case Id params
func (o *GetSuitesByTestCaseIDParams) SetTestCaseID(testCaseID int32) {
	o.TestCaseID = testCaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSuitesByTestCaseIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param testCaseId
	qrTestCaseID := o.TestCaseID
	qTestCaseID := swag.FormatInt32(qrTestCaseID)
	if qTestCaseID != "" {
		if err := r.SetQueryParam("testCaseId", qTestCaseID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}