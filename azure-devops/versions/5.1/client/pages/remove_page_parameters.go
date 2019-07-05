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

// NewRemovePageParams creates a new RemovePageParams object
// with the default values initialized.
func NewRemovePageParams() *RemovePageParams {
	var ()
	return &RemovePageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemovePageParamsWithTimeout creates a new RemovePageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemovePageParamsWithTimeout(timeout time.Duration) *RemovePageParams {
	var ()
	return &RemovePageParams{

		timeout: timeout,
	}
}

// NewRemovePageParamsWithContext creates a new RemovePageParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemovePageParamsWithContext(ctx context.Context) *RemovePageParams {
	var ()
	return &RemovePageParams{

		Context: ctx,
	}
}

// NewRemovePageParamsWithHTTPClient creates a new RemovePageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemovePageParamsWithHTTPClient(client *http.Client) *RemovePageParams {
	var ()
	return &RemovePageParams{
		HTTPClient: client,
	}
}

/*RemovePageParams contains all the parameters to send to the API endpoint
for the remove page operation typically these are written to a http.Request
*/
type RemovePageParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PageID
	  The ID of the page

	*/
	PageID string
	/*ProcessID
	  The ID of the process

	*/
	ProcessID strfmt.UUID
	/*WitRefName
	  The reference name of the work item type

	*/
	WitRefName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove page params
func (o *RemovePageParams) WithTimeout(timeout time.Duration) *RemovePageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove page params
func (o *RemovePageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove page params
func (o *RemovePageParams) WithContext(ctx context.Context) *RemovePageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove page params
func (o *RemovePageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove page params
func (o *RemovePageParams) WithHTTPClient(client *http.Client) *RemovePageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove page params
func (o *RemovePageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the remove page params
func (o *RemovePageParams) WithAPIVersion(aPIVersion string) *RemovePageParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the remove page params
func (o *RemovePageParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the remove page params
func (o *RemovePageParams) WithOrganization(organization string) *RemovePageParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the remove page params
func (o *RemovePageParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPageID adds the pageID to the remove page params
func (o *RemovePageParams) WithPageID(pageID string) *RemovePageParams {
	o.SetPageID(pageID)
	return o
}

// SetPageID adds the pageId to the remove page params
func (o *RemovePageParams) SetPageID(pageID string) {
	o.PageID = pageID
}

// WithProcessID adds the processID to the remove page params
func (o *RemovePageParams) WithProcessID(processID strfmt.UUID) *RemovePageParams {
	o.SetProcessID(processID)
	return o
}

// SetProcessID adds the processId to the remove page params
func (o *RemovePageParams) SetProcessID(processID strfmt.UUID) {
	o.ProcessID = processID
}

// WithWitRefName adds the witRefName to the remove page params
func (o *RemovePageParams) WithWitRefName(witRefName string) *RemovePageParams {
	o.SetWitRefName(witRefName)
	return o
}

// SetWitRefName adds the witRefName to the remove page params
func (o *RemovePageParams) SetWitRefName(witRefName string) {
	o.WitRefName = witRefName
}

// WriteToRequest writes these params to a swagger request
func (o *RemovePageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param pageId
	if err := r.SetPathParam("pageId", o.PageID); err != nil {
		return err
	}

	// path param processId
	if err := r.SetPathParam("processId", o.ProcessID.String()); err != nil {
		return err
	}

	// path param witRefName
	if err := r.SetPathParam("witRefName", o.WitRefName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}