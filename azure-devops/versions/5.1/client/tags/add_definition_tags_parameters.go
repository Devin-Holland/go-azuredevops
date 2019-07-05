// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewAddDefinitionTagsParams creates a new AddDefinitionTagsParams object
// with the default values initialized.
func NewAddDefinitionTagsParams() *AddDefinitionTagsParams {
	var ()
	return &AddDefinitionTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddDefinitionTagsParamsWithTimeout creates a new AddDefinitionTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddDefinitionTagsParamsWithTimeout(timeout time.Duration) *AddDefinitionTagsParams {
	var ()
	return &AddDefinitionTagsParams{

		timeout: timeout,
	}
}

// NewAddDefinitionTagsParamsWithContext creates a new AddDefinitionTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddDefinitionTagsParamsWithContext(ctx context.Context) *AddDefinitionTagsParams {
	var ()
	return &AddDefinitionTagsParams{

		Context: ctx,
	}
}

// NewAddDefinitionTagsParamsWithHTTPClient creates a new AddDefinitionTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddDefinitionTagsParamsWithHTTPClient(client *http.Client) *AddDefinitionTagsParams {
	var ()
	return &AddDefinitionTagsParams{
		HTTPClient: client,
	}
}

/*AddDefinitionTagsParams contains all the parameters to send to the API endpoint
for the add definition tags operation typically these are written to a http.Request
*/
type AddDefinitionTagsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  The tags to add.

	*/
	Body []string
	/*DefinitionID
	  The ID of the definition.

	*/
	DefinitionID int32
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

// WithTimeout adds the timeout to the add definition tags params
func (o *AddDefinitionTagsParams) WithTimeout(timeout time.Duration) *AddDefinitionTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add definition tags params
func (o *AddDefinitionTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add definition tags params
func (o *AddDefinitionTagsParams) WithContext(ctx context.Context) *AddDefinitionTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add definition tags params
func (o *AddDefinitionTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add definition tags params
func (o *AddDefinitionTagsParams) WithHTTPClient(client *http.Client) *AddDefinitionTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add definition tags params
func (o *AddDefinitionTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the add definition tags params
func (o *AddDefinitionTagsParams) WithAPIVersion(aPIVersion string) *AddDefinitionTagsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the add definition tags params
func (o *AddDefinitionTagsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the add definition tags params
func (o *AddDefinitionTagsParams) WithBody(body []string) *AddDefinitionTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add definition tags params
func (o *AddDefinitionTagsParams) SetBody(body []string) {
	o.Body = body
}

// WithDefinitionID adds the definitionID to the add definition tags params
func (o *AddDefinitionTagsParams) WithDefinitionID(definitionID int32) *AddDefinitionTagsParams {
	o.SetDefinitionID(definitionID)
	return o
}

// SetDefinitionID adds the definitionId to the add definition tags params
func (o *AddDefinitionTagsParams) SetDefinitionID(definitionID int32) {
	o.DefinitionID = definitionID
}

// WithOrganization adds the organization to the add definition tags params
func (o *AddDefinitionTagsParams) WithOrganization(organization string) *AddDefinitionTagsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the add definition tags params
func (o *AddDefinitionTagsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the add definition tags params
func (o *AddDefinitionTagsParams) WithProject(project string) *AddDefinitionTagsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the add definition tags params
func (o *AddDefinitionTagsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *AddDefinitionTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param definitionId
	if err := r.SetPathParam("definitionId", swag.FormatInt32(o.DefinitionID)); err != nil {
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
