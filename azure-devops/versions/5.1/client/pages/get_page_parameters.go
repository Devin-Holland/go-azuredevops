// Code generated by go-swagger; DO NOT EDIT.

package pages

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

// NewGetPageParams creates a new GetPageParams object
// with the default values initialized.
func NewGetPageParams() *GetPageParams {
	var ()
	return &GetPageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPageParamsWithTimeout creates a new GetPageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPageParamsWithTimeout(timeout time.Duration) *GetPageParams {
	var ()
	return &GetPageParams{

		timeout: timeout,
	}
}

// NewGetPageParamsWithContext creates a new GetPageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPageParamsWithContext(ctx context.Context) *GetPageParams {
	var ()
	return &GetPageParams{

		Context: ctx,
	}
}

// NewGetPageParamsWithHTTPClient creates a new GetPageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPageParamsWithHTTPClient(client *http.Client) *GetPageParams {
	var ()
	return &GetPageParams{
		HTTPClient: client,
	}
}

/*GetPageParams contains all the parameters to send to the API endpoint
for the get page operation typically these are written to a http.Request
*/
type GetPageParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*IncludeContent
	  True to include the content of the page in the response for Json content type. Defaults to false (Optional)

	*/
	IncludeContent *bool
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Path
	  Wiki page path.

	*/
	Path *string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*RecursionLevel
	  Recursion level for subpages retrieval. Defaults to `None` (Optional).

	*/
	RecursionLevel *string
	/*VersionDescriptorVersion
	  Version string identifier (name of tag/branch, SHA1 of commit)

	*/
	VersionDescriptorVersion *string
	/*VersionDescriptorVersionOptions
	  Version options - Specify additional modifiers to version (e.g Previous)

	*/
	VersionDescriptorVersionOptions *string
	/*VersionDescriptorVersionType
	  Version type (branch, tag, or commit). Determines how Id is interpreted

	*/
	VersionDescriptorVersionType *string
	/*WikiIdentifier
	  Wiki Id or name.

	*/
	WikiIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get page params
