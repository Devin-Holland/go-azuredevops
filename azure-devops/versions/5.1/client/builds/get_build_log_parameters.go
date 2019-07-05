// Code generated by go-swagger; DO NOT EDIT.

package builds

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

// NewGetBuildLogParams creates a new GetBuildLogParams object
// with the default values initialized.
func NewGetBuildLogParams() *GetBuildLogParams {
	var ()
	return &GetBuildLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildLogParamsWithTimeout creates a new GetBuildLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBuildLogParamsWithTimeout(timeout time.Duration) *GetBuildLogParams {
	var ()
	return &GetBuildLogParams{

		timeout: timeout,
	}
}

// NewGetBuildLogParamsWithContext creates a new GetBuildLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBuildLogParamsWithContext(ctx context.Context) *GetBuildLogParams {
	var ()
	return &GetBuildLogParams{

		Context: ctx,
	}
}

// NewGetBuildLogParamsWithHTTPClient creates a new GetBuildLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBuildLogParamsWithHTTPClient(client *http.Client) *GetBuildLogParams {
	var ()
	return &GetBuildLogParams{
		HTTPClient: client,
	}
}

/*GetBuildLogParams contains all the parameters to send to the API endpoint
for the get build log operation typically these are written to a http.Request
*/
type GetBuildLogParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*BuildID
	  The ID of the build.

	*/
	BuildID int32
	/*EndLine
	  The end line.

	*/
	EndLine *int64
	/*LogID
	  The ID of the log file.

	*/
	LogID int32
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*StartLine
	  The start line.

	*/
	StartLine *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get build log params
func (o *GetBuildLogParams) WithTimeout(timeout time.Duration) *GetBuildLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build log params
func (o *GetBuildLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build log params
func (o *GetBuildLogParams) WithContext(ctx context.Context) *GetBuildLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build log params
func (o *GetBuildLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build log params
func (o *GetBuildLogParams) WithHTTPClient(client *http.Client) *GetBuildLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build log params
func (o *GetBuildLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get build log params
func (o *GetBuildLogParams) WithAPIVersion(aPIVersion string) *GetBuildLogParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get build log params
func (o *GetBuildLogParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBuildID adds the buildID to the get build log params
func (o *GetBuildLogParams) WithBuildID(buildID int32) *GetBuildLogParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the get build log params
func (o *GetBuildLogParams) SetBuildID(buildID int32) {
	o.BuildID = buildID
}

// WithEndLine adds the endLine to the get build log params
func (o *GetBuildLogParams) WithEndLine(endLine *int64) *GetBuildLogParams {
	o.SetEndLine(endLine)
	return o
}

// SetEndLine adds the endLine to the get build log params
func (o *GetBuildLogParams) SetEndLine(endLine *int64) {
	o.EndLine = endLine
}

// WithLogID adds the logID to the get build log params
func (o *GetBuildLogParams) WithLogID(logID int32) *GetBuildLogParams {
	o.SetLogID(logID)
	return o
}

// SetLogID adds the logId to the get build log params
func (o *GetBuildLogParams) SetLogID(logID int32) {
	o.LogID = logID
}

// WithOrganization adds the organization to the get build log params
func (o *GetBuildLogParams) WithOrganization(organization string) *GetBuildLogParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get build log params
func (o *GetBuildLogParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get build log params
func (o *GetBuildLogParams) WithProject(project string) *GetBuildLogParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get build log params
func (o *GetBuildLogParams) SetProject(project string) {
	o.Project = project
}

// WithStartLine adds the startLine to the get build log params
func (o *GetBuildLogParams) WithStartLine(startLine *int64) *GetBuildLogParams {
	o.SetStartLine(startLine)
	return o
}

// SetStartLine adds the startLine to the get build log params
func (o *GetBuildLogParams) SetStartLine(startLine *int64) {
	o.StartLine = startLine
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param buildId
	if err := r.SetPathParam("buildId", swag.FormatInt32(o.BuildID)); err != nil {
		return err
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

	// path param logId
	if err := r.SetPathParam("logId", swag.FormatInt32(o.LogID)); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
