// Code generated by go-swagger; DO NOT EDIT.

package availability

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

// NewCheckAvailabilityParams creates a new CheckAvailabilityParams object
// with the default values initialized.
func NewCheckAvailabilityParams() *CheckAvailabilityParams {
	var ()
	return &CheckAvailabilityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckAvailabilityParamsWithTimeout creates a new CheckAvailabilityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckAvailabilityParamsWithTimeout(timeout time.Duration) *CheckAvailabilityParams {
	var ()
	return &CheckAvailabilityParams{

		timeout: timeout,
	}
}

// NewCheckAvailabilityParamsWithContext creates a new CheckAvailabilityParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckAvailabilityParamsWithContext(ctx context.Context) *CheckAvailabilityParams {
	var ()
	return &CheckAvailabilityParams{

		Context: ctx,
	}
}

// NewCheckAvailabilityParamsWithHTTPClient creates a new CheckAvailabilityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckAvailabilityParamsWithHTTPClient(client *http.Client) *CheckAvailabilityParams {
	var ()
	return &CheckAvailabilityParams{
		HTTPClient: client,
	}
}

/*CheckAvailabilityParams contains all the parameters to send to the API endpoint
for the check availability operation typically these are written to a http.Request
*/
type CheckAvailabilityParams struct {

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

// WithTimeout adds the timeout to the check availability params
func (o *CheckAvailabilityParams) WithTimeout(timeout time.Duration) *CheckAvailabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check availability params
func (o *CheckAvailabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check availability params
func (o *CheckAvailabilityParams) WithContext(ctx context.Context) *CheckAvailabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check availability params
func (o *CheckAvailabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check availability params
func (o *CheckAvailabilityParams) WithHTTPClient(client *http.Client) *CheckAvailabilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check availability params
func (o *CheckAvailabilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the check availability params
func (o *CheckAvailabilityParams) WithAPIVersion(aPIVersion string) *CheckAvailabilityParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the check availability params
func (o *CheckAvailabilityParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the check availability params
func (o *CheckAvailabilityParams) WithOrganization(organization string) *CheckAvailabilityParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the check availability params
func (o *CheckAvailabilityParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *CheckAvailabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
