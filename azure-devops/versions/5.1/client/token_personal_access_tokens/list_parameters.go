// Code generated by go-swagger; DO NOT EDIT.

package token_personal_access_tokens

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

// NewListParams creates a new ListParams object
// with the default values initialized.
func NewListParams() *ListParams {
	var ()
	return &ListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListParamsWithTimeout creates a new ListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListParamsWithTimeout(timeout time.Duration) *ListParams {
	var ()
	return &ListParams{

		timeout: timeout,
	}
}

// NewListParamsWithContext creates a new ListParams object
// with the default values initialized, and the ability to set a context for a request
func NewListParamsWithContext(ctx context.Context) *ListParams {
	var ()
	return &ListParams{

		Context: ctx,
	}
}

// NewListParamsWithHTTPClient creates a new ListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListParamsWithHTTPClient(client *http.Client) *ListParams {
	var ()
	return &ListParams{
		HTTPClient: client,
	}
}

/*ListParams contains all the parameters to send to the API endpoint
for the list operation typically these are written to a http.Request
*/
type ListParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.1' to use this version of the api.

	*/
	APIVersion string
	/*Body*/
	Body []string
	/*ContinuationToken
	  An opaque data blob that allows the next page of data to resume immediately after where the previous page ended. The only reliable way to know if there is more data left is the presence of a continuation token.

	*/
	ContinuationToken *string
	/*IsPublic
	  Set to false for PAT tokens and true for SSH tokens.

	*/
	IsPublic *bool
	/*PageSize
	  The maximum number of results to return on each page.

	*/
	PageSize *int32
	/*SubjectDescriptor
	  The descriptor of the target user.

	*/
	SubjectDescriptor string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list params
func (o *ListParams) WithTimeout(timeout time.Duration) *ListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list params
func (o *ListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list params
func (o *ListParams) WithContext(ctx context.Context) *ListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list params
func (o *ListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list params
func (o *ListParams) WithHTTPClient(client *http.Client) *ListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list params
func (o *ListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the list params
func (o *ListParams) WithAPIVersion(aPIVersion string) *ListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list params
func (o *ListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the list params
func (o *ListParams) WithBody(body []string) *ListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list params
func (o *ListParams) SetBody(body []string) {
	o.Body = body
}

// WithContinuationToken adds the continuationToken to the list params
func (o *ListParams) WithContinuationToken(continuationToken *string) *ListParams {
	o.SetContinuationToken(continuationToken)
	return o
}

// SetContinuationToken adds the continuationToken to the list params
func (o *ListParams) SetContinuationToken(continuationToken *string) {
	o.ContinuationToken = continuationToken
}

// WithIsPublic adds the isPublic to the list params
func (o *ListParams) WithIsPublic(isPublic *bool) *ListParams {
	o.SetIsPublic(isPublic)
	return o
}

// SetIsPublic adds the isPublic to the list params
func (o *ListParams) SetIsPublic(isPublic *bool) {
	o.IsPublic = isPublic
}

// WithPageSize adds the pageSize to the list params
func (o *ListParams) WithPageSize(pageSize *int32) *ListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list params
func (o *ListParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSubjectDescriptor adds the subjectDescriptor to the list params
func (o *ListParams) WithSubjectDescriptor(subjectDescriptor string) *ListParams {
	o.SetSubjectDescriptor(subjectDescriptor)
	return o
}

// SetSubjectDescriptor adds the subjectDescriptor to the list params
func (o *ListParams) SetSubjectDescriptor(subjectDescriptor string) {
	o.SubjectDescriptor = subjectDescriptor
}

// WriteToRequest writes these params to a swagger request
func (o *ListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ContinuationToken != nil {

		// query param continuationToken
		var qrContinuationToken string
		if o.ContinuationToken != nil {
			qrContinuationToken = *o.ContinuationToken
		}
		qContinuationToken := qrContinuationToken
		if qContinuationToken != "" {
			if err := r.SetQueryParam("continuationToken", qContinuationToken); err != nil {
				return err
			}
		}

	}

	if o.IsPublic != nil {

		// query param isPublic
		var qrIsPublic bool
		if o.IsPublic != nil {
			qrIsPublic = *o.IsPublic
		}
		qIsPublic := swag.FormatBool(qrIsPublic)
		if qIsPublic != "" {
			if err := r.SetQueryParam("isPublic", qIsPublic); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

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
