// Code generated by go-swagger; DO NOT EDIT.

package folders

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

// NewPutOrganizationProjectApisBuildFoldersParams creates a new PutOrganizationProjectApisBuildFoldersParams object
// with the default values initialized.
func NewPutOrganizationProjectApisBuildFoldersParams() *PutOrganizationProjectApisBuildFoldersParams {
	var ()
	return &PutOrganizationProjectApisBuildFoldersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrganizationProjectApisBuildFoldersParamsWithTimeout creates a new PutOrganizationProjectApisBuildFoldersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOrganizationProjectApisBuildFoldersParamsWithTimeout(timeout time.Duration) *PutOrganizationProjectApisBuildFoldersParams {
	var ()
	return &PutOrganizationProjectApisBuildFoldersParams{

		timeout: timeout,
	}
}

// NewPutOrganizationProjectApisBuildFoldersParamsWithContext creates a new PutOrganizationProjectApisBuildFoldersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOrganizationProjectApisBuildFoldersParamsWithContext(ctx context.Context) *PutOrganizationProjectApisBuildFoldersParams {
	var ()
	return &PutOrganizationProjectApisBuildFoldersParams{

		Context: ctx,
	}
}

// NewPutOrganizationProjectApisBuildFoldersParamsWithHTTPClient creates a new PutOrganizationProjectApisBuildFoldersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOrganizationProjectApisBuildFoldersParamsWithHTTPClient(client *http.Client) *PutOrganizationProjectApisBuildFoldersParams {
	var ()
	return &PutOrganizationProjectApisBuildFoldersParams{
		HTTPClient: client,
	}
}

/*PutOrganizationProjectApisBuildFoldersParams contains all the parameters to send to the API endpoint
for the put organization project apis build folders operation typically these are written to a http.Request
*/
type PutOrganizationProjectApisBuildFoldersParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  The folder.

	*/
	Body *models.Folder
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Path
	  The full path of the folder.

	*/
	Path string
	/*Project
	  Project ID or project name

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) WithTimeout(timeout time.Duration) *PutOrganizationProjectApisBuildFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) WithContext(ctx context.Context) *PutOrganizationProjectApisBuildFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) WithHTTPClient(client *http.Client) *PutOrganizationProjectApisBuildFoldersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) WithAPIVersion(aPIVersion string) *PutOrganizationProjectApisBuildFoldersParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) WithBody(body *models.Folder) *PutOrganizationProjectApisBuildFoldersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) SetBody(body *models.Folder) {
	o.Body = body
}

// WithOrganization adds the organization to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) WithOrganization(organization string) *PutOrganizationProjectApisBuildFoldersParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPath adds the path to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) WithPath(path string) *PutOrganizationProjectApisBuildFoldersParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) SetPath(path string) {
	o.Path = path
}

// WithProject adds the project to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) WithProject(project string) *PutOrganizationProjectApisBuildFoldersParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the put organization project apis build folders params
func (o *PutOrganizationProjectApisBuildFoldersParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrganizationProjectApisBuildFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {
		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
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
