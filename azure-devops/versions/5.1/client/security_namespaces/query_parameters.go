// Code generated by go-swagger; DO NOT EDIT.

package security_namespaces

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

// NewQueryParams creates a new QueryParams object
// with the default values initialized.
func NewQueryParams() *QueryParams {
	var ()
	return &QueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryParamsWithTimeout creates a new QueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryParamsWithTimeout(timeout time.Duration) *QueryParams {
	var ()
	return &QueryParams{

		timeout: timeout,
	}
}

// NewQueryParamsWithContext creates a new QueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryParamsWithContext(ctx context.Context) *QueryParams {
	var ()
	return &QueryParams{

		Context: ctx,
	}
}

// NewQueryParamsWithHTTPClient creates a new QueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryParamsWithHTTPClient(client *http.Client) *QueryParams {
	var ()
	return &QueryParams{
		HTTPClient: client,
	}
}

/*QueryParams contains all the parameters to send to the API endpoint
for the query operation typically these are written to a http.Request
*/
type QueryParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*LocalOnly
	  If true, retrieve only local security namespaces.

	*/
	LocalOnly *bool
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*SecurityNamespaceID
	  Security namespace identifier.

	*/
	SecurityNamespaceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query params
func (o *QueryParams) WithTimeout(timeout time.Duration) *QueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query params
func (o *QueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query params
func (o *QueryParams) WithContext(ctx context.Context) *QueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query params
func (o *QueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query params
func (o *QueryParams) WithHTTPClient(client *http.Client) *QueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query params
func (o *QueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the query params
func (o *QueryParams) WithAPIVersion(aPIVersion string) *QueryParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the query params
func (o *QueryParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithLocalOnly adds the localOnly to the query params
func (o *QueryParams) WithLocalOnly(localOnly *bool) *QueryParams {
	o.SetLocalOnly(localOnly)
	return o
}

// SetLocalOnly adds the localOnly to the query params
func (o *QueryParams) SetLocalOnly(localOnly *bool) {
	o.LocalOnly = localOnly
}

// WithOrganization adds the organization to the query params
func (o *QueryParams) WithOrganization(organization string) *QueryParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the query params
func (o *QueryParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithSecurityNamespaceID adds the securityNamespaceID to the query params
func (o *QueryParams) WithSecurityNamespaceID(securityNamespaceID strfmt.UUID) *QueryParams {
	o.SetSecurityNamespaceID(securityNamespaceID)
	return o
}

// SetSecurityNamespaceID adds the securityNamespaceId to the query params
func (o *QueryParams) SetSecurityNamespaceID(securityNamespaceID strfmt.UUID) {
	o.SecurityNamespaceID = securityNamespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LocalOnly != nil {

		// query param localOnly
		var qrLocalOnly bool
		if o.LocalOnly != nil {
			qrLocalOnly = *o.LocalOnly
		}
		qLocalOnly := swag.FormatBool(qrLocalOnly)
		if qLocalOnly != "" {
			if err := r.SetQueryParam("localOnly", qLocalOnly); err != nil {
				return err
			}
		}

	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param securityNamespaceId
	if err := r.SetPathParam("securityNamespaceId", o.SecurityNamespaceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}