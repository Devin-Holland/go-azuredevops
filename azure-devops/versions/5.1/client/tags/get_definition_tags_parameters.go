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

// NewGetDefinitionTagsParams creates a new GetDefinitionTagsParams object
// with the default values initialized.
func NewGetDefinitionTagsParams() *GetDefinitionTagsParams {
	var ()
	return &GetDefinitionTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDefinitionTagsParamsWithTimeout creates a new GetDefinitionTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDefinitionTagsParamsWithTimeout(timeout time.Duration) *GetDefinitionTagsParams {
	var ()
	return &GetDefinitionTagsParams{

		timeout: timeout,
	}
}

// NewGetDefinitionTagsParamsWithContext creates a new GetDefinitionTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDefinitionTagsParamsWithContext(ctx context.Context) *GetDefinitionTagsParams {
	var ()
	return &GetDefinitionTagsParams{

		Context: ctx,
	}
}

// NewGetDefinitionTagsParamsWithHTTPClient creates a new GetDefinitionTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDefinitionTagsParamsWithHTTPClient(client *http.Client) *GetDefinitionTagsParams {
	var ()
	return &GetDefinitionTagsParams{
		HTTPClient: client,
	}
}

/*GetDefinitionTagsParams contains all the parameters to send to the API endpoint
for the get definition tags operation typically these are written to a http.Request
*/
type GetDefinitionTagsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
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
	/*Revision
	  The definition revision number. If not specified, uses the latest revision of the definition.

	*/
	Revision *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get definition tags params
func (o *GetDefinitionTagsParams) WithTimeout(timeout time.Duration) *GetDefinitionTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get definition tags params
func (o *GetDefinitionTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get definition tags params
func (o *GetDefinitionTagsParams) WithContext(ctx context.Context) *GetDefinitionTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get definition tags params
func (o *GetDefinitionTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get definition tags params
func (o *GetDefinitionTagsParams) WithHTTPClient(client *http.Client) *GetDefinitionTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get definition tags params
func (o *GetDefinitionTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get definition tags params
func (o *GetDefinitionTagsParams) WithAPIVersion(aPIVersion string) *GetDefinitionTagsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get definition tags params
func (o *GetDefinitionTagsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDefinitionID adds the definitionID to the get definition tags params
func (o *GetDefinitionTagsParams) WithDefinitionID(definitionID int32) *GetDefinitionTagsParams {
	o.SetDefinitionID(definitionID)
	return o
}

// SetDefinitionID adds the definitionId to the get definition tags params
func (o *GetDefinitionTagsParams) SetDefinitionID(definitionID int32) {
	o.DefinitionID = definitionID
}

// WithOrganization adds the organization to the get definition tags params
func (o *GetDefinitionTagsParams) WithOrganization(organization string) *GetDefinitionTagsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get definition tags params
func (o *GetDefinitionTagsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get definition tags params
func (o *GetDefinitionTagsParams) WithProject(project string) *GetDefinitionTagsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get definition tags params
func (o *GetDefinitionTagsParams) SetProject(project string) {
	o.Project = project
}

// WithRevision adds the revision to the get definition tags params
func (o *GetDefinitionTagsParams) WithRevision(revision *int32) *GetDefinitionTagsParams {
	o.SetRevision(revision)
	return o
}

// SetRevision adds the revision to the get definition tags params
func (o *GetDefinitionTagsParams) SetRevision(revision *int32) {
	o.Revision = revision
}

// WriteToRequest writes these params to a swagger request
func (o *GetDefinitionTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Revision != nil {

		// query param revision
		var qrRevision int32
		if o.Revision != nil {
			qrRevision = *o.Revision
		}
		qRevision := swag.FormatInt32(qrRevision)
		if qRevision != "" {
			if err := r.SetQueryParam("revision", qRevision); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}