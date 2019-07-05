// Code generated by go-swagger; DO NOT EDIT.

package artifact_details

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

// NewGetPackageParams creates a new GetPackageParams object
// with the default values initialized.
func NewGetPackageParams() *GetPackageParams {
	var ()
	return &GetPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackageParamsWithTimeout creates a new GetPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPackageParamsWithTimeout(timeout time.Duration) *GetPackageParams {
	var ()
	return &GetPackageParams{

		timeout: timeout,
	}
}

// NewGetPackageParamsWithContext creates a new GetPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPackageParamsWithContext(ctx context.Context) *GetPackageParams {
	var ()
	return &GetPackageParams{

		Context: ctx,
	}
}

// NewGetPackageParamsWithHTTPClient creates a new GetPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPackageParamsWithHTTPClient(client *http.Client) *GetPackageParams {
	var ()
	return &GetPackageParams{
		HTTPClient: client,
	}
}

/*GetPackageParams contains all the parameters to send to the API endpoint
for the get package operation typically these are written to a http.Request
*/
type GetPackageParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*FeedID
	  Name or Id of the feed.

	*/
	FeedID string
	/*IncludeAllVersions
	  True to return all versions of the package in the response. Default is false (latest version only).

	*/
	IncludeAllVersions *bool
	/*IncludeDeleted
	  Return deleted or unpublished versions of packages in the response. Default is False.

	*/
	IncludeDeleted *bool
	/*IncludeDescription
	  Return the description for every version of each package in the response. Default is False.

	*/
	IncludeDescription *bool
	/*IncludeUrls
	  True to return REST Urls with the response. Default is True.

	*/
	IncludeUrls *bool
	/*IsListed
	  Only applicable for NuGet packages, setting it for other package types will result in a 404. If false, delisted package versions will be returned. Use this to filter the response when includeAllVersions is set to true. Default is unset (do not return delisted packages).

	*/
	IsListed *bool
	/*IsRelease
	  Only applicable for Nuget packages. Use this to filter the response when includeAllVersions is set to true.  Default is True (only return packages without prerelease versioning).

	*/
	IsRelease *bool
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PackageID
	  The package Id (GUID Id, not the package name).

	*/
	PackageID string
	/*Project
	  Project ID or project name

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get package params
func (o *GetPackageParams) WithTimeout(timeout time.Duration) *GetPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get package params
func (o *GetPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get package params
func (o *GetPackageParams) WithContext(ctx context.Context) *GetPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get package params
func (o *GetPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get package params
func (o *GetPackageParams) WithHTTPClient(client *http.Client) *GetPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get package params
func (o *GetPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get package params
func (o *GetPackageParams) WithAPIVersion(aPIVersion string) *GetPackageParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get package params
func (o *GetPackageParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithFeedID adds the feedID to the get package params
func (o *GetPackageParams) WithFeedID(feedID string) *GetPackageParams {
	o.SetFeedID(feedID)
	return o
}

// SetFeedID adds the feedId to the get package params
func (o *GetPackageParams) SetFeedID(feedID string) {
	o.FeedID = feedID
}

// WithIncludeAllVersions adds the includeAllVersions to the get package params
func (o *GetPackageParams) WithIncludeAllVersions(includeAllVersions *bool) *GetPackageParams {
	o.SetIncludeAllVersions(includeAllVersions)
	return o
}

// SetIncludeAllVersions adds the includeAllVersions to the get package params
func (o *GetPackageParams) SetIncludeAllVersions(includeAllVersions *bool) {
	o.IncludeAllVersions = includeAllVersions
}

// WithIncludeDeleted adds the includeDeleted to the get package params
func (o *GetPackageParams) WithIncludeDeleted(includeDeleted *bool) *GetPackageParams {
	o.SetIncludeDeleted(includeDeleted)
	return o
}

// SetIncludeDeleted adds the includeDeleted to the get package params
func (o *GetPackageParams) SetIncludeDeleted(includeDeleted *bool) {
	o.IncludeDeleted = includeDeleted
}

// WithIncludeDescription adds the includeDescription to the get package params
func (o *GetPackageParams) WithIncludeDescription(includeDescription *bool) *GetPackageParams {
	o.SetIncludeDescription(includeDescription)
	return o
}

// SetIncludeDescription adds the includeDescription to the get package params
func (o *GetPackageParams) SetIncludeDescription(includeDescription *bool) {
	o.IncludeDescription = includeDescription
}

// WithIncludeUrls adds the includeUrls to the get package params
func (o *GetPackageParams) WithIncludeUrls(includeUrls *bool) *GetPackageParams {
	o.SetIncludeUrls(includeUrls)
	return o
}

// SetIncludeUrls adds the includeUrls to the get package params
func (o *GetPackageParams) SetIncludeUrls(includeUrls *bool) {
	o.IncludeUrls = includeUrls
}

// WithIsListed adds the isListed to the get package params
func (o *GetPackageParams) WithIsListed(isListed *bool) *GetPackageParams {
	o.SetIsListed(isListed)
	return o
}

// SetIsListed adds the isListed to the get package params
func (o *GetPackageParams) SetIsListed(isListed *bool) {
	o.IsListed = isListed
}

// WithIsRelease adds the isRelease to the get package params
func (o *GetPackageParams) WithIsRelease(isRelease *bool) *GetPackageParams {
	o.SetIsRelease(isRelease)
	return o
}

// SetIsRelease adds the isRelease to the get package params
func (o *GetPackageParams) SetIsRelease(isRelease *bool) {
	o.IsRelease = isRelease
}

// WithOrganization adds the organization to the get package params
func (o *GetPackageParams) WithOrganization(organization string) *GetPackageParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get package params
func (o *GetPackageParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPackageID adds the packageID to the get package params
func (o *GetPackageParams) WithPackageID(packageID string) *GetPackageParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the get package params
func (o *GetPackageParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WithProject adds the project to the get package params
func (o *GetPackageParams) WithProject(project string) *GetPackageParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get package params
func (o *GetPackageParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludeAllVersions != nil {

		// query param includeAllVersions
		var qrIncludeAllVersions bool
		if o.IncludeAllVersions != nil {
			qrIncludeAllVersions = *o.IncludeAllVersions
		}
		qIncludeAllVersions := swag.FormatBool(qrIncludeAllVersions)
		if qIncludeAllVersions != "" {
			if err := r.SetQueryParam("includeAllVersions", qIncludeAllVersions); err != nil {
				return err
			}
		}

	}

	if o.IncludeDeleted != nil {

		// query param includeDeleted
		var qrIncludeDeleted bool
		if o.IncludeDeleted != nil {
			qrIncludeDeleted = *o.IncludeDeleted
		}
		qIncludeDeleted := swag.FormatBool(qrIncludeDeleted)
		if qIncludeDeleted != "" {
			if err := r.SetQueryParam("includeDeleted", qIncludeDeleted); err != nil {
				return err
			}
		}

	}

	if o.IncludeDescription != nil {

		// query param includeDescription
		var qrIncludeDescription bool
		if o.IncludeDescription != nil {
			qrIncludeDescription = *o.IncludeDescription
		}
		qIncludeDescription := swag.FormatBool(qrIncludeDescription)
		if qIncludeDescription != "" {
			if err := r.SetQueryParam("includeDescription", qIncludeDescription); err != nil {
				return err
			}
		}

	}

	if o.IncludeUrls != nil {

		// query param includeUrls
		var qrIncludeUrls bool
		if o.IncludeUrls != nil {
			qrIncludeUrls = *o.IncludeUrls
		}
		qIncludeUrls := swag.FormatBool(qrIncludeUrls)
		if qIncludeUrls != "" {
			if err := r.SetQueryParam("includeUrls", qIncludeUrls); err != nil {
				return err
			}
		}

	}

	if o.IsListed != nil {

		// query param isListed
		var qrIsListed bool
		if o.IsListed != nil {
			qrIsListed = *o.IsListed
		}
		qIsListed := swag.FormatBool(qrIsListed)
		if qIsListed != "" {
			if err := r.SetQueryParam("isListed", qIsListed); err != nil {
				return err
			}
		}

	}

	if o.IsRelease != nil {

		// query param isRelease
		var qrIsRelease bool
		if o.IsRelease != nil {
			qrIsRelease = *o.IsRelease
		}
		qIsRelease := swag.FormatBool(qrIsRelease)
		if qIsRelease != "" {
			if err := r.SetQueryParam("isRelease", qIsRelease); err != nil {
				return err
			}
		}

	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID); err != nil {
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