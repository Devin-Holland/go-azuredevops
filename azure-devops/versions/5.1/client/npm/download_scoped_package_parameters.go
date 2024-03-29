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

// NewDownloadScopedPackageParams creates a new DownloadScopedPackageParams object
// with the default values initialized.
func NewDownloadScopedPackageParams() *DownloadScopedPackageParams {
	var ()
	return &DownloadScopedPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadScopedPackageParamsWithTimeout creates a new DownloadScopedPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadScopedPackageParamsWithTimeout(timeout time.Duration) *DownloadScopedPackageParams {
	var ()
	return &DownloadScopedPackageParams{

		timeout: timeout,
	}
}

// NewDownloadScopedPackageParamsWithContext creates a new DownloadScopedPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadScopedPackageParamsWithContext(ctx context.Context) *DownloadScopedPackageParams {
	var ()
	return &DownloadScopedPackageParams{

		Context: ctx,
	}
}

// NewDownloadScopedPackageParamsWithHTTPClient creates a new DownloadScopedPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadScopedPackageParamsWithHTTPClient(client *http.Client) *DownloadScopedPackageParams {
	var ()
	return &DownloadScopedPackageParams{
		HTTPClient: client,
	}
}

/*DownloadScopedPackageParams contains all the parameters to send to the API endpoint
for the download scoped package operation typically these are written to a http.Request
*/
type DownloadScopedPackageParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*FeedID*/
	FeedID string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PackageScope*/
	PackageScope string
	/*PackageVersion*/
	PackageVersion string
	/*UnscopedPackageName*/
	UnscopedPackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download scoped package params
func (o *DownloadScopedPackageParams) WithTimeout(timeout time.Duration) *DownloadScopedPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download scoped package params
func (o *DownloadScopedPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download scoped package params
func (o *DownloadScopedPackageParams) WithContext(ctx context.Context) *DownloadScopedPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download scoped package params
func (o *DownloadScopedPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download scoped package params
func (o *DownloadScopedPackageParams) WithHTTPClient(client *http.Client) *DownloadScopedPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download scoped package params
func (o *DownloadScopedPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the download scoped package params
func (o *DownloadScopedPackageParams) WithAPIVersion(aPIVersion string) *DownloadScopedPackageParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the download scoped package params
func (o *DownloadScopedPackageParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithFeedID adds the feedID to the download scoped package params
func (o *DownloadScopedPackageParams) WithFeedID(feedID string) *DownloadScopedPackageParams {
	o.SetFeedID(feedID)
	return o
}

// SetFeedID adds the feedId to the download scoped package params
func (o *DownloadScopedPackageParams) SetFeedID(feedID string) {
	o.FeedID = feedID
}

// WithOrganization adds the organization to the download scoped package params
func (o *DownloadScopedPackageParams) WithOrganization(organization string) *DownloadScopedPackageParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the download scoped package params
func (o *DownloadScopedPackageParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPackageScope adds the packageScope to the download scoped package params
func (o *DownloadScopedPackageParams) WithPackageScope(packageScope string) *DownloadScopedPackageParams {
	o.SetPackageScope(packageScope)
	return o
}

// SetPackageScope adds the packageScope to the download scoped package params
func (o *DownloadScopedPackageParams) SetPackageScope(packageScope string) {
	o.PackageScope = packageScope
}

// WithPackageVersion adds the packageVersion to the download scoped package params
func (o *DownloadScopedPackageParams) WithPackageVersion(packageVersion string) *DownloadScopedPackageParams {
	o.SetPackageVersion(packageVersion)
	return o
}

// SetPackageVersion adds the packageVersion to the download scoped package params
func (o *DownloadScopedPackageParams) SetPackageVersion(packageVersion string) {
	o.PackageVersion = packageVersion
}

// WithUnscopedPackageName adds the unscopedPackageName to the download scoped package params
func (o *DownloadScopedPackageParams) WithUnscopedPackageName(unscopedPackageName string) *DownloadScopedPackageParams {
	o.SetUnscopedPackageName(unscopedPackageName)
	return o
}

// SetUnscopedPackageName adds the unscopedPackageName to the download scoped package params
func (o *DownloadScopedPackageParams) SetUnscopedPackageName(unscopedPackageName string) {
	o.UnscopedPackageName = unscopedPackageName
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadScopedPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
