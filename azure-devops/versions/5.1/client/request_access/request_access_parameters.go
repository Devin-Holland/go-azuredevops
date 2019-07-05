// Code generated by go-swagger; DO NOT EDIT.

package request_access

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

// NewRequestAccessParams creates a new RequestAccessParams object
// with the default values initialized.
func NewRequestAccessParams() *RequestAccessParams {
	var ()
	return &RequestAccessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestAccessParamsWithTimeout creates a new RequestAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestAccessParamsWithTimeout(timeout time.Duration) *RequestAccessParams {
	var ()
	return &RequestAccessParams{

		timeout: timeout,
	}
}

// NewRequestAccessParamsWithContext creates a new RequestAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestAccessParamsWithContext(ctx context.Context) *RequestAccessParams {
	var ()
	return &RequestAccessParams{

		Context: ctx,
	}
}

// NewRequestAccessParamsWithHTTPClient creates a new RequestAccessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestAccessParamsWithHTTPClient(client *http.Client) *RequestAccessParams {
	var ()
	return &RequestAccessParams{
		HTTPClient: client,
	}
}

/*RequestAccessParams contains all the parameters to send to the API endpoint
for the request access operation typically these are written to a http.Request
*/
type RequestAccessParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body*/
	Body string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request access params
func (o *RequestAccessParams) WithTimeout(timeout time.Duration) *RequestAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request access params
func (o *RequestAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request access params
func (o *RequestAccessParams) WithContext(ctx context.Context) *RequestAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request access params
func (o *RequestAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request access params
func (o *RequestAccessParams) WithHTTPClient(client *http.Client) *RequestAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request access params
func (o *RequestAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the request access params
func (o *RequestAccessParams) WithAPIVersion(aPIVersion string) *RequestAccessParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the request access params
func (o *RequestAccessParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the request access params
func (o *RequestAccessParams) WithBody(body string) *RequestAccessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the request access params
func (o *RequestAccessParams) SetBody(body string) {
	o.Body = body
}

// WithOrganization adds the organization to the request access params
func (o *RequestAccessParams) WithOrganization(organization string) *RequestAccessParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the request access params
func (o *RequestAccessParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *RequestAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.Body); err != nil {
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