func (o *GetPageParams) WithTimeout(timeout time.Duration) *GetPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get page params
func (o *GetPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get page params
func (o *GetPageParams) WithContext(ctx context.Context) *GetPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get page params
func (o *GetPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get page params
func (o *GetPageParams) WithHTTPClient(client *http.Client) *GetPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get page params
func (o *GetPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get page params
func (o *GetPageParams) WithAPIVersion(aPIVersion string) *GetPageParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get page params
func (o *GetPageParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithIncludeContent adds the includeContent to the get page params
func (o *GetPageParams) WithIncludeContent(includeContent *bool) *GetPageParams {
	o.SetIncludeContent(includeContent)
	return o
}

// SetIncludeContent adds the includeContent to the get page params
func (o *GetPageParams) SetIncludeContent(includeContent *bool) {
	o.IncludeContent = includeContent
}

// WithOrganization adds the organization to the get page params
func (o *GetPageParams) WithOrganization(organization string) *GetPageParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get page params
func (o *GetPageParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPath adds the path to the get page params
func (o *GetPageParams) WithPath(path *string) *GetPageParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get page params
func (o *GetPageParams) SetPath(path *string) {
	o.Path = path
}

// WithProject adds the project to the get page params
func (o *GetPageParams) WithProject(project string) *GetPageParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get page params
func (o *GetPageParams) SetProject(project string) {
	o.Project = project
}

// WithRecursionLevel adds the recursionLevel to the get page params
func (o *GetPageParams) WithRecursionLevel(recursionLevel *string) *GetPageParams {
	o.SetRecursionLevel(recursionLevel)
	return o
}

// SetRecursionLevel adds the recursionLevel to the get page params
func (o *GetPageParams) SetRecursionLevel(recursionLevel *string) {
	o.RecursionLevel = recursionLevel
}

// WithVersionDescriptorVersion adds the versionDescriptorVersion to the get page params
func (o *GetPageParams) WithVersionDescriptorVersion(versionDescriptorVersion *string) *GetPageParams {
	o.SetVersionDescriptorVersion(versionDescriptorVersion)
	return o
}

// SetVersionDescriptorVersion adds the versionDescriptorVersion to the get page params
func (o *GetPageParams) SetVersionDescriptorVersion(versionDescriptorVersion *string) {
	o.VersionDescriptorVersion = versionDescriptorVersion
}

// WithVersionDescriptorVersionOptions adds the versionDescriptorVersionOptions to the get page params
func (o *GetPageParams) WithVersionDescriptorVersionOptions(versionDescriptorVersionOptions *string) *GetPageParams {
	o.SetVersionDescriptorVersionOptions(versionDescriptorVersionOptions)
	return o
}

// SetVersionDescriptorVersionOptions adds the versionDescriptorVersionOptions to the get page params
func (o *GetPageParams) SetVersionDescriptorVersionOptions(versionDescriptorVersionOptions *string) {
	o.VersionDescriptorVersionOptions = versionDescriptorVersionOptions
}

// WithVersionDescriptorVersionType adds the versionDescriptorVersionType to the get page params
func (o *GetPageParams) WithVersionDescriptorVersionType(versionDescriptorVersionType *string) *GetPageParams {
	o.SetVersionDescriptorVersionType(versionDescriptorVersionType)
	return o
}

// SetVersionDescriptorVersionType adds the versionDescriptorVersionType to the get page params
func (o *GetPageParams) SetVersionDescriptorVersionType(versionDescriptorVersionType *string) {
	o.VersionDescriptorVersionType = versionDescriptorVersionType
}

// WithWikiIdentifier adds the wikiIdentifier to the get page params
func (o *GetPageParams) WithWikiIdentifier(wikiIdentifier string) *GetPageParams {
	o.SetWikiIdentifier(wikiIdentifier)
	return o
}

// SetWikiIdentifier adds the wikiIdentifier to the get page params
func (o *GetPageParams) SetWikiIdentifier(wikiIdentifier string) {
	o.WikiIdentifier = wikiIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludeContent != nil {

		// query param includeContent
		var qrIncludeContent bool
		if o.IncludeContent != nil {
			qrIncludeContent = *o.IncludeContent
		}
		qIncludeContent := swag.FormatBool(qrIncludeContent)
		if qIncludeContent != "" {
			if err := r.SetQueryParam("includeContent", qIncludeContent); err != nil {
				return err
			}
		}

	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	if o.Path != nil {

		// query param path
		var qrPath string
		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {
			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}

	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if o.RecursionLevel != nil {

		// query param recursionLevel
		var qrRecursionLevel string
		if o.RecursionLevel != nil {
			qrRecursionLevel = *o.RecursionLevel
		}
		qRecursionLevel := qrRecursionLevel
		if qRecursionLevel != "" {
			if err := r.SetQueryParam("recursionLevel", qRecursionLevel); err != nil {
				return err
			}
		}

	}

	if o.VersionDescriptorVersion != nil {

		// query param versionDescriptor.version
		var qrVersionDescriptorVersion string
		if o.VersionDescriptorVersion != nil {
			qrVersionDescriptorVersion = *o.VersionDescriptorVersion
		}
		qVersionDescriptorVersion := qrVersionDescriptorVersion
		if qVersionDescriptorVersion != "" {
			if err := r.SetQueryParam("versionDescriptor.version", qVersionDescriptorVersion); err != nil {
				return err
			}
		}

	}

	if o.VersionDescriptorVersionOptions != nil {

		// query param versionDescriptor.versionOptions
		var qrVersionDescriptorVersionOptions string
		if o.VersionDescriptorVersionOptions != nil {
			qrVersionDescriptorVersionOptions = *o.VersionDescriptorVersionOptions
		}
		qVersionDescriptorVersionOptions := qrVersionDescriptorVersionOptions
		if qVersionDescriptorVersionOptions != "" {
			if err := r.SetQueryParam("versionDescriptor.versionOptions", qVersionDescriptorVersionOptions); err != nil {
				return err
			}
		}

	}

	if o.VersionDescriptorVersionType != nil {

		// query param versionDescriptor.versionType
		var qrVersionDescriptorVersionType string
		if o.VersionDescriptorVersionType != nil {
			qrVersionDescriptorVersionType = *o.VersionDescriptorVersionType
		}
		qVersionDescriptorVersionType := qrVersionDescriptorVersionType
		if qVersionDescriptorVersionType != "" {
			if err := r.SetQueryParam("versionDescriptor.versionType", qVersionDescriptorVersionType); err != nil {
				return err
			}
		}

	}

	// path param wikiIdentifier
	if err := r.SetPathParam("wikiIdentifier", o.WikiIdentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}