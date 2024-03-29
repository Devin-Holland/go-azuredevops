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

// NewUpdateRequestsRequestNameParams creates a new UpdateRequestsRequestNameParams object
// with the default values initialized.
func NewUpdateRequestsRequestNameParams() *UpdateRequestsRequestNameParams {
	var ()
	return &UpdateRequestsRequestNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRequestsRequestNameParamsWithTimeout creates a new UpdateRequestsRequestNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRequestsRequestNameParamsWithTimeout(timeout time.Duration) *UpdateRequestsRequestNameParams {
	var ()
	return &UpdateRequestsRequestNameParams{

		timeout: timeout,
	}
}

// NewUpdateRequestsRequestNameParamsWithContext creates a new UpdateRequestsRequestNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRequestsRequestNameParamsWithContext(ctx context.Context) *UpdateRequestsRequestNameParams {
	var ()
	return &UpdateRequestsRequestNameParams{

		Context: ctx,
	}
}

// NewUpdateRequestsRequestNameParamsWithHTTPClient creates a new UpdateRequestsRequestNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRequestsRequestNameParamsWithHTTPClient(client *http.Client) *UpdateRequestsRequestNameParams {
	var ()
	return &UpdateRequestsRequestNameParams{
		HTTPClient: client,
	}
}

/*UpdateRequestsRequestNameParams contains all the parameters to send to the API endpoint
for the update requests request name operation typically these are written to a http.Request
*/
type UpdateRequestsRequestNameParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  The symbol request.

	*/
	Body *models.Request
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*RequestName
	  The symbol request name.

	*/
	RequestName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update requests request name params
func (o *UpdateRequestsRequestNameParams) WithTimeout(timeout time.Duration) *UpdateRequestsRequestNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update requests request name params
func (o *UpdateRequestsRequestNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update requests request name params
func (o *UpdateRequestsRequestNameParams) WithContext(ctx context.Context) *UpdateRequestsRequestNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update requests request name params
func (o *UpdateRequestsRequestNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update requests request name params
func (o *UpdateRequestsRequestNameParams) WithHTTPClient(client *http.Client) *UpdateRequestsRequestNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update requests request name params
func (o *UpdateRequestsRequestNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update requests request name params
func (o *UpdateRequestsRequestNameParams) WithAPIVersion(aPIVersion string) *UpdateRequestsRequestNameParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update requests request name params
func (o *UpdateRequestsRequestNameParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update requests request name params
func (o *UpdateRequestsRequestNameParams) WithBody(body *models.Request) *UpdateRequestsRequestNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update requests request name params
func (o *UpdateRequestsRequestNameParams) SetBody(body *models.Request) {
	o.Body = body
}

// WithOrganization adds the organization to the update requests request name params
func (o *UpdateRequestsRequestNameParams) WithOrganization(organization string) *UpdateRequestsRequestNameParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update requests request name params
func (o *UpdateRequestsRequestNameParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithRequestName adds the requestName to the update requests request name params
func (o *UpdateRequestsRequestNameParams) WithRequestName(requestName string) *UpdateRequestsRequestNameParams {
	o.SetRequestName(requestName)
	return o
}

// SetRequestName adds the requestName to the update requests request name params
func (o *UpdateRequestsRequestNameParams) SetRequestName(requestName string) {
	o.RequestName = requestName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRequestsRequestNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param requestName
	qrRequestName := o.RequestName
	qRequestName := qrRequestName
	if qRequestName != "" {
		if err := r.SetQueryParam("requestName", qRequestName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
