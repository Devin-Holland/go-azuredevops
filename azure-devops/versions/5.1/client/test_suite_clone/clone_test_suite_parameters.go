// Code generated by go-swagger; DO NOT EDIT.

package test_suite_clone

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

// NewCloneTestSuiteParams creates a new CloneTestSuiteParams object
// with the default values initialized.
func NewCloneTestSuiteParams() *CloneTestSuiteParams {
	var ()
	return &CloneTestSuiteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloneTestSuiteParamsWithTimeout creates a new CloneTestSuiteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloneTestSuiteParamsWithTimeout(timeout time.Duration) *CloneTestSuiteParams {
	var ()
	return &CloneTestSuiteParams{

		timeout: timeout,
	}
}

// NewCloneTestSuiteParamsWithContext creates a new CloneTestSuiteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloneTestSuiteParamsWithContext(ctx context.Context) *CloneTestSuiteParams {
	var ()
	return &CloneTestSuiteParams{

		Context: ctx,
	}
}

// NewCloneTestSuiteParamsWithHTTPClient creates a new CloneTestSuiteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloneTestSuiteParamsWithHTTPClient(client *http.Client) *CloneTestSuiteParams {
	var ()
	return &CloneTestSuiteParams{
		HTTPClient: client,
	}
}

/*CloneTestSuiteParams contains all the parameters to send to the API endpoint
for the clone test suite operation typically these are written to a http.Request
*/
type CloneTestSuiteParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  Suite Clone Request Body detail TestSuiteCloneRequest

	*/
	Body *models.CloneTestSuiteParams
	/*DeepClone
	  Clones all the associated test cases as well

	*/
	DeepClone *bool
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

// WithTimeout adds the timeout to the clone test suite params
func (o *CloneTestSuiteParams) WithTimeout(timeout time.Duration) *CloneTestSuiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone test suite params
func (o *CloneTestSuiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone test suite params
func (o *CloneTestSuiteParams) WithContext(ctx context.Context) *CloneTestSuiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone test suite params
func (o *CloneTestSuiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone test suite params
func (o *CloneTestSuiteParams) WithHTTPClient(client *http.Client) *CloneTestSuiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone test suite params
func (o *CloneTestSuiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the clone test suite params
func (o *CloneTestSuiteParams) WithAPIVersion(aPIVersion string) *CloneTestSuiteParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the clone test suite params
func (o *CloneTestSuiteParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the clone test suite params
func (o *CloneTestSuiteParams) WithBody(body *models.CloneTestSuiteParams) *CloneTestSuiteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the clone test suite params
func (o *CloneTestSuiteParams) SetBody(body *models.CloneTestSuiteParams) {
	o.Body = body
}

// WithDeepClone adds the deepClone to the clone test suite params
func (o *CloneTestSuiteParams) WithDeepClone(deepClone *bool) *CloneTestSuiteParams {
	o.SetDeepClone(deepClone)
	return o
}

// SetDeepClone adds the deepClone to the clone test suite params
func (o *CloneTestSuiteParams) SetDeepClone(deepClone *bool) {
	o.DeepClone = deepClone
}

// WithOrganization adds the organization to the clone test suite params
func (o *CloneTestSuiteParams) WithOrganization(organization string) *CloneTestSuiteParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the clone test suite params
func (o *CloneTestSuiteParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the clone test suite params
func (o *CloneTestSuiteParams) WithProject(project string) *CloneTestSuiteParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the clone test suite params
func (o *CloneTestSuiteParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CloneTestSuiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DeepClone != nil {

		// query param deepClone
		var qrDeepClone bool
		if o.DeepClone != nil {
			qrDeepClone = *o.DeepClone
		}
		qDeepClone := swag.FormatBool(qrDeepClone)
		if qDeepClone != "" {
			if err := r.SetQueryParam("deepClone", qDeepClone); err != nil {
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