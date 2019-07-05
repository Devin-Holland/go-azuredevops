// Code generated by go-swagger; DO NOT EDIT.

package permissions

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

// NewHasPermissionsParams creates a new HasPermissionsParams object
// with the default values initialized.
func NewHasPermissionsParams() *HasPermissionsParams {
	var ()
	return &HasPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHasPermissionsParamsWithTimeout creates a new HasPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHasPermissionsParamsWithTimeout(timeout time.Duration) *HasPermissionsParams {
	var ()
	return &HasPermissionsParams{

		timeout: timeout,
	}
}

// NewHasPermissionsParamsWithContext creates a new HasPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewHasPermissionsParamsWithContext(ctx context.Context) *HasPermissionsParams {
	var ()
	return &HasPermissionsParams{

		Context: ctx,
	}
}

// NewHasPermissionsParamsWithHTTPClient creates a new HasPermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHasPermissionsParamsWithHTTPClient(client *http.Client) *HasPermissionsParams {
	var ()
	return &HasPermissionsParams{
		HTTPClient: client,
	}
}

/*HasPermissionsParams contains all the parameters to send to the API endpoint
for the has permissions operation typically these are written to a http.Request
*/
type HasPermissionsParams struct {

	/*AlwaysAllowAdministrators
	  If true and if the caller is an administrator, always return true.

	*/
	AlwaysAllowAdministrators *bool
	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Delimiter
	  Optional security token separator. Defaults to ",".

	*/
	Delimiter *string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*Permissions
	  Permissions to evaluate.

	*/
	Permissions int32
	/*SecurityNamespaceID
	  Security namespace identifier.

	*/
	SecurityNamespaceID strfmt.UUID
	/*Tokens
	  One or more security tokens to evaluate.

	*/
	Tokens *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the has permissions params
func (o *HasPermissionsParams) WithTimeout(timeout time.Duration) *HasPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the has permissions params
func (o *HasPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the has permissions params
func (o *HasPermissionsParams) WithContext(ctx context.Context) *HasPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the has permissions params
func (o *HasPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the has permissions params
func (o *HasPermissionsParams) WithHTTPClient(client *http.Client) *HasPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the has permissions params
func (o *HasPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlwaysAllowAdministrators adds the alwaysAllowAdministrators to the has permissions params
func (o *HasPermissionsParams) WithAlwaysAllowAdministrators(alwaysAllowAdministrators *bool) *HasPermissionsParams {
	o.SetAlwaysAllowAdministrators(alwaysAllowAdministrators)
	return o
}

// SetAlwaysAllowAdministrators adds the alwaysAllowAdministrators to the has permissions params
func (o *HasPermissionsParams) SetAlwaysAllowAdministrators(alwaysAllowAdministrators *bool) {
	o.AlwaysAllowAdministrators = alwaysAllowAdministrators
}

// WithAPIVersion adds the aPIVersion to the has permissions params
func (o *HasPermissionsParams) WithAPIVersion(aPIVersion string) *HasPermissionsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the has permissions params
func (o *HasPermissionsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDelimiter adds the delimiter to the has permissions params
func (o *HasPermissionsParams) WithDelimiter(delimiter *string) *HasPermissionsParams {
	o.SetDelimiter(delimiter)
	return o
}

// SetDelimiter adds the delimiter to the has permissions params
func (o *HasPermissionsParams) SetDelimiter(delimiter *string) {
	o.Delimiter = delimiter
}

// WithOrganization adds the organization to the has permissions params
func (o *HasPermissionsParams) WithOrganization(organization string) *HasPermissionsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the has permissions params
func (o *HasPermissionsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPermissions adds the permissions to the has permissions params
func (o *HasPermissionsParams) WithPermissions(permissions int32) *HasPermissionsParams {
	o.SetPermissions(permissions)
	return o
}

// SetPermissions adds the permissions to the has permissions params
func (o *HasPermissionsParams) SetPermissions(permissions int32) {
	o.Permissions = permissions
}

// WithSecurityNamespaceID adds the securityNamespaceID to the has permissions params
func (o *HasPermissionsParams) WithSecurityNamespaceID(securityNamespaceID strfmt.UUID) *HasPermissionsParams {
	o.SetSecurityNamespaceID(securityNamespaceID)
	return o
}

// SetSecurityNamespaceID adds the securityNamespaceId to the has permissions params
func (o *HasPermissionsParams) SetSecurityNamespaceID(securityNamespaceID strfmt.UUID) {
	o.SecurityNamespaceID = securityNamespaceID
}

// WithTokens adds the tokens to the has permissions params
func (o *HasPermissionsParams) WithTokens(tokens *string) *HasPermissionsParams {
	o.SetTokens(tokens)
	return o
}

// SetTokens adds the tokens to the has permissions params
func (o *HasPermissionsParams) SetTokens(tokens *string) {
	o.Tokens = tokens
}

// WriteToRequest writes these params to a swagger request
func (o *HasPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AlwaysAllowAdministrators != nil {

		// query param alwaysAllowAdministrators
		var qrAlwaysAllowAdministrators bool
		if o.AlwaysAllowAdministrators != nil {
			qrAlwaysAllowAdministrators = *o.AlwaysAllowAdministrators
		}
		qAlwaysAllowAdministrators := swag.FormatBool(qrAlwaysAllowAdministrators)
		if qAlwaysAllowAdministrators != "" {
			if err := r.SetQueryParam("alwaysAllowAdministrators", qAlwaysAllowAdministrators); err != nil {
				return err
			}
		}

	}

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	if o.Delimiter != nil {

		// query param delimiter
		var qrDelimiter string
		if o.Delimiter != nil {
			qrDelimiter = *o.Delimiter
		}
		qDelimiter := qrDelimiter
		if qDelimiter != "" {
			if err := r.SetQueryParam("delimiter", qDelimiter); err != nil {
				return err
			}
		}

	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param permissions
	if err := r.SetPathParam("permissions", swag.FormatInt32(o.Permissions)); err != nil {
		return err
	}

	// path param securityNamespaceId
	if err := r.SetPathParam("securityNamespaceId", o.SecurityNamespaceID.String()); err != nil {
		return err
	}

	if o.Tokens != nil {

		// query param tokens
		var qrTokens string
		if o.Tokens != nil {
			qrTokens = *o.Tokens
		}
		qTokens := qrTokens
		if qTokens != "" {
			if err := r.SetQueryParam("tokens", qTokens); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
