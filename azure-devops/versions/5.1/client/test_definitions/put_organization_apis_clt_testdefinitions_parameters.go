// Code generated by go-swagger; DO NOT EDIT.

package test_definitions

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

// NewPutOrganizationApisCltTestdefinitionsParams creates a new PutOrganizationApisCltTestdefinitionsParams object
// with the default values initialized.
func NewPutOrganizationApisCltTestdefinitionsParams() *PutOrganizationApisCltTestdefinitionsParams {
	var ()
	return &PutOrganizationApisCltTestdefinitionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrganizationApisCltTestdefinitionsParamsWithTimeout creates a new PutOrganizationApisCltTestdefinitionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOrganizationApisCltTestdefinitionsParamsWithTimeout(timeout time.Duration) *PutOrganizationApisCltTestdefinitionsParams {
	var ()
	return &PutOrganizationApisCltTestdefinitionsParams{

		timeout: timeout,
	}
}

// NewPutOrganizationApisCltTestdefinitionsParamsWithContext creates a new PutOrganizationApisCltTestdefinitionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOrganizationApisCltTestdefinitionsParamsWithContext(ctx context.Context) *PutOrganizationApisCltTestdefinitionsParams {
	var ()
	return &PutOrganizationApisCltTestdefinitionsParams{

		Context: ctx,
	}
}

// NewPutOrganizationApisCltTestdefinitionsParamsWithHTTPClient creates a new PutOrganizationApisCltTestdefinitionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOrganizationApisCltTestdefinitionsParamsWithHTTPClient(client *http.Client) *PutOrganizationApisCltTestdefinitionsParams {
	var ()
	return &PutOrganizationApisCltTestdefinitionsParams{
		HTTPClient: client,
	}
}

/*PutOrganizationApisCltTestdefinitionsParams contains all the parameters to send to the API endpoint
for the put organization apis clt testdefinitions operation typically these are written to a http.Request
*/
type PutOrganizationApisCltTestdefinitionsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body*/
	Body *models.TestDefinition
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) WithTimeout(timeout time.Duration) *PutOrganizationApisCltTestdefinitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) WithContext(ctx context.Context) *PutOrganizationApisCltTestdefinitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) WithHTTPClient(client *http.Client) *PutOrganizationApisCltTestdefinitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) WithAPIVersion(aPIVersion string) *PutOrganizationApisCltTestdefinitionsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) WithBody(body *models.TestDefinition) *PutOrganizationApisCltTestdefinitionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) SetBody(body *models.TestDefinition) {
	o.Body = body
}

// WithOrganization adds the organization to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) WithOrganization(organization string) *PutOrganizationApisCltTestdefinitionsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the put organization apis clt testdefinitions params
func (o *PutOrganizationApisCltTestdefinitionsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrganizationApisCltTestdefinitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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