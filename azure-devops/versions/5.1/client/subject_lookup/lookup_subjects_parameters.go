// Code generated by go-swagger; DO NOT EDIT.

package subject_lookup

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

	models "azure-devops/5.1/models"
)

// NewLookupSubjectsParams creates a new LookupSubjectsParams object
// with the default values initialized.
func NewLookupSubjectsParams() *LookupSubjectsParams {
	var ()
	return &LookupSubjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLookupSubjectsParamsWithTimeout creates a new LookupSubjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLookupSubjectsParamsWithTimeout(timeout time.Duration) *LookupSubjectsParams {
	var ()
	return &LookupSubjectsParams{

		timeout: timeout,
	}
}

// NewLookupSubjectsParamsWithContext creates a new LookupSubjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewLookupSubjectsParamsWithContext(ctx context.Context) *LookupSubjectsParams {
	var ()
	return &LookupSubjectsParams{

		Context: ctx,
	}
}

// NewLookupSubjectsParamsWithHTTPClient creates a new LookupSubjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLookupSubjectsParamsWithHTTPClient(client *http.Client) *LookupSubjectsParams {
	var ()
	return &LookupSubjectsParams{
		HTTPClient: client,
	}
}

/*LookupSubjectsParams contains all the parameters to send to the API endpoint
for the lookup subjects operation typically these are written to a http.Request
*/
type LookupSubjectsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  A list of descriptors that specifies a subset of subjects to retrieve. Each descriptor uniquely identifies the subject across all instance scopes, but only at a single point in time.

	*/
	Body *models.GraphSubjectLookup
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lookup subjects params
func (o *LookupSubjectsParams) WithTimeout(timeout time.Duration) *LookupSubjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lookup subjects params
func (o *LookupSubjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lookup subjects params
func (o *LookupSubjectsParams) WithContext(ctx context.Context) *LookupSubjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lookup subjects params
func (o *LookupSubjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lookup subjects params
func (o *LookupSubjectsParams) WithHTTPClient(client *http.Client) *LookupSubjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lookup subjects params
func (o *LookupSubjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the lookup subjects params
func (o *LookupSubjectsParams) WithAPIVersion(aPIVersion string) *LookupSubjectsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the lookup subjects params
func (o *LookupSubjectsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the lookup subjects params
func (o *LookupSubjectsParams) WithBody(body *models.GraphSubjectLookup) *LookupSubjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lookup subjects params
func (o *LookupSubjectsParams) SetBody(body *models.GraphSubjectLookup) {
	o.Body = body
}

// WithOrganization adds the organization to the lookup subjects params
func (o *LookupSubjectsParams) WithOrganization(organization string) *LookupSubjectsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the lookup subjects params
func (o *LookupSubjectsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *LookupSubjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}