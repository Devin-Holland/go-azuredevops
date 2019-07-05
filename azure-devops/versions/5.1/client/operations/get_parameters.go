// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetParams creates a new GetParams object
// with the default values initialized.
func NewGetParams() *GetParams {
	var ()
	return &GetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParamsWithTimeout creates a new GetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParamsWithTimeout(timeout time.Duration) *GetParams {
	var ()
	return &GetParams{

		timeout: timeout,
	}
}

// NewGetParamsWithContext creates a new GetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetParamsWithContext(ctx context.Context) *GetParams {
	var ()
	return &GetParams{

		Context: ctx,
	}
}

// NewGetParamsWithHTTPClient creates a new GetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetParamsWithHTTPClient(client *http.Client) *GetParams {
	var ()
	return &GetParams{
		HTTPClient: client,
	}
}

/*GetParams contains all the parameters to send to the API endpoint
for the get operation typically these are written to a http.Request
*/
type GetParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*OperationID
	  The ID for the operation.

	*/
	OperationID strfmt.UUID
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PluginID
	  The ID for the plugin.

	*/
	PluginID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get params
func (o *GetParams) WithTimeout(timeout time.Duration) *GetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get params
func (o *GetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get params
func (o *GetParams) WithContext(ctx context.Context) *GetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get params
func (o *GetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get params
func (o *GetParams) WithHTTPClient(client *http.Client) *GetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get params
func (o *GetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get params
func (o *GetParams) WithAPIVersion(aPIVersion string) *GetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get params
func (o *GetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOperationID adds the operationID to the get params
func (o *GetParams) WithOperationID(operationID strfmt.UUID) *GetParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the get params
func (o *GetParams) SetOperationID(operationID strfmt.UUID) {
	o.OperationID = operationID
}

// WithOrganization adds the organization to the get params
func (o *GetParams) WithOrganization(organization string) *GetParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get params
func (o *GetParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPluginID adds the pluginID to the get params
func (o *GetParams) WithPluginID(pluginID *strfmt.UUID) *GetParams {
	o.SetPluginID(pluginID)
	return o
}

// SetPluginID adds the pluginId to the get params
func (o *GetParams) SetPluginID(pluginID *strfmt.UUID) {
	o.PluginID = pluginID
}

// WriteToRequest writes these params to a swagger request
func (o *GetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param operationId
	if err := r.SetPathParam("operationId", o.OperationID.String()); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	if o.PluginID != nil {

		// query param pluginId
		var qrPluginID strfmt.UUID
		if o.PluginID != nil {
			qrPluginID = *o.PluginID
		}
		qPluginID := qrPluginID.String()
		if qPluginID != "" {
			if err := r.SetQueryParam("pluginId", qPluginID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
