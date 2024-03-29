// Code generated by go-swagger; DO NOT EDIT.

package user_entitlements

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

	models "azure-devops/5.1/models"
)

// NewUpdateUserEntitlementParams creates a new UpdateUserEntitlementParams object
// with the default values initialized.
func NewUpdateUserEntitlementParams() *UpdateUserEntitlementParams {
	var ()
	return &UpdateUserEntitlementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserEntitlementParamsWithTimeout creates a new UpdateUserEntitlementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserEntitlementParamsWithTimeout(timeout time.Duration) *UpdateUserEntitlementParams {
	var ()
	return &UpdateUserEntitlementParams{

		timeout: timeout,
	}
}

// NewUpdateUserEntitlementParamsWithContext creates a new UpdateUserEntitlementParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserEntitlementParamsWithContext(ctx context.Context) *UpdateUserEntitlementParams {
	var ()
	return &UpdateUserEntitlementParams{

		Context: ctx,
	}
}

// NewUpdateUserEntitlementParamsWithHTTPClient creates a new UpdateUserEntitlementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserEntitlementParamsWithHTTPClient(client *http.Client) *UpdateUserEntitlementParams {
	var ()
	return &UpdateUserEntitlementParams{
		HTTPClient: client,
	}
}

/*UpdateUserEntitlementParams contains all the parameters to send to the API endpoint
for the update user entitlement operation typically these are written to a http.Request
*/
type UpdateUserEntitlementParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Body
	  JsonPatchDocument containing the operations to perform on the user.

	*/
	Body *models.JSONPatchDocument
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*UserID
	  ID of the user.

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user entitlement params
func (o *UpdateUserEntitlementParams) WithTimeout(timeout time.Duration) *UpdateUserEntitlementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user entitlement params
func (o *UpdateUserEntitlementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user entitlement params
func (o *UpdateUserEntitlementParams) WithContext(ctx context.Context) *UpdateUserEntitlementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user entitlement params
func (o *UpdateUserEntitlementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user entitlement params
func (o *UpdateUserEntitlementParams) WithHTTPClient(client *http.Client) *UpdateUserEntitlementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user entitlement params
func (o *UpdateUserEntitlementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update user entitlement params
func (o *UpdateUserEntitlementParams) WithAPIVersion(aPIVersion string) *UpdateUserEntitlementParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update user entitlement params
func (o *UpdateUserEntitlementParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update user entitlement params
func (o *UpdateUserEntitlementParams) WithBody(body *models.JSONPatchDocument) *UpdateUserEntitlementParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user entitlement params
func (o *UpdateUserEntitlementParams) SetBody(body *models.JSONPatchDocument) {
	o.Body = body
}

// WithOrganization adds the organization to the update user entitlement params
func (o *UpdateUserEntitlementParams) WithOrganization(organization string) *UpdateUserEntitlementParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the update user entitlement params
func (o *UpdateUserEntitlementParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithUserID adds the userID to the update user entitlement params
func (o *UpdateUserEntitlementParams) WithUserID(userID strfmt.UUID) *UpdateUserEntitlementParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update user entitlement params
func (o *UpdateUserEntitlementParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserEntitlementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
