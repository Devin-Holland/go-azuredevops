// Code generated by go-swagger; DO NOT EDIT.

package npm

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

// NewGetScopedPackageVersionFromRecycleBinParams creates a new GetScopedPackageVersionFromRecycleBinParams object
// with the default values initialized.
func NewGetScopedPackageVersionFromRecycleBinParams() *GetScopedPackageVersionFromRecycleBinParams {
	var ()
	return &GetScopedPackageVersionFromRecycleBinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScopedPackageVersionFromRecycleBinParamsWithTimeout creates a new GetScopedPackageVersionFromRecycleBinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScopedPackageVersionFromRecycleBinParamsWithTimeout(timeout time.Duration) *GetScopedPackageVersionFromRecycleBinParams {
	var ()
	return &GetScopedPackageVersionFromRecycleBinParams{

		timeout: timeout,
	}
}

// NewGetScopedPackageVersionFromRecycleBinParamsWithContext creates a new GetScopedPackageVersionFromRecycleBinParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScopedPackageVersionFromRecycleBinParamsWithContext(ctx context.Context) *GetScopedPackageVersionFromRecycleBinParams {
	var ()
	return &GetScopedPackageVersionFromRecycleBinParams{

		Context: ctx,
	}
}

// NewGetScopedPackageVersionFromRecycleBinParamsWithHTTPClient creates a new GetScopedPackageVersionFromRecycleBinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScopedPackageVersionFromRecycleBinParamsWithHTTPClient(client *http.Client) *GetScopedPackageVersionFromRecycleBinParams {
	var ()
	return &GetScopedPackageVersionFromRecycleBinParams{
		HTTPClient: client,
	}
}

/*GetScopedPackageVersionFromRecycleBinParams contains all the parameters to send to the API endpoint
for the get scoped package version from recycle bin operation typically these are written to a http.Request
*/
type GetScopedPackageVersionFromRecycleBinParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*FeedID
	  Name or ID of the feed.

	*/
	FeedID string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PackageScope
	  Scope of the package (the 'scope' part of @scope/name)

	*/
	PackageScope string
	/*PackageVersion
	  Version of the package.

	*/
	PackageVersion string
	/*UnscopedPackageName
	  Name of the package (the 'name' part of @scope/name).

	*/
	UnscopedPackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) WithTimeout(timeout time.Duration) *GetScopedPackageVersionFromRecycleBinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) WithContext(ctx context.Context) *GetScopedPackageVersionFromRecycleBinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) WithHTTPClient(client *http.Client) *GetScopedPackageVersionFromRecycleBinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) WithAPIVersion(aPIVersion string) *GetScopedPackageVersionFromRecycleBinParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithFeedID adds the feedID to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) WithFeedID(feedID string) *GetScopedPackageVersionFromRecycleBinParams {
	o.SetFeedID(feedID)
	return o
}

// SetFeedID adds the feedId to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) SetFeedID(feedID string) {
	o.FeedID = feedID
}

// WithOrganization adds the organization to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) WithOrganization(organization string) *GetScopedPackageVersionFromRecycleBinParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPackageScope adds the packageScope to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) WithPackageScope(packageScope string) *GetScopedPackageVersionFromRecycleBinParams {
	o.SetPackageScope(packageScope)
	return o
}

// SetPackageScope adds the packageScope to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) SetPackageScope(packageScope string) {
	o.PackageScope = packageScope
}

// WithPackageVersion adds the packageVersion to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) WithPackageVersion(packageVersion string) *GetScopedPackageVersionFromRecycleBinParams {
	o.SetPackageVersion(packageVersion)
	return o
}

// SetPackageVersion adds the packageVersion to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) SetPackageVersion(packageVersion string) {
	o.PackageVersion = packageVersion
}

// WithUnscopedPackageName adds the unscopedPackageName to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) WithUnscopedPackageName(unscopedPackageName string) *GetScopedPackageVersionFromRecycleBinParams {
	o.SetUnscopedPackageName(unscopedPackageName)
	return o
}

// SetUnscopedPackageName adds the unscopedPackageName to the get scoped package version from recycle bin params
func (o *GetScopedPackageVersionFromRecycleBinParams) SetUnscopedPackageName(unscopedPackageName string) {
	o.UnscopedPackageName = unscopedPackageName
}

// WriteToRequest writes these params to a swagger request
func (o *GetScopedPackageVersionFromRecycleBinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param feedId
	if err := r.SetPathParam("feedId", o.FeedID); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param packageScope
	if err := r.SetPathParam("packageScope", o.PackageScope); err != nil {
		return err
	}

	// path param packageVersion
	if err := r.SetPathParam("packageVersion", o.PackageVersion); err != nil {
		return err
	}

	// path param unscopedPackageName
	if err := r.SetPathParam("unscopedPackageName", o.UnscopedPackageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}