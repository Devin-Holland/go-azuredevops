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

// NewDeleteScopedPackageVersionFromRecycleBinParams creates a new DeleteScopedPackageVersionFromRecycleBinParams object
// with the default values initialized.
func NewDeleteScopedPackageVersionFromRecycleBinParams() *DeleteScopedPackageVersionFromRecycleBinParams {
	var ()
	return &DeleteScopedPackageVersionFromRecycleBinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScopedPackageVersionFromRecycleBinParamsWithTimeout creates a new DeleteScopedPackageVersionFromRecycleBinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteScopedPackageVersionFromRecycleBinParamsWithTimeout(timeout time.Duration) *DeleteScopedPackageVersionFromRecycleBinParams {
	var ()
	return &DeleteScopedPackageVersionFromRecycleBinParams{

		timeout: timeout,
	}
}

// NewDeleteScopedPackageVersionFromRecycleBinParamsWithContext creates a new DeleteScopedPackageVersionFromRecycleBinParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteScopedPackageVersionFromRecycleBinParamsWithContext(ctx context.Context) *DeleteScopedPackageVersionFromRecycleBinParams {
	var ()
	return &DeleteScopedPackageVersionFromRecycleBinParams{

		Context: ctx,
	}
}

// NewDeleteScopedPackageVersionFromRecycleBinParamsWithHTTPClient creates a new DeleteScopedPackageVersionFromRecycleBinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteScopedPackageVersionFromRecycleBinParamsWithHTTPClient(client *http.Client) *DeleteScopedPackageVersionFromRecycleBinParams {
	var ()
	return &DeleteScopedPackageVersionFromRecycleBinParams{
		HTTPClient: client,
	}
}

/*DeleteScopedPackageVersionFromRecycleBinParams contains all the parameters to send to the API endpoint
for the delete scoped package version from recycle bin operation typically these are written to a http.Request
*/
type DeleteScopedPackageVersionFromRecycleBinParams struct {

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

// WithTimeout adds the timeout to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WithTimeout(timeout time.Duration) *DeleteScopedPackageVersionFromRecycleBinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WithContext(ctx context.Context) *DeleteScopedPackageVersionFromRecycleBinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WithHTTPClient(client *http.Client) *DeleteScopedPackageVersionFromRecycleBinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WithAPIVersion(aPIVersion string) *DeleteScopedPackageVersionFromRecycleBinParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithFeedID adds the feedID to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WithFeedID(feedID string) *DeleteScopedPackageVersionFromRecycleBinParams {
	o.SetFeedID(feedID)
	return o
}

// SetFeedID adds the feedId to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) SetFeedID(feedID string) {
	o.FeedID = feedID
}

// WithOrganization adds the organization to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WithOrganization(organization string) *DeleteScopedPackageVersionFromRecycleBinParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPackageScope adds the packageScope to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WithPackageScope(packageScope string) *DeleteScopedPackageVersionFromRecycleBinParams {
	o.SetPackageScope(packageScope)
	return o
}

// SetPackageScope adds the packageScope to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) SetPackageScope(packageScope string) {
	o.PackageScope = packageScope
}

// WithPackageVersion adds the packageVersion to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WithPackageVersion(packageVersion string) *DeleteScopedPackageVersionFromRecycleBinParams {
	o.SetPackageVersion(packageVersion)
	return o
}

// SetPackageVersion adds the packageVersion to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) SetPackageVersion(packageVersion string) {
	o.PackageVersion = packageVersion
}

// WithUnscopedPackageName adds the unscopedPackageName to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WithUnscopedPackageName(unscopedPackageName string) *DeleteScopedPackageVersionFromRecycleBinParams {
	o.SetUnscopedPackageName(unscopedPackageName)
	return o
}

// SetUnscopedPackageName adds the unscopedPackageName to the delete scoped package version from recycle bin params
func (o *DeleteScopedPackageVersionFromRecycleBinParams) SetUnscopedPackageName(unscopedPackageName string) {
	o.UnscopedPackageName = unscopedPackageName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScopedPackageVersionFromRecycleBinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
