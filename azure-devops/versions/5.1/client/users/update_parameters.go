// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUpdateParams creates a new UpdateParams object
// with the default values initialized.
func NewUpdateParams() *UpdateParams {
	var ()
	return &UpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateParamsWithTimeout creates a new UpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateParamsWithTimeout(timeout time.Duration) *UpdateParams {
	var ()
	return &UpdateParams{

		timeout: timeout,
	}
}

// NewUpdateParamsWithContext creates a new UpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateParamsWithContext(ctx context.Context) *UpdateParams {
	var ()
	return &UpdateParams{

		Context: ctx,
	}
}

// NewUpdateParamsWithHTTPClient creates a new UpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateParamsWithHTTPClient(client *http.Client) *UpdateParams {
	var ()
	return &UpdateParams{
		HTTPClient: client,
	}
}

/*UpdateParams contains all the parameters to send to the API endpoint
for the update operation typically these are written to a http.Request
*/
type UpdateParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  The subset of the full graph user used to uniquely find the graph subject in an external provider.

	*/
	Body models.GraphUserUpdateContext
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*UserDescriptor
	  the descriptor of the user to update

	*/
	UserDescriptor string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update params
func (o *UpdateParams) WithTimeout(timeout time.Duration) *UpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update params
func (o *UpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update params
func (o *UpdateParams) WithContext(ctx context.Context) *UpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update params
func (o *UpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) WithHTTPClient(client *http.Client) *UpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update params
func (o *UpdateParams) WithAPIVersion(aPIVersion string) *UpdateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update params
func (o *UpdateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update params
func (o *UpdateParams) WithBody(body models.GraphUserUpdateContext) *UpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update params
func (o *UpdateParams) SetBody(body models.GraphUserUpdateContext) {
	o.Body = body
}

// WithOrganization adds the organization to the update params
func (o *UpdateParams) WithOrganization(organization string) *UpdateParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update params
func (o *UpdateParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithUserDescriptor adds the userDescriptor to the update params
func (o *UpdateParams) WithUserDescriptor(userDescriptor string) *UpdateParams {
	o.SetUserDescriptor(userDescriptor)
	return o
}

// SetUserDescriptor adds the userDescriptor to the update params
func (o *UpdateParams) SetUserDescriptor(userDescriptor string) {
	o.UserDescriptor = userDescriptor
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userDescriptor
	if err := r.SetPathParam("userDescriptor", o.UserDescriptor); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
