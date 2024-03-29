// Code generated by go-swagger; DO NOT EDIT.

package resources

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

// NewAuthorizeDefinitionResourcesParams creates a new AuthorizeDefinitionResourcesParams object
// with the default values initialized.
func NewAuthorizeDefinitionResourcesParams() *AuthorizeDefinitionResourcesParams {
	var ()
	return &AuthorizeDefinitionResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthorizeDefinitionResourcesParamsWithTimeout creates a new AuthorizeDefinitionResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthorizeDefinitionResourcesParamsWithTimeout(timeout time.Duration) *AuthorizeDefinitionResourcesParams {
	var ()
	return &AuthorizeDefinitionResourcesParams{

		timeout: timeout,
	}
}

// NewAuthorizeDefinitionResourcesParamsWithContext creates a new AuthorizeDefinitionResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthorizeDefinitionResourcesParamsWithContext(ctx context.Context) *AuthorizeDefinitionResourcesParams {
	var ()
	return &AuthorizeDefinitionResourcesParams{

		Context: ctx,
	}
}

// NewAuthorizeDefinitionResourcesParamsWithHTTPClient creates a new AuthorizeDefinitionResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthorizeDefinitionResourcesParamsWithHTTPClient(client *http.Client) *AuthorizeDefinitionResourcesParams {
	var ()
	return &AuthorizeDefinitionResourcesParams{
		HTTPClient: client,
	}
}

/*AuthorizeDefinitionResourcesParams contains all the parameters to send to the API endpoint
for the authorize definition resources operation typically these are written to a http.Request
*/
type AuthorizeDefinitionResourcesParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body*/
	Body []*models.DefinitionResourceReference
	/*DefinitionID*/
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

// WithTimeout adds the timeout to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) WithTimeout(timeout time.Duration) *AuthorizeDefinitionResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) WithContext(ctx context.Context) *AuthorizeDefinitionResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) WithHTTPClient(client *http.Client) *AuthorizeDefinitionResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) WithAPIVersion(aPIVersion string) *AuthorizeDefinitionResourcesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) WithBody(body []*models.DefinitionResourceReference) *AuthorizeDefinitionResourcesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) SetBody(body []*models.DefinitionResourceReference) {
	o.Body = body
}

// WithDefinitionID adds the definitionID to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) WithDefinitionID(definitionID int32) *AuthorizeDefinitionResourcesParams {
	o.SetDefinitionID(definitionID)
	return o
}

// SetDefinitionID adds the definitionId to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) SetDefinitionID(definitionID int32) {
	o.DefinitionID = definitionID
}

// WithOrganization adds the organization to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) WithOrganization(organization string) *AuthorizeDefinitionResourcesParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) WithProject(project string) *AuthorizeDefinitionResourcesParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the authorize definition resources params
func (o *AuthorizeDefinitionResourcesParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *AuthorizeDefinitionResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
