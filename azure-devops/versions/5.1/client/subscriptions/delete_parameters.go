// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// NewDeleteParams creates a new DeleteParams object
// with the default values initialized.
func NewDeleteParams() *DeleteParams {
	var ()
	return &DeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteParamsWithTimeout creates a new DeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteParamsWithTimeout(timeout time.Duration) *DeleteParams {
	var ()
	return &DeleteParams{

		timeout: timeout,
	}
}

// NewDeleteParamsWithContext creates a new DeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteParamsWithContext(ctx context.Context) *DeleteParams {
	var ()
	return &DeleteParams{

		Context: ctx,
	}
}

// NewDeleteParamsWithHTTPClient creates a new DeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteParamsWithHTTPClient(client *http.Client) *DeleteParams {
	var ()
	return &DeleteParams{
		HTTPClient: client,
	}
}

/*DeleteParams contains all the parameters to send to the API endpoint
for the delete operation typically these are written to a http.Request
*/
type DeleteParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*SubscriptionID
	  ID for a subscription.

	*/
	SubscriptionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete params
func (o *DeleteParams) WithTimeout(timeout time.Duration) *DeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete params
func (o *DeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete params
func (o *DeleteParams) WithContext(ctx context.Context) *DeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete params
func (o *DeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete params
func (o *DeleteParams) WithHTTPClient(client *http.Client) *DeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete params
func (o *DeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete params
func (o *DeleteParams) WithAPIVersion(aPIVersion string) *DeleteParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete params
func (o *DeleteParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the delete params
func (o *DeleteParams) WithOrganization(organization string) *DeleteParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete params
func (o *DeleteParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithSubscriptionID adds the subscriptionID to the delete params
func (o *DeleteParams) WithSubscriptionID(subscriptionID strfmt.UUID) *DeleteParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the delete params
func (o *DeleteParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
