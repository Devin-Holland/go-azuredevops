// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewRemoveGroupParams creates a new RemoveGroupParams object
// with the default values initialized.
func NewRemoveGroupParams() *RemoveGroupParams {
	var ()
	return &RemoveGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveGroupParamsWithTimeout creates a new RemoveGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveGroupParamsWithTimeout(timeout time.Duration) *RemoveGroupParams {
	var ()
	return &RemoveGroupParams{

		timeout: timeout,
	}
}

// NewRemoveGroupParamsWithContext creates a new RemoveGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveGroupParamsWithContext(ctx context.Context) *RemoveGroupParams {
	var ()
	return &RemoveGroupParams{

		Context: ctx,
	}
}

// NewRemoveGroupParamsWithHTTPClient creates a new RemoveGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveGroupParamsWithHTTPClient(client *http.Client) *RemoveGroupParams {
	var ()
	return &RemoveGroupParams{
		HTTPClient: client,
	}
}

/*RemoveGroupParams contains all the parameters to send to the API endpoint
for the remove group operation typically these are written to a http.Request
*/
type RemoveGroupParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*GroupID
	  The ID of the group

	*/
	GroupID string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PageID
	  The ID of the page the group is in

	*/
	PageID string
	/*ProcessID
	  The ID of the process

	*/
	ProcessID strfmt.UUID
	/*SectionID
	  The ID of the section to the group is in

	*/
	SectionID string
	/*WitRefName
	  The reference name of the work item type

	*/
	WitRefName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove group params
func (o *RemoveGroupParams) WithTimeout(timeout time.Duration) *RemoveGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove group params
func (o *RemoveGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove group params
func (o *RemoveGroupParams) WithContext(ctx context.Context) *RemoveGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove group params
func (o *RemoveGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove group params
func (o *RemoveGroupParams) WithHTTPClient(client *http.Client) *RemoveGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove group params
func (o *RemoveGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the remove group params
func (o *RemoveGroupParams) WithAPIVersion(aPIVersion string) *RemoveGroupParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the remove group params
func (o *RemoveGroupParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGroupID adds the groupID to the remove group params
func (o *RemoveGroupParams) WithGroupID(groupID string) *RemoveGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the remove group params
func (o *RemoveGroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithOrganization adds the organization to the remove group params
func (o *RemoveGroupParams) WithOrganization(organization string) *RemoveGroupParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the remove group params
func (o *RemoveGroupParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPageID adds the pageID to the remove group params
func (o *RemoveGroupParams) WithPageID(pageID string) *RemoveGroupParams {
	o.SetPageID(pageID)
	return o
}

// SetPageID adds the pageId to the remove group params
func (o *RemoveGroupParams) SetPageID(pageID string) {
	o.PageID = pageID
}

// WithProcessID adds the processID to the remove group params
func (o *RemoveGroupParams) WithProcessID(processID strfmt.UUID) *RemoveGroupParams {
	o.SetProcessID(processID)
	return o
}

// SetProcessID adds the processId to the remove group params
func (o *RemoveGroupParams) SetProcessID(processID strfmt.UUID) {
	o.ProcessID = processID
}

// WithSectionID adds the sectionID to the remove group params
func (o *RemoveGroupParams) WithSectionID(sectionID string) *RemoveGroupParams {
	o.SetSectionID(sectionID)
	return o
}

// SetSectionID adds the sectionId to the remove group params
func (o *RemoveGroupParams) SetSectionID(sectionID string) {
	o.SectionID = sectionID
}

// WithWitRefName adds the witRefName to the remove group params
func (o *RemoveGroupParams) WithWitRefName(witRefName string) *RemoveGroupParams {
	o.SetWitRefName(witRefName)
	return o
}

// SetWitRefName adds the witRefName to the remove group params
func (o *RemoveGroupParams) SetWitRefName(witRefName string) {
	o.WitRefName = witRefName
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
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

	// path param sectionId
	if err := r.SetPathParam("sectionId", o.SectionID); err != nil {
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
