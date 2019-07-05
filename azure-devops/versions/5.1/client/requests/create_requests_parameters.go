// Code generated by go-swagger; DO NOT EDIT.

package requests

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

// NewCreateRequestsParams creates a new CreateRequestsParams object
// with the default values initialized.
func NewCreateRequestsParams() *CreateRequestsParams {
	var ()
	return &CreateRequestsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRequestsParamsWithTimeout creates a new CreateRequestsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRequestsParamsWithTimeout(timeout time.Duration) *CreateRequestsParams {
	var ()
	return &CreateRequestsParams{

		timeout: timeout,
	}
}

// NewCreateRequestsParamsWithContext creates a new CreateRequestsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRequestsParamsWithContext(ctx context.Context) *CreateRequestsParams {
	var ()
	return &CreateRequestsParams{

		Context: ctx,
	}
}

// NewCreateRequestsParamsWithHTTPClient creates a new CreateRequestsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRequestsParamsWithHTTPClient(client *http.Client) *CreateRequestsParams {
	var ()
	return &CreateRequestsParams{
		HTTPClient: client,
	}
}

/*CreateRequestsParams contains all the parameters to send to the API endpoint
for the create requests operation typically these are written to a http.Request
*/
type CreateRequestsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  The symbol request to create.

	*/
	Body *models.Request
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create requests params
func (o *CreateRequestsParams) WithTimeout(timeout time.Duration) *CreateRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create requests params
func (o *CreateRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create requests params
func (o *CreateRequestsParams) WithContext(ctx context.Context) *CreateRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create requests params
func (o *CreateRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create requests params
func (o *CreateRequestsParams) WithHTTPClient(client *http.Client) *CreateRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create requests params
func (o *CreateRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create requests params
func (o *CreateRequestsParams) WithAPIVersion(aPIVersion string) *CreateRequestsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create requests params
func (o *CreateRequestsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create requests params
func (o *CreateRequestsParams) WithBody(body *models.Request) *CreateRequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create requests params
func (o *CreateRequestsParams) SetBody(body *models.Request) {
	o.Body = body
}

// WithOrganization adds the organization to the create requests params
func (o *CreateRequestsParams) WithOrganization(organization string) *CreateRequestsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the create requests params
func (o *CreateRequestsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
