// Code generated by go-swagger; DO NOT EDIT.

package pages

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

	models "azure-devops/5.1/models"
)

// NewUpdateParams creates a new UpdateParams object
// with the default values initialized.
func NewUpdateParams() *UpdateParams {
	var ()
	return &UpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateParamsWithTimeout creates a new UpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateParamsWithTimeout(timeout time.Duration) *UpdateParams {
	var ()
	return &UpdateParams{

		timeout: timeout,
	}
}

// NewUpdateParamsWithContext creates a new UpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateParamsWithContext(ctx context.Context) *UpdateParams {
	var ()
	return &UpdateParams{

		Context: ctx,
	}
}

// NewUpdateParamsWithHTTPClient creates a new UpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateParamsWithHTTPClient(client *http.Client) *UpdateParams {
	var ()
	return &UpdateParams{
		HTTPClient: client,
	}
}

/*UpdateParams contains all the parameters to send to the API endpoint
for the update operation typically these are written to a http.Request
*/
type UpdateParams struct {

	/*Version
	  Version of the page on which the change is to be made. Mandatory for `Edit` scenario. To be populated in the If-Match header of the request.

	*/
	Version string
	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  Wiki update operation parameters.

	*/
	Body *models.WikiPageCreateOrUpdateParameters
	/*Comment
	  Comment to be associated with the page operation.

	*/
	Comment *string
	/*ID
	  Wiki page id.

	*/
	ID int32
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*WikiIdentifier
	  Wiki Id or name.

	*/
	WikiIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update params
func (o *UpdateParams) WithTimeout(timeout time.Duration) *UpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update params
func (o *UpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update params
func (o *UpdateParams) WithContext(ctx context.Context) *UpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update params
func (o *UpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) WithHTTPClient(client *http.Client) *UpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersion adds the version to the update params
func (o *UpdateParams) WithVersion(version string) *UpdateParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update params
func (o *UpdateParams) SetVersion(version string) {
	o.Version = version
}

// WithAPIVersion adds the aPIVersion to the update params
func (o *UpdateParams) WithAPIVersion(aPIVersion string) *UpdateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update params
func (o *UpdateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update params
func (o *UpdateParams) WithBody(body *models.WikiPageCreateOrUpdateParameters) *UpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update params
func (o *UpdateParams) SetBody(body *models.WikiPageCreateOrUpdateParameters) {
	o.Body = body
}

// WithComment adds the comment to the update params
func (o *UpdateParams) WithComment(comment *string) *UpdateParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the update params
func (o *UpdateParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithID adds the id to the update params
func (o *UpdateParams) WithID(id int32) *UpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update params
func (o *UpdateParams) SetID(id int32) {
	o.ID = id
}

// WithOrganization adds the organization to the update params
func (o *UpdateParams) WithOrganization(organization string) *UpdateParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update params
func (o *UpdateParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the update params
func (o *UpdateParams) WithProject(project string) *UpdateParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update params
func (o *UpdateParams) SetProject(project string) {
	o.Project = project
}

// WithWikiIdentifier adds the wikiIdentifier to the update params
func (o *UpdateParams) WithWikiIdentifier(wikiIdentifier string) *UpdateParams {
	o.SetWikiIdentifier(wikiIdentifier)
	return o
}

// SetWikiIdentifier adds the wikiIdentifier to the update params
func (o *UpdateParams) SetWikiIdentifier(wikiIdentifier string) {
	o.WikiIdentifier = wikiIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Version
	if err := r.SetHeaderParam("Version", o.Version); err != nil {
		return err
	}

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

	if o.Comment != nil {

		// query param comment
		var qrComment string
		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {
			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
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

	// path param wikiIdentifier
	if err := r.SetPathParam("wikiIdentifier", o.WikiIdentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
