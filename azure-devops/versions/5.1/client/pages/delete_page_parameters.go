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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePageParams creates a new DeletePageParams object
// with the default values initialized.
func NewDeletePageParams() *DeletePageParams {
	var ()
	return &DeletePageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePageParamsWithTimeout creates a new DeletePageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePageParamsWithTimeout(timeout time.Duration) *DeletePageParams {
	var ()
	return &DeletePageParams{

		timeout: timeout,
	}
}

// NewDeletePageParamsWithContext creates a new DeletePageParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePageParamsWithContext(ctx context.Context) *DeletePageParams {
	var ()
	return &DeletePageParams{

		Context: ctx,
	}
}

// NewDeletePageParamsWithHTTPClient creates a new DeletePageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePageParamsWithHTTPClient(client *http.Client) *DeletePageParams {
	var ()
	return &DeletePageParams{
		HTTPClient: client,
	}
}

/*DeletePageParams contains all the parameters to send to the API endpoint
for the delete page operation typically these are written to a http.Request
*/
type DeletePageParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Comment
	  Comment to be associated with this page delete.

	*/
	Comment *string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Path
	  Wiki page path.

	*/
	Path string
	/*Project
	  Project ID or project name

	*/
	Project string
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

// WithTimeout adds the timeout to the delete page params
func (o *DeletePageParams) WithTimeout(timeout time.Duration) *DeletePageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete page params
func (o *DeletePageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete page params
func (o *DeletePageParams) WithContext(ctx context.Context) *DeletePageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete page params
func (o *DeletePageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete page params
func (o *DeletePageParams) WithHTTPClient(client *http.Client) *DeletePageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete page params
func (o *DeletePageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete page params
func (o *DeletePageParams) WithAPIVersion(aPIVersion string) *DeletePageParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete page params
func (o *DeletePageParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithComment adds the comment to the delete page params
func (o *DeletePageParams) WithComment(comment *string) *DeletePageParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the delete page params
func (o *DeletePageParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithOrganization adds the organization to the delete page params
func (o *DeletePageParams) WithOrganization(organization string) *DeletePageParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete page params
func (o *DeletePageParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPath adds the path to the delete page params
func (o *DeletePageParams) WithPath(path string) *DeletePageParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the delete page params
func (o *DeletePageParams) SetPath(path string) {
	o.Path = path
}

// WithProject adds the project to the delete page params
func (o *DeletePageParams) WithProject(project string) *DeletePageParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete page params
func (o *DeletePageParams) SetProject(project string) {
	o.Project = project
}

// WithVersionDescriptorVersion adds the versionDescriptorVersion to the delete page params
func (o *DeletePageParams) WithVersionDescriptorVersion(versionDescriptorVersion *string) *DeletePageParams {
	o.SetVersionDescriptorVersion(versionDescriptorVersion)
	return o
}

// SetVersionDescriptorVersion adds the versionDescriptorVersion to the delete page params
func (o *DeletePageParams) SetVersionDescriptorVersion(versionDescriptorVersion *string) {
	o.VersionDescriptorVersion = versionDescriptorVersion
}

// WithVersionDescriptorVersionOptions adds the versionDescriptorVersionOptions to the delete page params
func (o *DeletePageParams) WithVersionDescriptorVersionOptions(versionDescriptorVersionOptions *string) *DeletePageParams {
	o.SetVersionDescriptorVersionOptions(versionDescriptorVersionOptions)
	return o
}

// SetVersionDescriptorVersionOptions adds the versionDescriptorVersionOptions to the delete page params
func (o *DeletePageParams) SetVersionDescriptorVersionOptions(versionDescriptorVersionOptions *string) {
	o.VersionDescriptorVersionOptions = versionDescriptorVersionOptions
}

// WithVersionDescriptorVersionType adds the versionDescriptorVersionType to the delete page params
func (o *DeletePageParams) WithVersionDescriptorVersionType(versionDescriptorVersionType *string) *DeletePageParams {
	o.SetVersionDescriptorVersionType(versionDescriptorVersionType)
	return o
}

// SetVersionDescriptorVersionType adds the versionDescriptorVersionType to the delete page params
func (o *DeletePageParams) SetVersionDescriptorVersionType(versionDescriptorVersionType *string) {
	o.VersionDescriptorVersionType = versionDescriptorVersionType
}

// WithWikiIdentifier adds the wikiIdentifier to the delete page params
func (o *DeletePageParams) WithWikiIdentifier(wikiIdentifier string) *DeletePageParams {
	o.SetWikiIdentifier(wikiIdentifier)
	return o
}

// SetWikiIdentifier adds the wikiIdentifier to the delete page params
func (o *DeletePageParams) SetWikiIdentifier(wikiIdentifier string) {
	o.WikiIdentifier = wikiIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Comment != nil {

		// query param comment
		var qrComment string
		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {
			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}

	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {
		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
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
