// Code generated by go-swagger; DO NOT EDIT.

package python

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

// NewGetPackageVersionFromRecycleBinParams creates a new GetPackageVersionFromRecycleBinParams object
// with the default values initialized.
func NewGetPackageVersionFromRecycleBinParams() *GetPackageVersionFromRecycleBinParams {
	var ()
	return &GetPackageVersionFromRecycleBinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackageVersionFromRecycleBinParamsWithTimeout creates a new GetPackageVersionFromRecycleBinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPackageVersionFromRecycleBinParamsWithTimeout(timeout time.Duration) *GetPackageVersionFromRecycleBinParams {
	var ()
	return &GetPackageVersionFromRecycleBinParams{

		timeout: timeout,
	}
}

// NewGetPackageVersionFromRecycleBinParamsWithContext creates a new GetPackageVersionFromRecycleBinParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPackageVersionFromRecycleBinParamsWithContext(ctx context.Context) *GetPackageVersionFromRecycleBinParams {
	var ()
	return &GetPackageVersionFromRecycleBinParams{

		Context: ctx,
	}
}

// NewGetPackageVersionFromRecycleBinParamsWithHTTPClient creates a new GetPackageVersionFromRecycleBinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPackageVersionFromRecycleBinParamsWithHTTPClient(client *http.Client) *GetPackageVersionFromRecycleBinParams {
	var ()
	return &GetPackageVersionFromRecycleBinParams{
		HTTPClient: client,
	}
}

/*GetPackageVersionFromRecycleBinParams contains all the parameters to send to the API endpoint
for the get package version from recycle bin operation typically these are written to a http.Request
*/
type GetPackageVersionFromRecycleBinParams struct {

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
	/*PackageName
	  Name of the package.

	*/
	PackageName string
	/*PackageVersion
	  Version of the package.

	*/
	PackageVersion string
	/*Project
	  Project ID or project name

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) WithTimeout(timeout time.Duration) *GetPackageVersionFromRecycleBinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) WithContext(ctx context.Context) *GetPackageVersionFromRecycleBinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) WithHTTPClient(client *http.Client) *GetPackageVersionFromRecycleBinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) WithAPIVersion(aPIVersion string) *GetPackageVersionFromRecycleBinParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithFeedID adds the feedID to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) WithFeedID(feedID string) *GetPackageVersionFromRecycleBinParams {
	o.SetFeedID(feedID)
	return o
}

// SetFeedID adds the feedId to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) SetFeedID(feedID string) {
	o.FeedID = feedID
}

// WithOrganization adds the organization to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) WithOrganization(organization string) *GetPackageVersionFromRecycleBinParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPackageName adds the packageName to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) WithPackageName(packageName string) *GetPackageVersionFromRecycleBinParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithPackageVersion adds the packageVersion to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) WithPackageVersion(packageVersion string) *GetPackageVersionFromRecycleBinParams {
	o.SetPackageVersion(packageVersion)
	return o
}

// SetPackageVersion adds the packageVersion to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) SetPackageVersion(packageVersion string) {
	o.PackageVersion = packageVersion
}

// WithProject adds the project to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) WithProject(project string) *GetPackageVersionFromRecycleBinParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get package version from recycle bin params
func (o *GetPackageVersionFromRecycleBinParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackageVersionFromRecycleBinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
	}

	// path param packageVersion
	if err := r.SetPathParam("packageVersion", o.PackageVersion); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
