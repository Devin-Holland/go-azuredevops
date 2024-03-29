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

// NewGetConsumerActionParams creates a new GetConsumerActionParams object
// with the default values initialized.
func NewGetConsumerActionParams() *GetConsumerActionParams {
	var ()
	return &GetConsumerActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConsumerActionParamsWithTimeout creates a new GetConsumerActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConsumerActionParamsWithTimeout(timeout time.Duration) *GetConsumerActionParams {
	var ()
	return &GetConsumerActionParams{

		timeout: timeout,
	}
}

// NewGetConsumerActionParamsWithContext creates a new GetConsumerActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConsumerActionParamsWithContext(ctx context.Context) *GetConsumerActionParams {
	var ()
	return &GetConsumerActionParams{

		Context: ctx,
	}
}

// NewGetConsumerActionParamsWithHTTPClient creates a new GetConsumerActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConsumerActionParamsWithHTTPClient(client *http.Client) *GetConsumerActionParams {
	var ()
	return &GetConsumerActionParams{
		HTTPClient: client,
	}
}

/*GetConsumerActionParams contains all the parameters to send to the API endpoint
for the get consumer action operation typically these are written to a http.Request
*/
type GetConsumerActionParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*ConsumerActionID
	  ID for a consumerActionId.

	*/
	ConsumerActionID string
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

// WithTimeout adds the timeout to the get consumer action params
func (o *GetConsumerActionParams) WithTimeout(timeout time.Duration) *GetConsumerActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get consumer action params
func (o *GetConsumerActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get consumer action params
func (o *GetConsumerActionParams) WithContext(ctx context.Context) *GetConsumerActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get consumer action params
func (o *GetConsumerActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get consumer action params
func (o *GetConsumerActionParams) WithHTTPClient(client *http.Client) *GetConsumerActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get consumer action params
func (o *GetConsumerActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get consumer action params
func (o *GetConsumerActionParams) WithAPIVersion(aPIVersion string) *GetConsumerActionParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get consumer action params
func (o *GetConsumerActionParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithConsumerActionID adds the consumerActionID to the get consumer action params
func (o *GetConsumerActionParams) WithConsumerActionID(consumerActionID string) *GetConsumerActionParams {
	o.SetConsumerActionID(consumerActionID)
	return o
}

// SetConsumerActionID adds the consumerActionId to the get consumer action params
func (o *GetConsumerActionParams) SetConsumerActionID(consumerActionID string) {
	o.ConsumerActionID = consumerActionID
}

// WithConsumerID adds the consumerID to the get consumer action params
func (o *GetConsumerActionParams) WithConsumerID(consumerID string) *GetConsumerActionParams {
	o.SetConsumerID(consumerID)
	return o
}

// SetConsumerID adds the consumerId to the get consumer action params
func (o *GetConsumerActionParams) SetConsumerID(consumerID string) {
	o.ConsumerID = consumerID
}

// WithOrganization adds the organization to the get consumer action params
func (o *GetConsumerActionParams) WithOrganization(organization string) *GetConsumerActionParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get consumer action params
func (o *GetConsumerActionParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPublisherID adds the publisherID to the get consumer action params
func (o *GetConsumerActionParams) WithPublisherID(publisherID *string) *GetConsumerActionParams {
	o.SetPublisherID(publisherID)
	return o
}

// SetPublisherID adds the publisherId to the get consumer action params
func (o *GetConsumerActionParams) SetPublisherID(publisherID *string) {
	o.PublisherID = publisherID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConsumerActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param consumerActionId
	if err := r.SetPathParam("consumerActionId", o.ConsumerActionID); err != nil {
		return err
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
