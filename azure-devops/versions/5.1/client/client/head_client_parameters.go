// Code generated by go-swagger; DO NOT EDIT.

package client

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

// NewHeadClientParams creates a new HeadClientParams object
// with the default values initialized.
func NewHeadClientParams() *HeadClientParams {
	var ()
	return &HeadClientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHeadClientParamsWithTimeout creates a new HeadClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHeadClientParamsWithTimeout(timeout time.Duration) *HeadClientParams {
	var ()
	return &HeadClientParams{

		timeout: timeout,
	}
}

// NewHeadClientParamsWithContext creates a new HeadClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewHeadClientParamsWithContext(ctx context.Context) *HeadClientParams {
	var ()
	return &HeadClientParams{

		Context: ctx,
	}
}

// NewHeadClientParamsWithHTTPClient creates a new HeadClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHeadClientParamsWithHTTPClient(client *http.Client) *HeadClientParams {
	var ()
	return &HeadClientParams{
		HTTPClient: client,
	}
}

/*HeadClientParams contains all the parameters to send to the API endpoint
for the head client operation typically these are written to a http.Request
*/
type HeadClientParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the head client params
func (o *HeadClientParams) WithTimeout(timeout time.Duration) *HeadClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head client params
func (o *HeadClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head client params
func (o *HeadClientParams) WithContext(ctx context.Context) *HeadClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head client params
func (o *HeadClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head client params
func (o *HeadClientParams) WithHTTPClient(client *http.Client) *HeadClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head client params
func (o *HeadClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the head client params
func (o *HeadClientParams) WithAPIVersion(aPIVersion string) *HeadClientParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the head client params
func (o *HeadClientParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the head client params
func (o *HeadClientParams) WithOrganization(organization string) *HeadClientParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the head client params
func (o *HeadClientParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *HeadClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}