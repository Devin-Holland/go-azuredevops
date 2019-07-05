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

// NewGetScopedPackageVersionParams creates a new GetScopedPackageVersionParams object
// with the default values initialized.
func NewGetScopedPackageVersionParams() *GetScopedPackageVersionParams {
	var ()
	return &GetScopedPackageVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScopedPackageVersionParamsWithTimeout creates a new GetScopedPackageVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScopedPackageVersionParamsWithTimeout(timeout time.Duration) *GetScopedPackageVersionParams {
	var ()
	return &GetScopedPackageVersionParams{

		timeout: timeout,
	}
}

// NewGetScopedPackageVersionParamsWithContext creates a new GetScopedPackageVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScopedPackageVersionParamsWithContext(ctx context.Context) *GetScopedPackageVersionParams {
	var ()
	return &GetScopedPackageVersionParams{

		Context: ctx,
	}
}

// NewGetScopedPackageVersionParamsWithHTTPClient creates a new GetScopedPackageVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScopedPackageVersionParamsWithHTTPClient(client *http.Client) *GetScopedPackageVersionParams {
	var ()
	return &GetScopedPackageVersionParams{
		HTTPClient: client,
	}
}

/*GetScopedPackageVersionParams contains all the parameters to send to the API endpoint
for the get scoped package version operation typically these are written to a http.Request
*/
type GetScopedPackageVersionParams struct {

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
	  Scope of the package (the 'scope' part of @scope/name).

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

// WithTimeout adds the timeout to the get scoped package version params
func (o *GetScopedPackageVersionParams) WithTimeout(timeout time.Duration) *GetScopedPackageVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scoped package version params
func (o *GetScopedPackageVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scoped package version params
func (o *GetScopedPackageVersionParams) WithContext(ctx context.Context) *GetScopedPackageVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scoped package version params
func (o *GetScopedPackageVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scoped package version params
func (o *GetScopedPackageVersionParams) WithHTTPClient(client *http.Client) *GetScopedPackageVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scoped package version params
func (o *GetScopedPackageVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get scoped package version params
func (o *GetScopedPackageVersionParams) WithAPIVersion(aPIVersion string) *GetScopedPackageVersionParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get scoped package version params
func (o *GetScopedPackageVersionParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithFeedID adds the feedID to the get scoped package version params
func (o *GetScopedPackageVersionParams) WithFeedID(feedID string) *GetScopedPackageVersionParams {
	o.SetFeedID(feedID)
	return o
}

// SetFeedID adds the feedId to the get scoped package version params
func (o *GetScopedPackageVersionParams) SetFeedID(feedID string) {
	o.FeedID = feedID
}

// WithOrganization adds the organization to the get scoped package version params
func (o *GetScopedPackageVersionParams) WithOrganization(organization string) *GetScopedPackageVersionParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get scoped package version params
func (o *GetScopedPackageVersionParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPackageScope adds the packageScope to the get scoped package version params
func (o *GetScopedPackageVersionParams) WithPackageScope(packageScope string) *GetScopedPackageVersionParams {
	o.SetPackageScope(packageScope)
	return o
}

// SetPackageScope adds the packageScope to the get scoped package version params
func (o *GetScopedPackageVersionParams) SetPackageScope(packageScope string) {
	o.PackageScope = packageScope
}

// WithPackageVersion adds the packageVersion to the get scoped package version params
func (o *GetScopedPackageVersionParams) WithPackageVersion(packageVersion string) *GetScopedPackageVersionParams {
	o.SetPackageVersion(packageVersion)
	return o
}

// SetPackageVersion adds the packageVersion to the get scoped package version params
func (o *GetScopedPackageVersionParams) SetPackageVersion(packageVersion string) {
	o.PackageVersion = packageVersion
}

// WithUnscopedPackageName adds the unscopedPackageName to the get scoped package version params
func (o *GetScopedPackageVersionParams) WithUnscopedPackageName(unscopedPackageName string) *GetScopedPackageVersionParams {
	o.SetUnscopedPackageName(unscopedPackageName)
	return o
}

// SetUnscopedPackageName adds the unscopedPackageName to the get scoped package version params
func (o *GetScopedPackageVersionParams) SetUnscopedPackageName(unscopedPackageName string) {
	o.UnscopedPackageName = unscopedPackageName
}

// WriteToRequest writes these params to a swagger request
func (o *GetScopedPackageVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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