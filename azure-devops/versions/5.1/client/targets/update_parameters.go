// Code generated by go-swagger; DO NOT EDIT.

package targets

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

// NewUpdateParams creates a new UpdateParams object
// with the default values initialized.
func NewUpdateParams() *UpdateParams {
	var ()
	return &UpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateParamsWithTimeout creates a new UpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateParamsWithTimeout(timeout time.Duration) *UpdateParams {
	var ()
	return &UpdateParams{

		timeout: timeout,
	}
}

// NewUpdateParamsWithContext creates a new UpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateParamsWithContext(ctx context.Context) *UpdateParams {
	var ()
	return &UpdateParams{

		Context: ctx,
	}
}

// NewUpdateParamsWithHTTPClient creates a new UpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateParamsWithHTTPClient(client *http.Client) *UpdateParams {
	var ()
	return &UpdateParams{
		HTTPClient: client,
	}
}

/*UpdateParams contains all the parameters to send to the API endpoint
for the update operation typically these are written to a http.Request
*/
type UpdateParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  Deployment targets with tags to udpdate.

	*/
	Body []*models.DeploymentTargetUpdateParameter
	/*DeploymentGroupID
	  ID of the deployment group in which deployment targets are updated.

	*/
	DeploymentGroupID int32
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

// WithTimeout adds the timeout to the update params
func (o *UpdateParams) WithTimeout(timeout time.Duration) *UpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update params
func (o *UpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update params
func (o *UpdateParams) WithContext(ctx context.Context) *UpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update params
func (o *UpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) WithHTTPClient(client *http.Client) *UpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update params
func (o *UpdateParams) WithAPIVersion(aPIVersion string) *UpdateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update params
func (o *UpdateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update params
func (o *UpdateParams) WithBody(body []*models.DeploymentTargetUpdateParameter) *UpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update params
func (o *UpdateParams) SetBody(body []*models.DeploymentTargetUpdateParameter) {
	o.Body = body
}

// WithDeploymentGroupID adds the deploymentGroupID to the update params
func (o *UpdateParams) WithDeploymentGroupID(deploymentGroupID int32) *UpdateParams {
	o.SetDeploymentGroupID(deploymentGroupID)
	return o
}

// SetDeploymentGroupID adds the deploymentGroupId to the update params
func (o *UpdateParams) SetDeploymentGroupID(deploymentGroupID int32) {
	o.DeploymentGroupID = deploymentGroupID
}

// WithOrganization adds the organization to the update params
func (o *UpdateParams) WithOrganization(organization string) *UpdateParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update params
func (o *UpdateParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the update params
func (o *UpdateParams) WithProject(project string) *UpdateParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update params
func (o *UpdateParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deploymentGroupId
	if err := r.SetPathParam("deploymentGroupId", swag.FormatInt32(o.DeploymentGroupID)); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
