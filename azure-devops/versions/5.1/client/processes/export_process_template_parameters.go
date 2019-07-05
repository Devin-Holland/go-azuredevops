// Code generated by go-swagger; DO NOT EDIT.

package processes

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

// NewExportProcessTemplateParams creates a new ExportProcessTemplateParams object
// with the default values initialized.
func NewExportProcessTemplateParams() *ExportProcessTemplateParams {
	var ()
	return &ExportProcessTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExportProcessTemplateParamsWithTimeout creates a new ExportProcessTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportProcessTemplateParamsWithTimeout(timeout time.Duration) *ExportProcessTemplateParams {
	var ()
	return &ExportProcessTemplateParams{

		timeout: timeout,
	}
}

// NewExportProcessTemplateParamsWithContext creates a new ExportProcessTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportProcessTemplateParamsWithContext(ctx context.Context) *ExportProcessTemplateParams {
	var ()
	return &ExportProcessTemplateParams{

		Context: ctx,
	}
}

// NewExportProcessTemplateParamsWithHTTPClient creates a new ExportProcessTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportProcessTemplateParamsWithHTTPClient(client *http.Client) *ExportProcessTemplateParams {
	var ()
	return &ExportProcessTemplateParams{
		HTTPClient: client,
	}
}

/*ExportProcessTemplateParams contains all the parameters to send to the API endpoint
for the export process template operation typically these are written to a http.Request
*/
type ExportProcessTemplateParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*ID
	  The ID of the process

	*/
	ID strfmt.UUID
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the export process template params
func (o *ExportProcessTemplateParams) WithTimeout(timeout time.Duration) *ExportProcessTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export process template params
func (o *ExportProcessTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export process template params
func (o *ExportProcessTemplateParams) WithContext(ctx context.Context) *ExportProcessTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export process template params
func (o *ExportProcessTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export process template params
func (o *ExportProcessTemplateParams) WithHTTPClient(client *http.Client) *ExportProcessTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export process template params
func (o *ExportProcessTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the export process template params
func (o *ExportProcessTemplateParams) WithAPIVersion(aPIVersion string) *ExportProcessTemplateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the export process template params
func (o *ExportProcessTemplateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the export process template params
func (o *ExportProcessTemplateParams) WithID(id strfmt.UUID) *ExportProcessTemplateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the export process template params
func (o *ExportProcessTemplateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithOrganization adds the organization to the export process template params
func (o *ExportProcessTemplateParams) WithOrganization(organization string) *ExportProcessTemplateParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the export process template params
func (o *ExportProcessTemplateParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *ExportProcessTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
