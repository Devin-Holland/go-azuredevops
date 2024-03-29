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
)

// NewGetTaskLogParams creates a new GetTaskLogParams object
// with the default values initialized.
func NewGetTaskLogParams() *GetTaskLogParams {
	var ()
	return &GetTaskLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskLogParamsWithTimeout creates a new GetTaskLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTaskLogParamsWithTimeout(timeout time.Duration) *GetTaskLogParams {
	var ()
	return &GetTaskLogParams{

		timeout: timeout,
	}
}

// NewGetTaskLogParamsWithContext creates a new GetTaskLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTaskLogParamsWithContext(ctx context.Context) *GetTaskLogParams {
	var ()
	return &GetTaskLogParams{

		Context: ctx,
	}
}

// NewGetTaskLogParamsWithHTTPClient creates a new GetTaskLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTaskLogParamsWithHTTPClient(client *http.Client) *GetTaskLogParams {
	var ()
	return &GetTaskLogParams{
		HTTPClient: client,
	}
}

/*GetTaskLogParams contains all the parameters to send to the API endpoint
for the get task log operation typically these are written to a http.Request
*/
type GetTaskLogParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*EndLine
	  Ending line number for logs

	*/
	EndLine *int64
	/*EnvironmentID
	  Id of release environment.

	*/
	EnvironmentID int32
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*ReleaseDeployPhaseID
	  Release deploy phase Id.

	*/
	ReleaseDeployPhaseID int32
	/*ReleaseID
	  Id of the release.

	*/
	ReleaseID int32
	/*StartLine
	  Starting line number for logs

	*/
	StartLine *int64
	/*TaskID
	  ReleaseTask Id for the log.

	*/
	TaskID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get task log params
func (o *GetTaskLogParams) WithTimeout(timeout time.Duration) *GetTaskLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task log params
func (o *GetTaskLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task log params
func (o *GetTaskLogParams) WithContext(ctx context.Context) *GetTaskLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task log params
func (o *GetTaskLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task log params
func (o *GetTaskLogParams) WithHTTPClient(client *http.Client) *GetTaskLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task log params
func (o *GetTaskLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get task log params
func (o *GetTaskLogParams) WithAPIVersion(aPIVersion string) *GetTaskLogParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get task log params
func (o *GetTaskLogParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithEndLine adds the endLine to the get task log params
func (o *GetTaskLogParams) WithEndLine(endLine *int64) *GetTaskLogParams {
	o.SetEndLine(endLine)
	return o
}

// SetEndLine adds the endLine to the get task log params
func (o *GetTaskLogParams) SetEndLine(endLine *int64) {
	o.EndLine = endLine
}

// WithEnvironmentID adds the environmentID to the get task log params
func (o *GetTaskLogParams) WithEnvironmentID(environmentID int32) *GetTaskLogParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the get task log params
func (o *GetTaskLogParams) SetEnvironmentID(environmentID int32) {
	o.EnvironmentID = environmentID
}

// WithOrganization adds the organization to the get task log params
func (o *GetTaskLogParams) WithOrganization(organization string) *GetTaskLogParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get task log params
func (o *GetTaskLogParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get task log params
func (o *GetTaskLogParams) WithProject(project string) *GetTaskLogParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get task log params
func (o *GetTaskLogParams) SetProject(project string) {
	o.Project = project
}

// WithReleaseDeployPhaseID adds the releaseDeployPhaseID to the get task log params
func (o *GetTaskLogParams) WithReleaseDeployPhaseID(releaseDeployPhaseID int32) *GetTaskLogParams {
	o.SetReleaseDeployPhaseID(releaseDeployPhaseID)
	return o
}

// SetReleaseDeployPhaseID adds the releaseDeployPhaseId to the get task log params
func (o *GetTaskLogParams) SetReleaseDeployPhaseID(releaseDeployPhaseID int32) {
	o.ReleaseDeployPhaseID = releaseDeployPhaseID
}

// WithReleaseID adds the releaseID to the get task log params
func (o *GetTaskLogParams) WithReleaseID(releaseID int32) *GetTaskLogParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the get task log params
func (o *GetTaskLogParams) SetReleaseID(releaseID int32) {
	o.ReleaseID = releaseID
}

// WithStartLine adds the startLine to the get task log params
func (o *GetTaskLogParams) WithStartLine(startLine *int64) *GetTaskLogParams {
	o.SetStartLine(startLine)
	return o
}

// SetStartLine adds the startLine to the get task log params
func (o *GetTaskLogParams) SetStartLine(startLine *int64) {
	o.StartLine = startLine
}

// WithTaskID adds the taskID to the get task log params
func (o *GetTaskLogParams) WithTaskID(taskID int32) *GetTaskLogParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get task log params
func (o *GetTaskLogParams) SetTaskID(taskID int32) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EndLine != nil {

		// query param endLine
		var qrEndLine int64
		if o.EndLine != nil {
			qrEndLine = *o.EndLine
		}
		qEndLine := swag.FormatInt64(qrEndLine)
		if qEndLine != "" {
			if err := r.SetQueryParam("endLine", qEndLine); err != nil {
				return err
			}
		}

	}

	// path param environmentId
	if err := r.SetPathParam("environmentId", swag.FormatInt32(o.EnvironmentID)); err != nil {
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

	// path param releaseDeployPhaseId
	if err := r.SetPathParam("releaseDeployPhaseId", swag.FormatInt32(o.ReleaseDeployPhaseID)); err != nil {
		return err
	}

	// path param releaseId
	if err := r.SetPathParam("releaseId", swag.FormatInt32(o.ReleaseID)); err != nil {
		return err
	}

	if o.StartLine != nil {

		// query param startLine
		var qrStartLine int64
		if o.StartLine != nil {
			qrStartLine = *o.StartLine
		}
		qStartLine := swag.FormatInt64(qrStartLine)
		if qStartLine != "" {
			if err := r.SetQueryParam("startLine", qStartLine); err != nil {
				return err
			}
		}

	}

	// path param taskId
	if err := r.SetPathParam("taskId", swag.FormatInt32(o.TaskID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
