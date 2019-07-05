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

// NewUpdateWidgetParams creates a new UpdateWidgetParams object
// with the default values initialized.
func NewUpdateWidgetParams() *UpdateWidgetParams {
	var ()
	return &UpdateWidgetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWidgetParamsWithTimeout creates a new UpdateWidgetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateWidgetParamsWithTimeout(timeout time.Duration) *UpdateWidgetParams {
	var ()
	return &UpdateWidgetParams{

		timeout: timeout,
	}
}

// NewUpdateWidgetParamsWithContext creates a new UpdateWidgetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateWidgetParamsWithContext(ctx context.Context) *UpdateWidgetParams {
	var ()
	return &UpdateWidgetParams{

		Context: ctx,
	}
}

// NewUpdateWidgetParamsWithHTTPClient creates a new UpdateWidgetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateWidgetParamsWithHTTPClient(client *http.Client) *UpdateWidgetParams {
	var ()
	return &UpdateWidgetParams{
		HTTPClient: client,
	}
}

/*UpdateWidgetParams contains all the parameters to send to the API endpoint
for the update widget operation typically these are written to a http.Request
*/
type UpdateWidgetParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  Description of the widget changes to apply. All non-null fields will be replaced.

	*/
	Body *models.Widget
	/*DashboardID
	  ID of the dashboard containing the widget.

	*/
	DashboardID strfmt.UUID
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
	/*WidgetID
	  ID of the widget to update.

	*/
	WidgetID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update widget params
func (o *UpdateWidgetParams) WithTimeout(timeout time.Duration) *UpdateWidgetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update widget params
func (o *UpdateWidgetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update widget params
func (o *UpdateWidgetParams) WithContext(ctx context.Context) *UpdateWidgetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update widget params
func (o *UpdateWidgetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update widget params
func (o *UpdateWidgetParams) WithHTTPClient(client *http.Client) *UpdateWidgetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update widget params
func (o *UpdateWidgetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update widget params
func (o *UpdateWidgetParams) WithAPIVersion(aPIVersion string) *UpdateWidgetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update widget params
func (o *UpdateWidgetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update widget params
func (o *UpdateWidgetParams) WithBody(body *models.Widget) *UpdateWidgetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update widget params
func (o *UpdateWidgetParams) SetBody(body *models.Widget) {
	o.Body = body
}

// WithDashboardID adds the dashboardID to the update widget params
func (o *UpdateWidgetParams) WithDashboardID(dashboardID strfmt.UUID) *UpdateWidgetParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the update widget params
func (o *UpdateWidgetParams) SetDashboardID(dashboardID strfmt.UUID) {
	o.DashboardID = dashboardID
}

// WithOrganization adds the organization to the update widget params
func (o *UpdateWidgetParams) WithOrganization(organization string) *UpdateWidgetParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update widget params
func (o *UpdateWidgetParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the update widget params
func (o *UpdateWidgetParams) WithProject(project string) *UpdateWidgetParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update widget params
func (o *UpdateWidgetParams) SetProject(project string) {
	o.Project = project
}

// WithTeam adds the team to the update widget params
func (o *UpdateWidgetParams) WithTeam(team string) *UpdateWidgetParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the update widget params
func (o *UpdateWidgetParams) SetTeam(team string) {
	o.Team = team
}

// WithWidgetID adds the widgetID to the update widget params
func (o *UpdateWidgetParams) WithWidgetID(widgetID strfmt.UUID) *UpdateWidgetParams {
	o.SetWidgetID(widgetID)
	return o
}

// SetWidgetID adds the widgetId to the update widget params
func (o *UpdateWidgetParams) SetWidgetID(widgetID strfmt.UUID) {
	o.WidgetID = widgetID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWidgetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param widgetId
	if err := r.SetPathParam("widgetId", o.WidgetID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
