// Code generated by go-swagger; DO NOT EDIT.

package cardsettings

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

// NewUpdateTaskboardCardSettingsParams creates a new UpdateTaskboardCardSettingsParams object
// with the default values initialized.
func NewUpdateTaskboardCardSettingsParams() *UpdateTaskboardCardSettingsParams {
	var ()
	return &UpdateTaskboardCardSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTaskboardCardSettingsParamsWithTimeout creates a new UpdateTaskboardCardSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTaskboardCardSettingsParamsWithTimeout(timeout time.Duration) *UpdateTaskboardCardSettingsParams {
	var ()
	return &UpdateTaskboardCardSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateTaskboardCardSettingsParamsWithContext creates a new UpdateTaskboardCardSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTaskboardCardSettingsParamsWithContext(ctx context.Context) *UpdateTaskboardCardSettingsParams {
	var ()
	return &UpdateTaskboardCardSettingsParams{

		Context: ctx,
	}
}

// NewUpdateTaskboardCardSettingsParamsWithHTTPClient creates a new UpdateTaskboardCardSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTaskboardCardSettingsParamsWithHTTPClient(client *http.Client) *UpdateTaskboardCardSettingsParams {
	var ()
	return &UpdateTaskboardCardSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateTaskboardCardSettingsParams contains all the parameters to send to the API endpoint
for the update taskboard card settings operation typically these are written to a http.Request
*/
type UpdateTaskboardCardSettingsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Body*/
	Body *models.BoardCardSettings
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

// WithTimeout adds the timeout to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) WithTimeout(timeout time.Duration) *UpdateTaskboardCardSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) WithContext(ctx context.Context) *UpdateTaskboardCardSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) WithHTTPClient(client *http.Client) *UpdateTaskboardCardSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) WithAPIVersion(aPIVersion string) *UpdateTaskboardCardSettingsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) WithBody(body *models.BoardCardSettings) *UpdateTaskboardCardSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) SetBody(body *models.BoardCardSettings) {
	o.Body = body
}

// WithOrganization adds the organization to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) WithOrganization(organization string) *UpdateTaskboardCardSettingsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) WithProject(project string) *UpdateTaskboardCardSettingsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) SetProject(project string) {
	o.Project = project
}

// WithTeam adds the team to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) WithTeam(team string) *UpdateTaskboardCardSettingsParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the update taskboard card settings params
func (o *UpdateTaskboardCardSettingsParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTaskboardCardSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}