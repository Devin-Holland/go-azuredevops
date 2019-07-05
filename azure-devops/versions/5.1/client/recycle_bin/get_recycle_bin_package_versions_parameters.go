// Code generated by go-swagger; DO NOT EDIT.

package recycle_bin

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

// NewGetRecycleBinPackageVersionsParams creates a new GetRecycleBinPackageVersionsParams object
// with the default values initialized.
func NewGetRecycleBinPackageVersionsParams() *GetRecycleBinPackageVersionsParams {
	var ()
	return &GetRecycleBinPackageVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecycleBinPackageVersionsParamsWithTimeout creates a new GetRecycleBinPackageVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecycleBinPackageVersionsParamsWithTimeout(timeout time.Duration) *GetRecycleBinPackageVersionsParams {
	var ()
	return &GetRecycleBinPackageVersionsParams{

		timeout: timeout,
	}
}

// NewGetRecycleBinPackageVersionsParamsWithContext creates a new GetRecycleBinPackageVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecycleBinPackageVersionsParamsWithContext(ctx context.Context) *GetRecycleBinPackageVersionsParams {
	var ()
	return &GetRecycleBinPackageVersionsParams{

		Context: ctx,
	}
}

// NewGetRecycleBinPackageVersionsParamsWithHTTPClient creates a new GetRecycleBinPackageVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecycleBinPackageVersionsParamsWithHTTPClient(client *http.Client) *GetRecycleBinPackageVersionsParams {
	var ()
	return &GetRecycleBinPackageVersionsParams{
		HTTPClient: client,
	}
}

/*GetRecycleBinPackageVersionsParams contains all the parameters to send to the API endpoint
for the get recycle bin package versions operation typically these are written to a http.Request
*/
type GetRecycleBinPackageVersionsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*FeedID
	  Name or Id of the feed.

	*/
	FeedID string
	/*IncludeUrls
	  True to return REST Urls with the response.  Default is True.

	*/
	IncludeUrls *bool
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PackageID
	  The package Id (GUID Id, not the package name).

	*/
	PackageID strfmt.UUID
	/*Project
	  Project ID or project name

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) WithTimeout(timeout time.Duration) *GetRecycleBinPackageVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) WithContext(ctx context.Context) *GetRecycleBinPackageVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) WithHTTPClient(client *http.Client) *GetRecycleBinPackageVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) WithAPIVersion(aPIVersion string) *GetRecycleBinPackageVersionsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithFeedID adds the feedID to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) WithFeedID(feedID string) *GetRecycleBinPackageVersionsParams {
	o.SetFeedID(feedID)
	return o
}

// SetFeedID adds the feedId to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) SetFeedID(feedID string) {
	o.FeedID = feedID
}

// WithIncludeUrls adds the includeUrls to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) WithIncludeUrls(includeUrls *bool) *GetRecycleBinPackageVersionsParams {
	o.SetIncludeUrls(includeUrls)
	return o
}

// SetIncludeUrls adds the includeUrls to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) SetIncludeUrls(includeUrls *bool) {
	o.IncludeUrls = includeUrls
}

// WithOrganization adds the organization to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) WithOrganization(organization string) *GetRecycleBinPackageVersionsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPackageID adds the packageID to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) WithPackageID(packageID strfmt.UUID) *GetRecycleBinPackageVersionsParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) SetPackageID(packageID strfmt.UUID) {
	o.PackageID = packageID
}

// WithProject adds the project to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) WithProject(project string) *GetRecycleBinPackageVersionsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get recycle bin package versions params
func (o *GetRecycleBinPackageVersionsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecycleBinPackageVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID.String()); err != nil {
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