// Code generated by go-swagger; DO NOT EDIT.

package variablegroups

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
)

// NewGetVariableGroupsByIDParams creates a new GetVariableGroupsByIDParams object
// with the default values initialized.
func NewGetVariableGroupsByIDParams() *GetVariableGroupsByIDParams {
	var ()
	return &GetVariableGroupsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVariableGroupsByIDParamsWithTimeout creates a new GetVariableGroupsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVariableGroupsByIDParamsWithTimeout(timeout time.Duration) *GetVariableGroupsByIDParams {
	var ()
	return &GetVariableGroupsByIDParams{

		timeout: timeout,
	}
}

// NewGetVariableGroupsByIDParamsWithContext creates a new GetVariableGroupsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVariableGroupsByIDParamsWithContext(ctx context.Context) *GetVariableGroupsByIDParams {
	var ()
	return &GetVariableGroupsByIDParams{

		Context: ctx,
	}
}

// NewGetVariableGroupsByIDParamsWithHTTPClient creates a new GetVariableGroupsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVariableGroupsByIDParamsWithHTTPClient(client *http.Client) *GetVariableGroupsByIDParams {
	var ()
	return &GetVariableGroupsByIDParams{
		HTTPClient: client,
	}
}

/*GetVariableGroupsByIDParams contains all the parameters to send to the API endpoint
for the get variable groups by Id operation typically these are written to a http.Request
*/
type GetVariableGroupsByIDParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*GroupIds
	  Comma separated list of Ids of variable groups.

	*/
	GroupIds string
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

// WithTimeout adds the timeout to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) WithTimeout(timeout time.Duration) *GetVariableGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) WithContext(ctx context.Context) *GetVariableGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) WithHTTPClient(client *http.Client) *GetVariableGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) WithAPIVersion(aPIVersion string) *GetVariableGroupsByIDParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGroupIds adds the groupIds to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) WithGroupIds(groupIds string) *GetVariableGroupsByIDParams {
	o.SetGroupIds(groupIds)
	return o
}

// SetGroupIds adds the groupIds to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) SetGroupIds(groupIds string) {
	o.GroupIds = groupIds
}

// WithOrganization adds the organization to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) WithOrganization(organization string) *GetVariableGroupsByIDParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) WithProject(project string) *GetVariableGroupsByIDParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get variable groups by Id params
func (o *GetVariableGroupsByIDParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetVariableGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param groupIds
	qrGroupIds := o.GroupIds
	qGroupIds := qrGroupIds
	if qGroupIds != "" {
		if err := r.SetQueryParam("groupIds", qGroupIds); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
