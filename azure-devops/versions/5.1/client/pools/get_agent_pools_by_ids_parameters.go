// Code generated by go-swagger; DO NOT EDIT.

package pools

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

// NewGetAgentPoolsByIdsParams creates a new GetAgentPoolsByIdsParams object
// with the default values initialized.
func NewGetAgentPoolsByIdsParams() *GetAgentPoolsByIdsParams {
	var ()
	return &GetAgentPoolsByIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentPoolsByIdsParamsWithTimeout creates a new GetAgentPoolsByIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAgentPoolsByIdsParamsWithTimeout(timeout time.Duration) *GetAgentPoolsByIdsParams {
	var ()
	return &GetAgentPoolsByIdsParams{

		timeout: timeout,
	}
}

// NewGetAgentPoolsByIdsParamsWithContext creates a new GetAgentPoolsByIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAgentPoolsByIdsParamsWithContext(ctx context.Context) *GetAgentPoolsByIdsParams {
	var ()
	return &GetAgentPoolsByIdsParams{

		Context: ctx,
	}
}

// NewGetAgentPoolsByIdsParamsWithHTTPClient creates a new GetAgentPoolsByIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAgentPoolsByIdsParamsWithHTTPClient(client *http.Client) *GetAgentPoolsByIdsParams {
	var ()
	return &GetAgentPoolsByIdsParams{
		HTTPClient: client,
	}
}

/*GetAgentPoolsByIdsParams contains all the parameters to send to the API endpoint
for the get agent pools by ids operation typically these are written to a http.Request
*/
type GetAgentPoolsByIdsParams struct {

	/*ActionFilter
	  Filter by whether the calling user has use or manage permissions

	*/
	ActionFilter *string
	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PoolIds
	  pool Ids to fetch

	*/
	PoolIds string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) WithTimeout(timeout time.Duration) *GetAgentPoolsByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) WithContext(ctx context.Context) *GetAgentPoolsByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) WithHTTPClient(client *http.Client) *GetAgentPoolsByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionFilter adds the actionFilter to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) WithActionFilter(actionFilter *string) *GetAgentPoolsByIdsParams {
	o.SetActionFilter(actionFilter)
	return o
}

// SetActionFilter adds the actionFilter to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) SetActionFilter(actionFilter *string) {
	o.ActionFilter = actionFilter
}

// WithAPIVersion adds the aPIVersion to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) WithAPIVersion(aPIVersion string) *GetAgentPoolsByIdsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) WithOrganization(organization string) *GetAgentPoolsByIdsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPoolIds adds the poolIds to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) WithPoolIds(poolIds string) *GetAgentPoolsByIdsParams {
	o.SetPoolIds(poolIds)
	return o
}

// SetPoolIds adds the poolIds to the get agent pools by ids params
func (o *GetAgentPoolsByIdsParams) SetPoolIds(poolIds string) {
	o.PoolIds = poolIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentPoolsByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionFilter != nil {

		// query param actionFilter
		var qrActionFilter string
		if o.ActionFilter != nil {
			qrActionFilter = *o.ActionFilter
		}
		qActionFilter := qrActionFilter
		if qActionFilter != "" {
			if err := r.SetQueryParam("actionFilter", qActionFilter); err != nil {
				return err
			}
		}

	}

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

	// query param poolIds
	qrPoolIds := o.PoolIds
	qPoolIds := qrPoolIds
	if qPoolIds != "" {
		if err := r.SetQueryParam("poolIds", qPoolIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}