// Code generated by go-swagger; DO NOT EDIT.

package comments

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

// NewGetCommentsBatchParams creates a new GetCommentsBatchParams object
// with the default values initialized.
func NewGetCommentsBatchParams() *GetCommentsBatchParams {
	var ()
	return &GetCommentsBatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCommentsBatchParamsWithTimeout creates a new GetCommentsBatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCommentsBatchParamsWithTimeout(timeout time.Duration) *GetCommentsBatchParams {
	var ()
	return &GetCommentsBatchParams{

		timeout: timeout,
	}
}

// NewGetCommentsBatchParamsWithContext creates a new GetCommentsBatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCommentsBatchParamsWithContext(ctx context.Context) *GetCommentsBatchParams {
	var ()
	return &GetCommentsBatchParams{

		Context: ctx,
	}
}

// NewGetCommentsBatchParamsWithHTTPClient creates a new GetCommentsBatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCommentsBatchParamsWithHTTPClient(client *http.Client) *GetCommentsBatchParams {
	var ()
	return &GetCommentsBatchParams{
		HTTPClient: client,
	}
}

/*GetCommentsBatchParams contains all the parameters to send to the API endpoint
for the get comments batch operation typically these are written to a http.Request
*/
type GetCommentsBatchParams struct {

	/*NrDollarExpand
	  Specifies the additional data retrieval options for work item comments.

	*/
	DollarExpand *string
	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.3' to use this version of the api.

	*/
	APIVersion string
	/*Ids
	  Comma-separated list of comment ids to return.

	*/
	Ids string
	/*IncludeDeleted
	  Specify if the deleted comments should be retrieved.

	*/
	IncludeDeleted *bool
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*WorkItemID
	  Id of a work item to get comments for.

	*/
	WorkItemID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get comments batch params
func (o *GetCommentsBatchParams) WithTimeout(timeout time.Duration) *GetCommentsBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get comments batch params
func (o *GetCommentsBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get comments batch params
func (o *GetCommentsBatchParams) WithContext(ctx context.Context) *GetCommentsBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get comments batch params
func (o *GetCommentsBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get comments batch params
func (o *GetCommentsBatchParams) WithHTTPClient(client *http.Client) *GetCommentsBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get comments batch params
func (o *GetCommentsBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarExpand adds the dollarExpand to the get comments batch params
func (o *GetCommentsBatchParams) WithDollarExpand(dollarExpand *string) *GetCommentsBatchParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the get comments batch params
func (o *GetCommentsBatchParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithAPIVersion adds the aPIVersion to the get comments batch params
func (o *GetCommentsBatchParams) WithAPIVersion(aPIVersion string) *GetCommentsBatchParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get comments batch params
func (o *GetCommentsBatchParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithIds adds the ids to the get comments batch params
func (o *GetCommentsBatchParams) WithIds(ids string) *GetCommentsBatchParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get comments batch params
func (o *GetCommentsBatchParams) SetIds(ids string) {
	o.Ids = ids
}

// WithIncludeDeleted adds the includeDeleted to the get comments batch params
func (o *GetCommentsBatchParams) WithIncludeDeleted(includeDeleted *bool) *GetCommentsBatchParams {
	o.SetIncludeDeleted(includeDeleted)
	return o
}

// SetIncludeDeleted adds the includeDeleted to the get comments batch params
func (o *GetCommentsBatchParams) SetIncludeDeleted(includeDeleted *bool) {
	o.IncludeDeleted = includeDeleted
}

// WithOrganization adds the organization to the get comments batch params
func (o *GetCommentsBatchParams) WithOrganization(organization string) *GetCommentsBatchParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get comments batch params
func (o *GetCommentsBatchParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get comments batch params
func (o *GetCommentsBatchParams) WithProject(project string) *GetCommentsBatchParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get comments batch params
func (o *GetCommentsBatchParams) SetProject(project string) {
	o.Project = project
}

// WithWorkItemID adds the workItemID to the get comments batch params
func (o *GetCommentsBatchParams) WithWorkItemID(workItemID int32) *GetCommentsBatchParams {
	o.SetWorkItemID(workItemID)
	return o
}

// SetWorkItemID adds the workItemId to the get comments batch params
func (o *GetCommentsBatchParams) SetWorkItemID(workItemID int32) {
	o.WorkItemID = workItemID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCommentsBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarExpand != nil {

		// query param $expand
		var qrNrDollarExpand string
		if o.DollarExpand != nil {
			qrNrDollarExpand = *o.DollarExpand
		}
		qNrDollarExpand := qrNrDollarExpand
		if qNrDollarExpand != "" {
			if err := r.SetQueryParam("$expand", qNrDollarExpand); err != nil {
				return err
			}
		}

	}

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {
		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if o.IncludeDeleted != nil {

		// query param includeDeleted
		var qrIncludeDeleted bool
		if o.IncludeDeleted != nil {
			qrIncludeDeleted = *o.IncludeDeleted
		}
		qIncludeDeleted := swag.FormatBool(qrIncludeDeleted)
		if qIncludeDeleted != "" {
			if err := r.SetQueryParam("includeDeleted", qIncludeDeleted); err != nil {
				return err
			}
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

	// path param workItemId
	if err := r.SetPathParam("workItemId", swag.FormatInt32(o.WorkItemID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
