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

// NewDeleteRequestsRequestNameParams creates a new DeleteRequestsRequestNameParams object
// with the default values initialized.
func NewDeleteRequestsRequestNameParams() *DeleteRequestsRequestNameParams {
	var ()
	return &DeleteRequestsRequestNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRequestsRequestNameParamsWithTimeout creates a new DeleteRequestsRequestNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRequestsRequestNameParamsWithTimeout(timeout time.Duration) *DeleteRequestsRequestNameParams {
	var ()
	return &DeleteRequestsRequestNameParams{

		timeout: timeout,
	}
}

// NewDeleteRequestsRequestNameParamsWithContext creates a new DeleteRequestsRequestNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRequestsRequestNameParamsWithContext(ctx context.Context) *DeleteRequestsRequestNameParams {
	var ()
	return &DeleteRequestsRequestNameParams{

		Context: ctx,
	}
}

// NewDeleteRequestsRequestNameParamsWithHTTPClient creates a new DeleteRequestsRequestNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRequestsRequestNameParamsWithHTTPClient(client *http.Client) *DeleteRequestsRequestNameParams {
	var ()
	return &DeleteRequestsRequestNameParams{
		HTTPClient: client,
	}
}

/*DeleteRequestsRequestNameParams contains all the parameters to send to the API endpoint
for the delete requests request name operation typically these are written to a http.Request
*/
type DeleteRequestsRequestNameParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*RequestName
	  The symbol request name.

	*/
	RequestName string
	/*Synchronous
	  If true, delete all the debug entries under this request synchronously in the current session. If false, the deletion will be postponed to a later point and be executed automatically by the system.

	*/
	Synchronous *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) WithTimeout(timeout time.Duration) *DeleteRequestsRequestNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) WithContext(ctx context.Context) *DeleteRequestsRequestNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) WithHTTPClient(client *http.Client) *DeleteRequestsRequestNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) WithAPIVersion(aPIVersion string) *DeleteRequestsRequestNameParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) WithOrganization(organization string) *DeleteRequestsRequestNameParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithRequestName adds the requestName to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) WithRequestName(requestName string) *DeleteRequestsRequestNameParams {
	o.SetRequestName(requestName)
	return o
}

// SetRequestName adds the requestName to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) SetRequestName(requestName string) {
	o.RequestName = requestName
}

// WithSynchronous adds the synchronous to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) WithSynchronous(synchronous *bool) *DeleteRequestsRequestNameParams {
	o.SetSynchronous(synchronous)
	return o
}

// SetSynchronous adds the synchronous to the delete requests request name params
func (o *DeleteRequestsRequestNameParams) SetSynchronous(synchronous *bool) {
	o.Synchronous = synchronous
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRequestsRequestNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param requestName
	qrRequestName := o.RequestName
	qRequestName := qrRequestName
	if qRequestName != "" {
		if err := r.SetQueryParam("requestName", qRequestName); err != nil {
			return err
		}
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