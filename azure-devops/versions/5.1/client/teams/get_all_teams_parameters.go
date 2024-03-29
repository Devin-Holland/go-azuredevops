// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewGetAllTeamsParams creates a new GetAllTeamsParams object
// with the default values initialized.
func NewGetAllTeamsParams() *GetAllTeamsParams {
	var ()
	return &GetAllTeamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllTeamsParamsWithTimeout creates a new GetAllTeamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllTeamsParamsWithTimeout(timeout time.Duration) *GetAllTeamsParams {
	var ()
	return &GetAllTeamsParams{

		timeout: timeout,
	}
}

// NewGetAllTeamsParamsWithContext creates a new GetAllTeamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllTeamsParamsWithContext(ctx context.Context) *GetAllTeamsParams {
	var ()
	return &GetAllTeamsParams{

		Context: ctx,
	}
}

// NewGetAllTeamsParamsWithHTTPClient creates a new GetAllTeamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllTeamsParamsWithHTTPClient(client *http.Client) *GetAllTeamsParams {
	var ()
	return &GetAllTeamsParams{
		HTTPClient: client,
	}
}

/*GetAllTeamsParams contains all the parameters to send to the API endpoint
for the get all teams operation typically these are written to a http.Request
*/
type GetAllTeamsParams struct {

	/*NrDollarExpandIdentity
	  A value indicating whether or not to expand Identity information in the result WebApiTeam object.

	*/
	DollarExpandIdentity *bool
	/*NrDollarMine
	  If true, then return all teams requesting user is member. Otherwise return all teams user has read access.

	*/
	DollarMine *bool
	/*NrDollarSkip
	  Number of teams to skip.

	*/
	DollarSkip *int32
	/*NrDollarTop
	  Maximum number of teams to return.

	*/
	DollarTop *int32
	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.3' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all teams params
func (o *GetAllTeamsParams) WithTimeout(timeout time.Duration) *GetAllTeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all teams params
func (o *GetAllTeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all teams params
func (o *GetAllTeamsParams) WithContext(ctx context.Context) *GetAllTeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all teams params
func (o *GetAllTeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all teams params
func (o *GetAllTeamsParams) WithHTTPClient(client *http.Client) *GetAllTeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all teams params
func (o *GetAllTeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarExpandIdentity adds the dollarExpandIdentity to the get all teams params
func (o *GetAllTeamsParams) WithDollarExpandIdentity(dollarExpandIdentity *bool) *GetAllTeamsParams {
	o.SetDollarExpandIdentity(dollarExpandIdentity)
	return o
}

// SetDollarExpandIdentity adds the dollarExpandIdentity to the get all teams params
func (o *GetAllTeamsParams) SetDollarExpandIdentity(dollarExpandIdentity *bool) {
	o.DollarExpandIdentity = dollarExpandIdentity
}

// WithDollarMine adds the dollarMine to the get all teams params
func (o *GetAllTeamsParams) WithDollarMine(dollarMine *bool) *GetAllTeamsParams {
	o.SetDollarMine(dollarMine)
	return o
}

// SetDollarMine adds the dollarMine to the get all teams params
func (o *GetAllTeamsParams) SetDollarMine(dollarMine *bool) {
	o.DollarMine = dollarMine
}

// WithDollarSkip adds the dollarSkip to the get all teams params
func (o *GetAllTeamsParams) WithDollarSkip(dollarSkip *int32) *GetAllTeamsParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get all teams params
func (o *GetAllTeamsParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get all teams params
func (o *GetAllTeamsParams) WithDollarTop(dollarTop *int32) *GetAllTeamsParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get all teams params
func (o *GetAllTeamsParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get all teams params
func (o *GetAllTeamsParams) WithAPIVersion(aPIVersion string) *GetAllTeamsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get all teams params
func (o *GetAllTeamsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the get all teams params
func (o *GetAllTeamsParams) WithOrganization(organization string) *GetAllTeamsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get all teams params
func (o *GetAllTeamsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarExpandIdentity != nil {

		// query param $expandIdentity
		var qrNrDollarExpandIdentity bool
		if o.DollarExpandIdentity != nil {
			qrNrDollarExpandIdentity = *o.DollarExpandIdentity
		}
		qNrDollarExpandIdentity := swag.FormatBool(qrNrDollarExpandIdentity)
		if qNrDollarExpandIdentity != "" {
			if err := r.SetQueryParam("$expandIdentity", qNrDollarExpandIdentity); err != nil {
				return err
			}
		}

	}

	if o.DollarMine != nil {

		// query param $mine
		var qrNrDollarMine bool
		if o.DollarMine != nil {
			qrNrDollarMine = *o.DollarMine
		}
		qNrDollarMine := swag.FormatBool(qrNrDollarMine)
		if qNrDollarMine != "" {
			if err := r.SetQueryParam("$mine", qNrDollarMine); err != nil {
				return err
			}
		}

	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrNrDollarSkip int32
		if o.DollarSkip != nil {
			qrNrDollarSkip = *o.DollarSkip
		}
		qNrDollarSkip := swag.FormatInt32(qrNrDollarSkip)
		if qNrDollarSkip != "" {
			if err := r.SetQueryParam("$skip", qNrDollarSkip); err != nil {
				return err
			}
		}

	}

	if o.DollarTop != nil {

		// query param $top
		var qrNrDollarTop int32
		if o.DollarTop != nil {
			qrNrDollarTop = *o.DollarTop
		}
		qNrDollarTop := swag.FormatInt32(qrNrDollarTop)
		if qNrDollarTop != "" {
			if err := r.SetQueryParam("$top", qNrDollarTop); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
