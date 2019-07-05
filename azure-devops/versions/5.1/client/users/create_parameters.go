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

// NewCreateParams creates a new CreateParams object
// with the default values initialized.
func NewCreateParams() *CreateParams {
	var ()
	return &CreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateParamsWithTimeout creates a new CreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateParamsWithTimeout(timeout time.Duration) *CreateParams {
	var ()
	return &CreateParams{

		timeout: timeout,
	}
}

// NewCreateParamsWithContext creates a new CreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateParamsWithContext(ctx context.Context) *CreateParams {
	var ()
	return &CreateParams{

		Context: ctx,
	}
}

// NewCreateParamsWithHTTPClient creates a new CreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateParamsWithHTTPClient(client *http.Client) *CreateParams {
	var ()
	return &CreateParams{
		HTTPClient: client,
	}
}

/*CreateParams contains all the parameters to send to the API endpoint
for the create operation typically these are written to a http.Request
*/
type CreateParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  The subset of the full graph user used to uniquely find the graph subject in an external provider.

	*/
	Body *models.GraphUserCreationContext
	/*GroupDescriptors
	  A comma separated list of descriptors of groups you want the graph user to join

	*/
	GroupDescriptors *string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create params
func (o *CreateParams) WithTimeout(timeout time.Duration) *CreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create params
func (o *CreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create params
func (o *CreateParams) WithContext(ctx context.Context) *CreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create params
func (o *CreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create params
func (o *CreateParams) WithHTTPClient(client *http.Client) *CreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create params
func (o *CreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create params
func (o *CreateParams) WithAPIVersion(aPIVersion string) *CreateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create params
func (o *CreateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create params
func (o *CreateParams) WithBody(body *models.GraphUserCreationContext) *CreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create params
func (o *CreateParams) SetBody(body *models.GraphUserCreationContext) {
	o.Body = body
}

// WithGroupDescriptors adds the groupDescriptors to the create params
func (o *CreateParams) WithGroupDescriptors(groupDescriptors *string) *CreateParams {
	o.SetGroupDescriptors(groupDescriptors)
	return o
}

// SetGroupDescriptors adds the groupDescriptors to the create params
func (o *CreateParams) SetGroupDescriptors(groupDescriptors *string) {
	o.GroupDescriptors = groupDescriptors
}

// WithOrganization adds the organization to the create params
func (o *CreateParams) WithOrganization(organization string) *CreateParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the create params
func (o *CreateParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *CreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.GroupDescriptors != nil {

		// query param groupDescriptors
		var qrGroupDescriptors string
		if o.GroupDescriptors != nil {
			qrGroupDescriptors = *o.GroupDescriptors
		}
		qGroupDescriptors := qrGroupDescriptors
		if qGroupDescriptors != "" {
			if err := r.SetQueryParam("groupDescriptors", qGroupDescriptors); err != nil {
				return err
			}
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
