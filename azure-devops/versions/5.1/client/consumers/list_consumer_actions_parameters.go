// Code generated by go-swagger; DO NOT EDIT.

package consumers

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

// NewListConsumerActionsParams creates a new ListConsumerActionsParams object
// with the default values initialized.
func NewListConsumerActionsParams() *ListConsumerActionsParams {
	var ()
	return &ListConsumerActionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListConsumerActionsParamsWithTimeout creates a new ListConsumerActionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListConsumerActionsParamsWithTimeout(timeout time.Duration) *ListConsumerActionsParams {
	var ()
	return &ListConsumerActionsParams{

		timeout: timeout,
	}
}

// NewListConsumerActionsParamsWithContext creates a new ListConsumerActionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListConsumerActionsParamsWithContext(ctx context.Context) *ListConsumerActionsParams {
	var ()
	return &ListConsumerActionsParams{

		Context: ctx,
	}
}

// NewListConsumerActionsParamsWithHTTPClient creates a new ListConsumerActionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListConsumerActionsParamsWithHTTPClient(client *http.Client) *ListConsumerActionsParams {
	var ()
	return &ListConsumerActionsParams{
		HTTPClient: client,
	}
}

/*ListConsumerActionsParams contains all the parameters to send to the API endpoint
for the list consumer actions operation typically these are written to a http.Request
*/
type ListConsumerActionsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*ConsumerID
	  ID for a consumer.

	*/
	ConsumerID string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PublisherID*/
	PublisherID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list consumer actions params
func (o *ListConsumerActionsParams) WithTimeout(timeout time.Duration) *ListConsumerActionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list consumer actions params
func (o *ListConsumerActionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list consumer actions params
func (o *ListConsumerActionsParams) WithContext(ctx context.Context) *ListConsumerActionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list consumer actions params
func (o *ListConsumerActionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list consumer actions params
func (o *ListConsumerActionsParams) WithHTTPClient(client *http.Client) *ListConsumerActionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list consumer actions params
func (o *ListConsumerActionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the list consumer actions params
func (o *ListConsumerActionsParams) WithAPIVersion(aPIVersion string) *ListConsumerActionsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list consumer actions params
func (o *ListConsumerActionsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithConsumerID adds the consumerID to the list consumer actions params
func (o *ListConsumerActionsParams) WithConsumerID(consumerID string) *ListConsumerActionsParams {
	o.SetConsumerID(consumerID)
	return o
}

// SetConsumerID adds the consumerId to the list consumer actions params
func (o *ListConsumerActionsParams) SetConsumerID(consumerID string) {
	o.ConsumerID = consumerID
}

// WithOrganization adds the organization to the list consumer actions params
func (o *ListConsumerActionsParams) WithOrganization(organization string) *ListConsumerActionsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the list consumer actions params
func (o *ListConsumerActionsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPublisherID adds the publisherID to the list consumer actions params
func (o *ListConsumerActionsParams) WithPublisherID(publisherID *string) *ListConsumerActionsParams {
	o.SetPublisherID(publisherID)
	return o
}

// SetPublisherID adds the publisherId to the list consumer actions params
func (o *ListConsumerActionsParams) SetPublisherID(publisherID *string) {
	o.PublisherID = publisherID
}

// WriteToRequest writes these params to a swagger request
func (o *ListConsumerActionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param consumerId
	if err := r.SetPathParam("consumerId", o.ConsumerID); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	if o.PublisherID != nil {

		// query param publisherId
		var qrPublisherID string
		if o.PublisherID != nil {
			qrPublisherID = *o.PublisherID
		}
		qPublisherID := qrPublisherID
		if qPublisherID != "" {
			if err := r.SetQueryParam("publisherId", qPublisherID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}