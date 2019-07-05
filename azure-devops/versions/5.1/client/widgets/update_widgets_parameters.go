// Code generated by go-swagger; DO NOT EDIT.

package widgets

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

// NewUpdateWidgetsParams creates a new UpdateWidgetsParams object
// with the default values initialized.
func NewUpdateWidgetsParams() *UpdateWidgetsParams {
	var ()
	return &UpdateWidgetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWidgetsParamsWithTimeout creates a new UpdateWidgetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateWidgetsParamsWithTimeout(timeout time.Duration) *UpdateWidgetsParams {
	var ()
	return &UpdateWidgetsParams{

		timeout: timeout,
	}
}

// NewUpdateWidgetsParamsWithContext creates a new UpdateWidgetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateWidgetsParamsWithContext(ctx context.Context) *UpdateWidgetsParams {
	var ()
	return &UpdateWidgetsParams{

		Context: ctx,
	}
}

// NewUpdateWidgetsParamsWithHTTPClient creates a new UpdateWidgetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateWidgetsParamsWithHTTPClient(client *http.Client) *UpdateWidgetsParams {
	var ()
	return &UpdateWidgetsParams{
		HTTPClient: client,
	}
}

/*UpdateWidgetsParams contains all the parameters to send to the API endpoint
for the update widgets operation typically these are written to a http.Request
*/
type UpdateWidgetsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  The set of widget states to update on the dashboard.

	*/
	Body []*models.Widget
	/*DashboardID
	  ID of the Dashboard to modify.

	*/
	DashboardID strfmt.UUID
	/*ETag
	  Dashboard Widgets Version

	*/
	ETag *string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*Team
	  Team ID or team name

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update widgets params
func (o *UpdateWidgetsParams) WithTimeout(timeout time.Duration) *UpdateWidgetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update widgets params
func (o *UpdateWidgetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update widgets params
func (o *UpdateWidgetsParams) WithContext(ctx context.Context) *UpdateWidgetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update widgets params
func (o *UpdateWidgetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update widgets params
func (o *UpdateWidgetsParams) WithHTTPClient(client *http.Client) *UpdateWidgetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update widgets params
func (o *UpdateWidgetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update widgets params
func (o *UpdateWidgetsParams) WithAPIVersion(aPIVersion string) *UpdateWidgetsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update widgets params
func (o *UpdateWidgetsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update widgets params
func (o *UpdateWidgetsParams) WithBody(body []*models.Widget) *UpdateWidgetsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update widgets params
func (o *UpdateWidgetsParams) SetBody(body []*models.Widget) {
	o.Body = body
}

// WithDashboardID adds the dashboardID to the update widgets params
func (o *UpdateWidgetsParams) WithDashboardID(dashboardID strfmt.UUID) *UpdateWidgetsParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the update widgets params
func (o *UpdateWidgetsParams) SetDashboardID(dashboardID strfmt.UUID) {
	o.DashboardID = dashboardID
}

// WithETag adds the eTag to the update widgets params
func (o *UpdateWidgetsParams) WithETag(eTag *string) *UpdateWidgetsParams {
	o.SetETag(eTag)
	return o
}

// SetETag adds the eTag to the update widgets params
func (o *UpdateWidgetsParams) SetETag(eTag *string) {
	o.ETag = eTag
}

// WithOrganization adds the organization to the update widgets params
func (o *UpdateWidgetsParams) WithOrganization(organization string) *UpdateWidgetsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update widgets params
func (o *UpdateWidgetsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the update widgets params
func (o *UpdateWidgetsParams) WithProject(project string) *UpdateWidgetsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update widgets params
func (o *UpdateWidgetsParams) SetProject(project string) {
	o.Project = project
}

// WithTeam adds the team to the update widgets params
func (o *UpdateWidgetsParams) WithTeam(team string) *UpdateWidgetsParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the update widgets params
func (o *UpdateWidgetsParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWidgetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param dashboardId
	if err := r.SetPathParam("dashboardId", o.DashboardID.String()); err != nil {
		return err
	}

	if o.ETag != nil {

		// header param eTag
		if err := r.SetHeaderParam("eTag", *o.ETag); err != nil {
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

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
