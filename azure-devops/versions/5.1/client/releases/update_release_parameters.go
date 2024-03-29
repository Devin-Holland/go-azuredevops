// Code generated by go-swagger; DO NOT EDIT.

package releases

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

// NewUpdateReleaseParams creates a new UpdateReleaseParams object
// with the default values initialized.
func NewUpdateReleaseParams() *UpdateReleaseParams {
	var ()
	return &UpdateReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReleaseParamsWithTimeout creates a new UpdateReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateReleaseParamsWithTimeout(timeout time.Duration) *UpdateReleaseParams {
	var ()
	return &UpdateReleaseParams{

		timeout: timeout,
	}
}

// NewUpdateReleaseParamsWithContext creates a new UpdateReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateReleaseParamsWithContext(ctx context.Context) *UpdateReleaseParams {
	var ()
	return &UpdateReleaseParams{

		Context: ctx,
	}
}

// NewUpdateReleaseParamsWithHTTPClient creates a new UpdateReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateReleaseParamsWithHTTPClient(client *http.Client) *UpdateReleaseParams {
	var ()
	return &UpdateReleaseParams{
		HTTPClient: client,
	}
}

/*UpdateReleaseParams contains all the parameters to send to the API endpoint
for the update release operation typically these are written to a http.Request
*/
type UpdateReleaseParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.8' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  Release object for update.

	*/
	Body *models.Release
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*ReleaseID
	  Id of the release to update.

	*/
	ReleaseID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update release params
func (o *UpdateReleaseParams) WithTimeout(timeout time.Duration) *UpdateReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update release params
func (o *UpdateReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update release params
func (o *UpdateReleaseParams) WithContext(ctx context.Context) *UpdateReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update release params
func (o *UpdateReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update release params
func (o *UpdateReleaseParams) WithHTTPClient(client *http.Client) *UpdateReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update release params
func (o *UpdateReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update release params
func (o *UpdateReleaseParams) WithAPIVersion(aPIVersion string) *UpdateReleaseParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update release params
func (o *UpdateReleaseParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update release params
func (o *UpdateReleaseParams) WithBody(body *models.Release) *UpdateReleaseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update release params
func (o *UpdateReleaseParams) SetBody(body *models.Release) {
	o.Body = body
}

// WithOrganization adds the organization to the update release params
func (o *UpdateReleaseParams) WithOrganization(organization string) *UpdateReleaseParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update release params
func (o *UpdateReleaseParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the update release params
func (o *UpdateReleaseParams) WithProject(project string) *UpdateReleaseParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update release params
func (o *UpdateReleaseParams) SetProject(project string) {
	o.Project = project
}

// WithReleaseID adds the releaseID to the update release params
func (o *UpdateReleaseParams) WithReleaseID(releaseID int32) *UpdateReleaseParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the update release params
func (o *UpdateReleaseParams) SetReleaseID(releaseID int32) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param releaseId
	if err := r.SetPathParam("releaseId", swag.FormatInt32(o.ReleaseID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
