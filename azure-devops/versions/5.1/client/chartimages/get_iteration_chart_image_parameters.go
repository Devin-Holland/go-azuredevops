// Code generated by go-swagger; DO NOT EDIT.

package chartimages

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

// NewGetIterationChartImageParams creates a new GetIterationChartImageParams object
// with the default values initialized.
func NewGetIterationChartImageParams() *GetIterationChartImageParams {
	var ()
	return &GetIterationChartImageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIterationChartImageParamsWithTimeout creates a new GetIterationChartImageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIterationChartImageParamsWithTimeout(timeout time.Duration) *GetIterationChartImageParams {
	var ()
	return &GetIterationChartImageParams{

		timeout: timeout,
	}
}

// NewGetIterationChartImageParamsWithContext creates a new GetIterationChartImageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIterationChartImageParamsWithContext(ctx context.Context) *GetIterationChartImageParams {
	var ()
	return &GetIterationChartImageParams{

		Context: ctx,
	}
}

// NewGetIterationChartImageParamsWithHTTPClient creates a new GetIterationChartImageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIterationChartImageParamsWithHTTPClient(client *http.Client) *GetIterationChartImageParams {
	var ()
	return &GetIterationChartImageParams{
		HTTPClient: client,
	}
}

/*GetIterationChartImageParams contains all the parameters to send to the API endpoint
for the get iteration chart image operation typically these are written to a http.Request
*/
type GetIterationChartImageParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Height
	  The height of the chart in pixels. Must be greater than 0.

	*/
	Height *int32
	/*IterationID
	  ID of the iteration.

	*/
	IterationID strfmt.UUID
	/*Name
	  The chart name. e.g. Burndown.

	*/
	Name string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*ShowDetails
	  Whether or not the chart should include detailed information (e.g. axis labels, titles, trend lines, etc.)

	*/
	ShowDetails *bool
	/*Team
	  Team ID or team name

	*/
	Team string
	/*Title
	  The title of the chart. Can only be dislayed if ShowLabels is true.

	*/
	Title *string
	/*Width
	  The width of the chart in pixels. Must be greater than 0.

	*/
	Width *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iteration chart image params
func (o *GetIterationChartImageParams) WithTimeout(timeout time.Duration) *GetIterationChartImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iteration chart image params
func (o *GetIterationChartImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iteration chart image params
func (o *GetIterationChartImageParams) WithContext(ctx context.Context) *GetIterationChartImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iteration chart image params
func (o *GetIterationChartImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iteration chart image params
func (o *GetIterationChartImageParams) WithHTTPClient(client *http.Client) *GetIterationChartImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iteration chart image params
func (o *GetIterationChartImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get iteration chart image params
func (o *GetIterationChartImageParams) WithAPIVersion(aPIVersion string) *GetIterationChartImageParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get iteration chart image params
func (o *GetIterationChartImageParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithHeight adds the height to the get iteration chart image params
func (o *GetIterationChartImageParams) WithHeight(height *int32) *GetIterationChartImageParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the get iteration chart image params
func (o *GetIterationChartImageParams) SetHeight(height *int32) {
	o.Height = height
}

// WithIterationID adds the iterationID to the get iteration chart image params
func (o *GetIterationChartImageParams) WithIterationID(iterationID strfmt.UUID) *GetIterationChartImageParams {
	o.SetIterationID(iterationID)
	return o
}

// SetIterationID adds the iterationId to the get iteration chart image params
func (o *GetIterationChartImageParams) SetIterationID(iterationID strfmt.UUID) {
	o.IterationID = iterationID
}

// WithName adds the name to the get iteration chart image params
func (o *GetIterationChartImageParams) WithName(name string) *GetIterationChartImageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get iteration chart image params
func (o *GetIterationChartImageParams) SetName(name string) {
	o.Name = name
}

// WithOrganization adds the organization to the get iteration chart image params
func (o *GetIterationChartImageParams) WithOrganization(organization string) *GetIterationChartImageParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get iteration chart image params
func (o *GetIterationChartImageParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get iteration chart image params
func (o *GetIterationChartImageParams) WithProject(project string) *GetIterationChartImageParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get iteration chart image params
func (o *GetIterationChartImageParams) SetProject(project string) {
	o.Project = project
}

// WithShowDetails adds the showDetails to the get iteration chart image params
func (o *GetIterationChartImageParams) WithShowDetails(showDetails *bool) *GetIterationChartImageParams {
	o.SetShowDetails(showDetails)
	return o
}

// SetShowDetails adds the showDetails to the get iteration chart image params
func (o *GetIterationChartImageParams) SetShowDetails(showDetails *bool) {
	o.ShowDetails = showDetails
}

// WithTeam adds the team to the get iteration chart image params
func (o *GetIterationChartImageParams) WithTeam(team string) *GetIterationChartImageParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the get iteration chart image params
func (o *GetIterationChartImageParams) SetTeam(team string) {
	o.Team = team
}

// WithTitle adds the title to the get iteration chart image params
func (o *GetIterationChartImageParams) WithTitle(title *string) *GetIterationChartImageParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the get iteration chart image params
func (o *GetIterationChartImageParams) SetTitle(title *string) {
	o.Title = title
}

// WithWidth adds the width to the get iteration chart image params
func (o *GetIterationChartImageParams) WithWidth(width *int32) *GetIterationChartImageParams {
	o.SetWidth(width)
	return o
}

// SetWidth adds the width to the get iteration chart image params
func (o *GetIterationChartImageParams) SetWidth(width *int32) {
	o.Width = width
}

// WriteToRequest writes these params to a swagger request
func (o *GetIterationChartImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Height != nil {

		// query param height
		var qrHeight int32
		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := swag.FormatInt32(qrHeight)
		if qHeight != "" {
			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}

	}

	// path param iterationId
	if err := r.SetPathParam("iterationId", o.IterationID.String()); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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

	if o.ShowDetails != nil {

		// query param showDetails
		var qrShowDetails bool
		if o.ShowDetails != nil {
			qrShowDetails = *o.ShowDetails
		}
		qShowDetails := swag.FormatBool(qrShowDetails)
		if qShowDetails != "" {
			if err := r.SetQueryParam("showDetails", qShowDetails); err != nil {
				return err
			}
		}

	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if o.Title != nil {

		// query param title
		var qrTitle string
		if o.Title != nil {
			qrTitle = *o.Title
		}
		qTitle := qrTitle
		if qTitle != "" {
			if err := r.SetQueryParam("title", qTitle); err != nil {
				return err
			}
		}

	}

	if o.Width != nil {

		// query param width
		var qrWidth int32
		if o.Width != nil {
			qrWidth = *o.Width
		}
		qWidth := swag.FormatInt32(qrWidth)
		if qWidth != "" {
			if err := r.SetQueryParam("width", qWidth); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
