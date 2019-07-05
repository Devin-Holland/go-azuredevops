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

// NewUpdateBoardCardSettingsParams creates a new UpdateBoardCardSettingsParams object
// with the default values initialized.
func NewUpdateBoardCardSettingsParams() *UpdateBoardCardSettingsParams {
	var ()
	return &UpdateBoardCardSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBoardCardSettingsParamsWithTimeout creates a new UpdateBoardCardSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateBoardCardSettingsParamsWithTimeout(timeout time.Duration) *UpdateBoardCardSettingsParams {
	var ()
	return &UpdateBoardCardSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateBoardCardSettingsParamsWithContext creates a new UpdateBoardCardSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateBoardCardSettingsParamsWithContext(ctx context.Context) *UpdateBoardCardSettingsParams {
	var ()
	return &UpdateBoardCardSettingsParams{

		Context: ctx,
	}
}

// NewUpdateBoardCardSettingsParamsWithHTTPClient creates a new UpdateBoardCardSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateBoardCardSettingsParamsWithHTTPClient(client *http.Client) *UpdateBoardCardSettingsParams {
	var ()
	return &UpdateBoardCardSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateBoardCardSettingsParams contains all the parameters to send to the API endpoint
for the update board card settings operation typically these are written to a http.Request
*/
type UpdateBoardCardSettingsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Board*/
	Board string
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

// WithTimeout adds the timeout to the update board card settings params
func (o *UpdateBoardCardSettingsParams) WithTimeout(timeout time.Duration) *UpdateBoardCardSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update board card settings params
func (o *UpdateBoardCardSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update board card settings params
func (o *UpdateBoardCardSettingsParams) WithContext(ctx context.Context) *UpdateBoardCardSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update board card settings params
func (o *UpdateBoardCardSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update board card settings params
func (o *UpdateBoardCardSettingsParams) WithHTTPClient(client *http.Client) *UpdateBoardCardSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update board card settings params
func (o *UpdateBoardCardSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update board card settings params
func (o *UpdateBoardCardSettingsParams) WithAPIVersion(aPIVersion string) *UpdateBoardCardSettingsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update board card settings params
func (o *UpdateBoardCardSettingsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBoard adds the board to the update board card settings params
func (o *UpdateBoardCardSettingsParams) WithBoard(board string) *UpdateBoardCardSettingsParams {
	o.SetBoard(board)
	return o
}

// SetBoard adds the board to the update board card settings params
func (o *UpdateBoardCardSettingsParams) SetBoard(board string) {
	o.Board = board
}

// WithBody adds the body to the update board card settings params
func (o *UpdateBoardCardSettingsParams) WithBody(body *models.BoardCardSettings) *UpdateBoardCardSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update board card settings params
func (o *UpdateBoardCardSettingsParams) SetBody(body *models.BoardCardSettings) {
	o.Body = body
}

// WithOrganization adds the organization to the update board card settings params
func (o *UpdateBoardCardSettingsParams) WithOrganization(organization string) *UpdateBoardCardSettingsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update board card settings params
func (o *UpdateBoardCardSettingsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the update board card settings params
func (o *UpdateBoardCardSettingsParams) WithProject(project string) *UpdateBoardCardSettingsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update board card settings params
func (o *UpdateBoardCardSettingsParams) SetProject(project string) {
	o.Project = project
}

// WithTeam adds the team to the update board card settings params
func (o *UpdateBoardCardSettingsParams) WithTeam(team string) *UpdateBoardCardSettingsParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the update board card settings params
func (o *UpdateBoardCardSettingsParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBoardCardSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param board
	if err := r.SetPathParam("board", o.Board); err != nil {
		return err
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
