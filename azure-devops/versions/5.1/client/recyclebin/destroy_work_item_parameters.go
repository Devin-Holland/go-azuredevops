// Code generated by go-swagger; DO NOT EDIT.

package recyclebin

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

// NewDestroyWorkItemParams creates a new DestroyWorkItemParams object
// with the default values initialized.
func NewDestroyWorkItemParams() *DestroyWorkItemParams {
	var ()
	return &DestroyWorkItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestroyWorkItemParamsWithTimeout creates a new DestroyWorkItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestroyWorkItemParamsWithTimeout(timeout time.Duration) *DestroyWorkItemParams {
	var ()
	return &DestroyWorkItemParams{

		timeout: timeout,
	}
}

// NewDestroyWorkItemParamsWithContext creates a new DestroyWorkItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestroyWorkItemParamsWithContext(ctx context.Context) *DestroyWorkItemParams {
	var ()
	return &DestroyWorkItemParams{

		Context: ctx,
	}
}

// NewDestroyWorkItemParamsWithHTTPClient creates a new DestroyWorkItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestroyWorkItemParamsWithHTTPClient(client *http.Client) *DestroyWorkItemParams {
	var ()
	return &DestroyWorkItemParams{
		HTTPClient: client,
	}
}

/*DestroyWorkItemParams contains all the parameters to send to the API endpoint
for the destroy work item operation typically these are written to a http.Request
*/
type DestroyWorkItemParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*ID
	  ID of the work item to be destroyed permanently

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the destroy work item params
func (o *DestroyWorkItemParams) WithTimeout(timeout time.Duration) *DestroyWorkItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destroy work item params
func (o *DestroyWorkItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destroy work item params
func (o *DestroyWorkItemParams) WithContext(ctx context.Context) *DestroyWorkItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destroy work item params
func (o *DestroyWorkItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destroy work item params
func (o *DestroyWorkItemParams) WithHTTPClient(client *http.Client) *DestroyWorkItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destroy work item params
func (o *DestroyWorkItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the destroy work item params
func (o *DestroyWorkItemParams) WithAPIVersion(aPIVersion string) *DestroyWorkItemParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the destroy work item params
func (o *DestroyWorkItemParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the destroy work item params
func (o *DestroyWorkItemParams) WithID(id int32) *DestroyWorkItemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the destroy work item params
func (o *DestroyWorkItemParams) SetID(id int32) {
	o.ID = id
}

// WithOrganization adds the organization to the destroy work item params
func (o *DestroyWorkItemParams) WithOrganization(organization string) *DestroyWorkItemParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the destroy work item params
func (o *DestroyWorkItemParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the destroy work item params
func (o *DestroyWorkItemParams) WithProject(project string) *DestroyWorkItemParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the destroy work item params
func (o *DestroyWorkItemParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *DestroyWorkItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
