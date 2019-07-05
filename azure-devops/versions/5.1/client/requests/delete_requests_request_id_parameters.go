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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteRequestsRequestIDParams creates a new DeleteRequestsRequestIDParams object
// with the default values initialized.
func NewDeleteRequestsRequestIDParams() *DeleteRequestsRequestIDParams {
	var ()
	return &DeleteRequestsRequestIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRequestsRequestIDParamsWithTimeout creates a new DeleteRequestsRequestIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRequestsRequestIDParamsWithTimeout(timeout time.Duration) *DeleteRequestsRequestIDParams {
	var ()
	return &DeleteRequestsRequestIDParams{

		timeout: timeout,
	}
}

// NewDeleteRequestsRequestIDParamsWithContext creates a new DeleteRequestsRequestIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRequestsRequestIDParamsWithContext(ctx context.Context) *DeleteRequestsRequestIDParams {
	var ()
	return &DeleteRequestsRequestIDParams{

		Context: ctx,
	}
}

// NewDeleteRequestsRequestIDParamsWithHTTPClient creates a new DeleteRequestsRequestIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRequestsRequestIDParamsWithHTTPClient(client *http.Client) *DeleteRequestsRequestIDParams {
	var ()
	return &DeleteRequestsRequestIDParams{
		HTTPClient: client,
	}
}

/*DeleteRequestsRequestIDParams contains all the parameters to send to the API endpoint
for the delete requests request Id operation typically these are written to a http.Request
*/
type DeleteRequestsRequestIDParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*RequestID
	  The symbol request identifier.

	*/
	RequestID string
	/*Synchronous
	  If true, delete all the debug entries under this request synchronously in the current session. If false, the deletion will be postponed to a later point and be executed automatically by the system.

	*/
	Synchronous *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) WithTimeout(timeout time.Duration) *DeleteRequestsRequestIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) WithContext(ctx context.Context) *DeleteRequestsRequestIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) WithHTTPClient(client *http.Client) *DeleteRequestsRequestIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) WithAPIVersion(aPIVersion string) *DeleteRequestsRequestIDParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) WithOrganization(organization string) *DeleteRequestsRequestIDParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithRequestID adds the requestID to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) WithRequestID(requestID string) *DeleteRequestsRequestIDParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WithSynchronous adds the synchronous to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) WithSynchronous(synchronous *bool) *DeleteRequestsRequestIDParams {
	o.SetSynchronous(synchronous)
	return o
}

// SetSynchronous adds the synchronous to the delete requests request Id params
func (o *DeleteRequestsRequestIDParams) SetSynchronous(synchronous *bool) {
	o.Synchronous = synchronous
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRequestsRequestIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID); err != nil {
		return err
	}

	if o.Synchronous != nil {

		// query param synchronous
		var qrSynchronous bool
		if o.Synchronous != nil {
			qrSynchronous = *o.Synchronous
		}
		qSynchronous := swag.FormatBool(qrSynchronous)
		if qSynchronous != "" {
			if err := r.SetQueryParam("synchronous", qSynchronous); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
