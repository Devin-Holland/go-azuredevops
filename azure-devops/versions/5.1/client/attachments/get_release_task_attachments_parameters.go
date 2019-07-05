// Code generated by go-swagger; DO NOT EDIT.

package attachments

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

// NewGetReleaseTaskAttachmentsParams creates a new GetReleaseTaskAttachmentsParams object
// with the default values initialized.
func NewGetReleaseTaskAttachmentsParams() *GetReleaseTaskAttachmentsParams {
	var ()
	return &GetReleaseTaskAttachmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReleaseTaskAttachmentsParamsWithTimeout creates a new GetReleaseTaskAttachmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReleaseTaskAttachmentsParamsWithTimeout(timeout time.Duration) *GetReleaseTaskAttachmentsParams {
	var ()
	return &GetReleaseTaskAttachmentsParams{

		timeout: timeout,
	}
}

// NewGetReleaseTaskAttachmentsParamsWithContext creates a new GetReleaseTaskAttachmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReleaseTaskAttachmentsParamsWithContext(ctx context.Context) *GetReleaseTaskAttachmentsParams {
	var ()
	return &GetReleaseTaskAttachmentsParams{

		Context: ctx,
	}
}

// NewGetReleaseTaskAttachmentsParamsWithHTTPClient creates a new GetReleaseTaskAttachmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReleaseTaskAttachmentsParamsWithHTTPClient(client *http.Client) *GetReleaseTaskAttachmentsParams {
	var ()
	return &GetReleaseTaskAttachmentsParams{
		HTTPClient: client,
	}
}

/*GetReleaseTaskAttachmentsParams contains all the parameters to send to the API endpoint
for the get release task attachments operation typically these are written to a http.Request
*/
type GetReleaseTaskAttachmentsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*AttemptID
	  Attempt number of deployment.

	*/
	AttemptID int32
	/*EnvironmentID
	  Id of the release environment.

	*/
	EnvironmentID int32
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PlanID
	  Plan Id of the deploy phase.

	*/
	PlanID strfmt.UUID
	/*Project
	  Project ID or project name

	*/
	Project string
	/*ReleaseID
	  Id of the release.

	*/
	ReleaseID int32
	/*Type
	  Type of the attachment.

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithTimeout(timeout time.Duration) *GetReleaseTaskAttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithContext(ctx context.Context) *GetReleaseTaskAttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithHTTPClient(client *http.Client) *GetReleaseTaskAttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithAPIVersion(aPIVersion string) *GetReleaseTaskAttachmentsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithAttemptID adds the attemptID to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithAttemptID(attemptID int32) *GetReleaseTaskAttachmentsParams {
	o.SetAttemptID(attemptID)
	return o
}

// SetAttemptID adds the attemptId to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetAttemptID(attemptID int32) {
	o.AttemptID = attemptID
}

// WithEnvironmentID adds the environmentID to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithEnvironmentID(environmentID int32) *GetReleaseTaskAttachmentsParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetEnvironmentID(environmentID int32) {
	o.EnvironmentID = environmentID
}

// WithOrganization adds the organization to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithOrganization(organization string) *GetReleaseTaskAttachmentsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPlanID adds the planID to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithPlanID(planID strfmt.UUID) *GetReleaseTaskAttachmentsParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetPlanID(planID strfmt.UUID) {
	o.PlanID = planID
}

// WithProject adds the project to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithProject(project string) *GetReleaseTaskAttachmentsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetProject(project string) {
	o.Project = project
}

// WithReleaseID adds the releaseID to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithReleaseID(releaseID int32) *GetReleaseTaskAttachmentsParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetReleaseID(releaseID int32) {
	o.ReleaseID = releaseID
}

// WithType adds the typeVar to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) WithType(typeVar string) *GetReleaseTaskAttachmentsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get release task attachments params
func (o *GetReleaseTaskAttachmentsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetReleaseTaskAttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param attemptId
	if err := r.SetPathParam("attemptId", swag.FormatInt32(o.AttemptID)); err != nil {
		return err
	}

	// path param environmentId
	if err := r.SetPathParam("environmentId", swag.FormatInt32(o.EnvironmentID)); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param planId
	if err := r.SetPathParam("planId", o.PlanID.String()); err != nil {
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

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
