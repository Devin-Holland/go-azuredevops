// Code generated by go-swagger; DO NOT EDIT.

package evaluations

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

// NewRequeuePolicyEvaluationParams creates a new RequeuePolicyEvaluationParams object
// with the default values initialized.
func NewRequeuePolicyEvaluationParams() *RequeuePolicyEvaluationParams {
	var ()
	return &RequeuePolicyEvaluationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequeuePolicyEvaluationParamsWithTimeout creates a new RequeuePolicyEvaluationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequeuePolicyEvaluationParamsWithTimeout(timeout time.Duration) *RequeuePolicyEvaluationParams {
	var ()
	return &RequeuePolicyEvaluationParams{

		timeout: timeout,
	}
}

// NewRequeuePolicyEvaluationParamsWithContext creates a new RequeuePolicyEvaluationParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequeuePolicyEvaluationParamsWithContext(ctx context.Context) *RequeuePolicyEvaluationParams {
	var ()
	return &RequeuePolicyEvaluationParams{

		Context: ctx,
	}
}

// NewRequeuePolicyEvaluationParamsWithHTTPClient creates a new RequeuePolicyEvaluationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequeuePolicyEvaluationParamsWithHTTPClient(client *http.Client) *RequeuePolicyEvaluationParams {
	var ()
	return &RequeuePolicyEvaluationParams{
		HTTPClient: client,
	}
}

/*RequeuePolicyEvaluationParams contains all the parameters to send to the API endpoint
for the requeue policy evaluation operation typically these are written to a http.Request
*/
type RequeuePolicyEvaluationParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*EvaluationID
	  ID of the policy evaluation to be retrieved.

	*/
	EvaluationID strfmt.UUID
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

// WithTimeout adds the timeout to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) WithTimeout(timeout time.Duration) *RequeuePolicyEvaluationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) WithContext(ctx context.Context) *RequeuePolicyEvaluationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) WithHTTPClient(client *http.Client) *RequeuePolicyEvaluationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) WithAPIVersion(aPIVersion string) *RequeuePolicyEvaluationParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithEvaluationID adds the evaluationID to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) WithEvaluationID(evaluationID strfmt.UUID) *RequeuePolicyEvaluationParams {
	o.SetEvaluationID(evaluationID)
	return o
}

// SetEvaluationID adds the evaluationId to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) SetEvaluationID(evaluationID strfmt.UUID) {
	o.EvaluationID = evaluationID
}

// WithOrganization adds the organization to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) WithOrganization(organization string) *RequeuePolicyEvaluationParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) WithProject(project string) *RequeuePolicyEvaluationParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the requeue policy evaluation params
func (o *RequeuePolicyEvaluationParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *RequeuePolicyEvaluationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param evaluationId
	if err := r.SetPathParam("evaluationId", o.EvaluationID.String()); err != nil {
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
