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
)

// NewGetWidgetParams creates a new GetWidgetParams object
// with the default values initialized.
func NewGetWidgetParams() *GetWidgetParams {
	var ()
	return &GetWidgetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWidgetParamsWithTimeout creates a new GetWidgetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWidgetParamsWithTimeout(timeout time.Duration) *GetWidgetParams {
	var ()
	return &GetWidgetParams{

		timeout: timeout,
	}
}

// NewGetWidgetParamsWithContext creates a new GetWidgetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWidgetParamsWithContext(ctx context.Context) *GetWidgetParams {
	var ()
	return &GetWidgetParams{

		Context: ctx,
	}
}

// NewGetWidgetParamsWithHTTPClient creates a new GetWidgetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWidgetParamsWithHTTPClient(client *http.Client) *GetWidgetParams {
	var ()
	return &GetWidgetParams{
		HTTPClient: client,
	}
}

/*GetWidgetParams contains all the parameters to send to the API endpoint
for the get widget operation typically these are written to a http.Request
*/
type GetWidgetParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
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
	  ID of the widget to read.

	*/
	WidgetID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get widget params
func (o *GetWidgetParams) WithTimeout(timeout time.Duration) *GetWidgetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get widget params
func (o *GetWidgetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get widget params
func (o *GetWidgetParams) WithContext(ctx context.Context) *GetWidgetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get widget params
func (o *GetWidgetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get widget params
func (o *GetWidgetParams) WithHTTPClient(client *http.Client) *GetWidgetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get widget params
func (o *GetWidgetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get widget params
func (o *GetWidgetParams) WithAPIVersion(aPIVersion string) *GetWidgetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get widget params
func (o *GetWidgetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDashboardID adds the dashboardID to the get widget params
func (o *GetWidgetParams) WithDashboardID(dashboardID strfmt.UUID) *GetWidgetParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the get widget params
func (o *GetWidgetParams) SetDashboardID(dashboardID strfmt.UUID) {
	o.DashboardID = dashboardID
}

// WithOrganization adds the organization to the get widget params
func (o *GetWidgetParams) WithOrganization(organization string) *GetWidgetParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get widget params
func (o *GetWidgetParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get widget params
func (o *GetWidgetParams) WithProject(project string) *GetWidgetParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get widget params
func (o *GetWidgetParams) SetProject(project string) {
	o.Project = project
}

// WithTeam adds the team to the get widget params
func (o *GetWidgetParams) WithTeam(team string) *GetWidgetParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the get widget params
func (o *GetWidgetParams) SetTeam(team string) {
	o.Team = team
}

// WithWidgetID adds the widgetID to the get widget params
func (o *GetWidgetParams) WithWidgetID(widgetID strfmt.UUID) *GetWidgetParams {
	o.SetWidgetID(widgetID)
	return o
}

// SetWidgetID adds the widgetId to the get widget params
func (o *GetWidgetParams) SetWidgetID(widgetID strfmt.UUID) {
	o.WidgetID = widgetID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWidgetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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