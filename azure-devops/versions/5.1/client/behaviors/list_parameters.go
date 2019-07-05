// Code generated by go-swagger; DO NOT EDIT.

package behaviors

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

// NewListParams creates a new ListParams object
// with the default values initialized.
func NewListParams() *ListParams {
	var ()
	return &ListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListParamsWithTimeout creates a new ListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListParamsWithTimeout(timeout time.Duration) *ListParams {
	var ()
	return &ListParams{

		timeout: timeout,
	}
}

// NewListParamsWithContext creates a new ListParams object
// with the default values initialized, and the ability to set a context for a request
func NewListParamsWithContext(ctx context.Context) *ListParams {
	var ()
	return &ListParams{

		Context: ctx,
	}
}

// NewListParamsWithHTTPClient creates a new ListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListParamsWithHTTPClient(client *http.Client) *ListParams {
	var ()
	return &ListParams{
		HTTPClient: client,
	}
}

/*ListParams contains all the parameters to send to the API endpoint
for the list operation typically these are written to a http.Request
*/
type ListParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*ProcessID
	  The ID of the process

	*/
	ProcessID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list params
func (o *ListParams) WithTimeout(timeout time.Duration) *ListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list params
func (o *ListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list params
func (o *ListParams) WithContext(ctx context.Context) *ListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list params
func (o *ListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list params
func (o *ListParams) WithHTTPClient(client *http.Client) *ListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list params
func (o *ListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the list params
func (o *ListParams) WithAPIVersion(aPIVersion string) *ListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list params
func (o *ListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the list params
func (o *ListParams) WithOrganization(organization string) *ListParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the list params
func (o *ListParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProcessID adds the processID to the list params
func (o *ListParams) WithProcessID(processID strfmt.UUID) *ListParams {
	o.SetProcessID(processID)
	return o
}

// SetProcessID adds the processId to the list params
func (o *ListParams) SetProcessID(processID strfmt.UUID) {
	o.ProcessID = processID
}

// WriteToRequest writes these params to a swagger request
func (o *ListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param processId
	if err := r.SetPathParam("processId", o.ProcessID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}