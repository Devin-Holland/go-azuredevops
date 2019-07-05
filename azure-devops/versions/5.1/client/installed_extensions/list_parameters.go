// Code generated by go-swagger; DO NOT EDIT.

package installed_extensions

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

// NewListParams creates a new ListParams object
// with the default values initialized.
func NewListParams() *ListParams {
	var ()
	return &ListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListParamsWithTimeout creates a new ListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListParamsWithTimeout(timeout time.Duration) *ListParams {
	var ()
	return &ListParams{

		timeout: timeout,
	}
}

// NewListParamsWithContext creates a new ListParams object
// with the default values initialized, and the ability to set a context for a request
func NewListParamsWithContext(ctx context.Context) *ListParams {
	var ()
	return &ListParams{

		Context: ctx,
	}
}

// NewListParamsWithHTTPClient creates a new ListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListParamsWithHTTPClient(client *http.Client) *ListParams {
	var ()
	return &ListParams{
		HTTPClient: client,
	}
}

/*ListParams contains all the parameters to send to the API endpoint
for the list operation typically these are written to a http.Request
*/
type ListParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*AssetTypes*/
	AssetTypes *string
	/*IncludeDisabledExtensions
	  If true (the default), include disabled extensions in the results.

	*/
	IncludeDisabledExtensions *bool
	/*IncludeErrors
	  If true, include installed extensions with errors.

	*/
	IncludeErrors *bool
	/*IncludeInstallationIssues*/
	IncludeInstallationIssues *bool
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list params
func (o *ListParams) WithTimeout(timeout time.Duration) *ListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list params
func (o *ListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list params
func (o *ListParams) WithContext(ctx context.Context) *ListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list params
func (o *ListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list params
func (o *ListParams) WithHTTPClient(client *http.Client) *ListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list params
func (o *ListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the list params
func (o *ListParams) WithAPIVersion(aPIVersion string) *ListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list params
func (o *ListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithAssetTypes adds the assetTypes to the list params
func (o *ListParams) WithAssetTypes(assetTypes *string) *ListParams {
	o.SetAssetTypes(assetTypes)
	return o
}

// SetAssetTypes adds the assetTypes to the list params
func (o *ListParams) SetAssetTypes(assetTypes *string) {
	o.AssetTypes = assetTypes
}

// WithIncludeDisabledExtensions adds the includeDisabledExtensions to the list params
func (o *ListParams) WithIncludeDisabledExtensions(includeDisabledExtensions *bool) *ListParams {
	o.SetIncludeDisabledExtensions(includeDisabledExtensions)
	return o
}

// SetIncludeDisabledExtensions adds the includeDisabledExtensions to the list params
func (o *ListParams) SetIncludeDisabledExtensions(includeDisabledExtensions *bool) {
	o.IncludeDisabledExtensions = includeDisabledExtensions
}

// WithIncludeErrors adds the includeErrors to the list params
func (o *ListParams) WithIncludeErrors(includeErrors *bool) *ListParams {
	o.SetIncludeErrors(includeErrors)
	return o
}

// SetIncludeErrors adds the includeErrors to the list params
func (o *ListParams) SetIncludeErrors(includeErrors *bool) {
	o.IncludeErrors = includeErrors
}

// WithIncludeInstallationIssues adds the includeInstallationIssues to the list params
func (o *ListParams) WithIncludeInstallationIssues(includeInstallationIssues *bool) *ListParams {
	o.SetIncludeInstallationIssues(includeInstallationIssues)
	return o
}

// SetIncludeInstallationIssues adds the includeInstallationIssues to the list params
func (o *ListParams) SetIncludeInstallationIssues(includeInstallationIssues *bool) {
	o.IncludeInstallationIssues = includeInstallationIssues
}

// WithOrganization adds the organization to the list params
func (o *ListParams) WithOrganization(organization string) *ListParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the list params
func (o *ListParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *ListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AssetTypes != nil {

		// query param assetTypes
		var qrAssetTypes string
		if o.AssetTypes != nil {
			qrAssetTypes = *o.AssetTypes
		}
		qAssetTypes := qrAssetTypes
		if qAssetTypes != "" {
			if err := r.SetQueryParam("assetTypes", qAssetTypes); err != nil {
				return err
			}
		}

	}

	if o.IncludeDisabledExtensions != nil {

		// query param includeDisabledExtensions
		var qrIncludeDisabledExtensions bool
		if o.IncludeDisabledExtensions != nil {
			qrIncludeDisabledExtensions = *o.IncludeDisabledExtensions
		}
		qIncludeDisabledExtensions := swag.FormatBool(qrIncludeDisabledExtensions)
		if qIncludeDisabledExtensions != "" {
			if err := r.SetQueryParam("includeDisabledExtensions", qIncludeDisabledExtensions); err != nil {
				return err
			}
		}

	}

	if o.IncludeErrors != nil {

		// query param includeErrors
		var qrIncludeErrors bool
		if o.IncludeErrors != nil {
			qrIncludeErrors = *o.IncludeErrors
		}
		qIncludeErrors := swag.FormatBool(qrIncludeErrors)
		if qIncludeErrors != "" {
			if err := r.SetQueryParam("includeErrors", qIncludeErrors); err != nil {
				return err
			}
		}

	}

	if o.IncludeInstallationIssues != nil {

		// query param includeInstallationIssues
		var qrIncludeInstallationIssues bool
		if o.IncludeInstallationIssues != nil {
			qrIncludeInstallationIssues = *o.IncludeInstallationIssues
		}
		qIncludeInstallationIssues := swag.FormatBool(qrIncludeInstallationIssues)
		if qIncludeInstallationIssues != "" {
			if err := r.SetQueryParam("includeInstallationIssues", qIncludeInstallationIssues); err != nil {
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
