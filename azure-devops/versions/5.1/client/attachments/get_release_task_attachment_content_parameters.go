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

// NewGetReleaseTaskAttachmentContentParams creates a new GetReleaseTaskAttachmentContentParams object
// with the default values initialized.
func NewGetReleaseTaskAttachmentContentParams() *GetReleaseTaskAttachmentContentParams {
	var ()
	return &GetReleaseTaskAttachmentContentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReleaseTaskAttachmentContentParamsWithTimeout creates a new GetReleaseTaskAttachmentContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReleaseTaskAttachmentContentParamsWithTimeout(timeout time.Duration) *GetReleaseTaskAttachmentContentParams {
	var ()
	return &GetReleaseTaskAttachmentContentParams{

		timeout: timeout,
	}
}

// NewGetReleaseTaskAttachmentContentParamsWithContext creates a new GetReleaseTaskAttachmentContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReleaseTaskAttachmentContentParamsWithContext(ctx context.Context) *GetReleaseTaskAttachmentContentParams {
	var ()
	return &GetReleaseTaskAttachmentContentParams{

		Context: ctx,
	}
}

// NewGetReleaseTaskAttachmentContentParamsWithHTTPClient creates a new GetReleaseTaskAttachmentContentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReleaseTaskAttachmentContentParamsWithHTTPClient(client *http.Client) *GetReleaseTaskAttachmentContentParams {
	var ()
	return &GetReleaseTaskAttachmentContentParams{
		HTTPClient: client,
	}
}

/*GetReleaseTaskAttachmentContentParams contains all the parameters to send to the API endpoint
for the get release task attachment content operation typically these are written to a http.Request
*/
type GetReleaseTaskAttachmentContentParams struct {

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
	/*Name
	  Name of the attachment.

	*/
	Name string
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
	/*RecordID
	  Record Id of attachment.

	*/
	RecordID strfmt.UUID
	/*ReleaseID
	  Id of the release.

	*/
	ReleaseID int32
	/*TimelineID
	  Timeline Id of the task.

	*/
	TimelineID strfmt.UUID
	/*Type
	  Type of the attachment.

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithTimeout(timeout time.Duration) *GetReleaseTaskAttachmentContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithContext(ctx context.Context) *GetReleaseTaskAttachmentContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithHTTPClient(client *http.Client) *GetReleaseTaskAttachmentContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithAPIVersion(aPIVersion string) *GetReleaseTaskAttachmentContentParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithAttemptID adds the attemptID to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithAttemptID(attemptID int32) *GetReleaseTaskAttachmentContentParams {
	o.SetAttemptID(attemptID)
	return o
}

// SetAttemptID adds the attemptId to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetAttemptID(attemptID int32) {
	o.AttemptID = attemptID
}

// WithEnvironmentID adds the environmentID to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithEnvironmentID(environmentID int32) *GetReleaseTaskAttachmentContentParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetEnvironmentID(environmentID int32) {
	o.EnvironmentID = environmentID
}

// WithName adds the name to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithName(name string) *GetReleaseTaskAttachmentContentParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetName(name string) {
	o.Name = name
}

// WithOrganization adds the organization to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithOrganization(organization string) *GetReleaseTaskAttachmentContentParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPlanID adds the planID to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithPlanID(planID strfmt.UUID) *GetReleaseTaskAttachmentContentParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetPlanID(planID strfmt.UUID) {
	o.PlanID = planID
}

// WithProject adds the project to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithProject(project string) *GetReleaseTaskAttachmentContentParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetProject(project string) {
	o.Project = project
}

// WithRecordID adds the recordID to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithRecordID(recordID strfmt.UUID) *GetReleaseTaskAttachmentContentParams {
	o.SetRecordID(recordID)
	return o
}

// SetRecordID adds the recordId to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetRecordID(recordID strfmt.UUID) {
	o.RecordID = recordID
}

// WithReleaseID adds the releaseID to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithReleaseID(releaseID int32) *GetReleaseTaskAttachmentContentParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetReleaseID(releaseID int32) {
	o.ReleaseID = releaseID
}

// WithTimelineID adds the timelineID to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithTimelineID(timelineID strfmt.UUID) *GetReleaseTaskAttachmentContentParams {
	o.SetTimelineID(timelineID)
	return o
}

// SetTimelineID adds the timelineId to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetTimelineID(timelineID strfmt.UUID) {
	o.TimelineID = timelineID
}

// WithType adds the typeVar to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) WithType(typeVar string) *GetReleaseTaskAttachmentContentParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get release task attachment content params
func (o *GetReleaseTaskAttachmentContentParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetReleaseTaskAttachmentContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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

	// path param recordId
	if err := r.SetPathParam("recordId", o.RecordID.String()); err != nil {
		return err
	}

	// path param releaseId
	if err := r.SetPathParam("releaseId", swag.FormatInt32(o.ReleaseID)); err != nil {
		return err
	}

	// path param timelineId
	if err := r.SetPathParam("timelineId", o.TimelineID.String()); err != nil {
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
