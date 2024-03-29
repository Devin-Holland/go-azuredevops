// Code generated by go-swagger; DO NOT EDIT.

package metrics

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

// NewGetDefinitionMetricsParams creates a new GetDefinitionMetricsParams object
// with the default values initialized.
func NewGetDefinitionMetricsParams() *GetDefinitionMetricsParams {
	var ()
	return &GetDefinitionMetricsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDefinitionMetricsParamsWithTimeout creates a new GetDefinitionMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDefinitionMetricsParamsWithTimeout(timeout time.Duration) *GetDefinitionMetricsParams {
	var ()
	return &GetDefinitionMetricsParams{

		timeout: timeout,
	}
}

// NewGetDefinitionMetricsParamsWithContext creates a new GetDefinitionMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDefinitionMetricsParamsWithContext(ctx context.Context) *GetDefinitionMetricsParams {
	var ()
	return &GetDefinitionMetricsParams{

		Context: ctx,
	}
}

// NewGetDefinitionMetricsParamsWithHTTPClient creates a new GetDefinitionMetricsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDefinitionMetricsParamsWithHTTPClient(client *http.Client) *GetDefinitionMetricsParams {
	var ()
	return &GetDefinitionMetricsParams{
		HTTPClient: client,
	}
}

/*GetDefinitionMetricsParams contains all the parameters to send to the API endpoint
for the get definition metrics operation typically these are written to a http.Request
*/
type GetDefinitionMetricsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*DefinitionID
	  The ID of the definition.

	*/
	DefinitionID int32
	/*MinMetricsTime
	  The date from which to calculate metrics.

	*/
	MinMetricsTime *strfmt.DateTime
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

// WithTimeout adds the timeout to the get definition metrics params
func (o *GetDefinitionMetricsParams) WithTimeout(timeout time.Duration) *GetDefinitionMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get definition metrics params
func (o *GetDefinitionMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get definition metrics params
func (o *GetDefinitionMetricsParams) WithContext(ctx context.Context) *GetDefinitionMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get definition metrics params
func (o *GetDefinitionMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get definition metrics params
func (o *GetDefinitionMetricsParams) WithHTTPClient(client *http.Client) *GetDefinitionMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get definition metrics params
func (o *GetDefinitionMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get definition metrics params
func (o *GetDefinitionMetricsParams) WithAPIVersion(aPIVersion string) *GetDefinitionMetricsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get definition metrics params
func (o *GetDefinitionMetricsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDefinitionID adds the definitionID to the get definition metrics params
func (o *GetDefinitionMetricsParams) WithDefinitionID(definitionID int32) *GetDefinitionMetricsParams {
	o.SetDefinitionID(definitionID)
	return o
}

// SetDefinitionID adds the definitionId to the get definition metrics params
func (o *GetDefinitionMetricsParams) SetDefinitionID(definitionID int32) {
	o.DefinitionID = definitionID
}

// WithMinMetricsTime adds the minMetricsTime to the get definition metrics params
func (o *GetDefinitionMetricsParams) WithMinMetricsTime(minMetricsTime *strfmt.DateTime) *GetDefinitionMetricsParams {
	o.SetMinMetricsTime(minMetricsTime)
	return o
}

// SetMinMetricsTime adds the minMetricsTime to the get definition metrics params
func (o *GetDefinitionMetricsParams) SetMinMetricsTime(minMetricsTime *strfmt.DateTime) {
	o.MinMetricsTime = minMetricsTime
}

// WithOrganization adds the organization to the get definition metrics params
func (o *GetDefinitionMetricsParams) WithOrganization(organization string) *GetDefinitionMetricsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get definition metrics params
func (o *GetDefinitionMetricsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get definition metrics params
func (o *GetDefinitionMetricsParams) WithProject(project string) *GetDefinitionMetricsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get definition metrics params
func (o *GetDefinitionMetricsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetDefinitionMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MinMetricsTime != nil {

		// query param minMetricsTime
		var qrMinMetricsTime strfmt.DateTime
		if o.MinMetricsTime != nil {
			qrMinMetricsTime = *o.MinMetricsTime
		}
		qMinMetricsTime := qrMinMetricsTime.String()
		if qMinMetricsTime != "" {
			if err := r.SetQueryParam("minMetricsTime", qMinMetricsTime); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
