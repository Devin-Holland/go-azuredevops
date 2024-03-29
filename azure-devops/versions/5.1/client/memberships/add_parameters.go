// Code generated by go-swagger; DO NOT EDIT.

package memberships

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

// NewAddParams creates a new AddParams object
// with the default values initialized.
func NewAddParams() *AddParams {
	var ()
	return &AddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddParamsWithTimeout creates a new AddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddParamsWithTimeout(timeout time.Duration) *AddParams {
	var ()
	return &AddParams{

		timeout: timeout,
	}
}

// NewAddParamsWithContext creates a new AddParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddParamsWithContext(ctx context.Context) *AddParams {
	var ()
	return &AddParams{

		Context: ctx,
	}
}

// NewAddParamsWithHTTPClient creates a new AddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddParamsWithHTTPClient(client *http.Client) *AddParams {
	var ()
	return &AddParams{
		HTTPClient: client,
	}
}

/*AddParams contains all the parameters to send to the API endpoint
for the add operation typically these are written to a http.Request
*/
type AddParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*ContainerDescriptor
	  A descriptor to a group that can be the container in the relationship.

	*/
	ContainerDescriptor string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*SubjectDescriptor
	  A descriptor to a group or user that can be the child subject in the relationship.

	*/
	SubjectDescriptor string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add params
func (o *AddParams) WithTimeout(timeout time.Duration) *AddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add params
func (o *AddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add params
func (o *AddParams) WithContext(ctx context.Context) *AddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add params
func (o *AddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add params
func (o *AddParams) WithHTTPClient(client *http.Client) *AddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add params
func (o *AddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the add params
func (o *AddParams) WithAPIVersion(aPIVersion string) *AddParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the add params
func (o *AddParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithContainerDescriptor adds the containerDescriptor to the add params
func (o *AddParams) WithContainerDescriptor(containerDescriptor string) *AddParams {
	o.SetContainerDescriptor(containerDescriptor)
	return o
}

// SetContainerDescriptor adds the containerDescriptor to the add params
func (o *AddParams) SetContainerDescriptor(containerDescriptor string) {
	o.ContainerDescriptor = containerDescriptor
}

// WithOrganization adds the organization to the add params
func (o *AddParams) WithOrganization(organization string) *AddParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the add params
func (o *AddParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithSubjectDescriptor adds the subjectDescriptor to the add params
func (o *AddParams) WithSubjectDescriptor(subjectDescriptor string) *AddParams {
	o.SetSubjectDescriptor(subjectDescriptor)
	return o
}

// SetSubjectDescriptor adds the subjectDescriptor to the add params
func (o *AddParams) SetSubjectDescriptor(subjectDescriptor string) {
	o.SubjectDescriptor = subjectDescriptor
}

// WriteToRequest writes these params to a swagger request
func (o *AddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param containerDescriptor
	if err := r.SetPathParam("containerDescriptor", o.ContainerDescriptor); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param subjectDescriptor
	if err := r.SetPathParam("subjectDescriptor", o.SubjectDescriptor); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
