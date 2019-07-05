// Code generated by go-swagger; DO NOT EDIT.

package test_point

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

// NewGetPointsParams creates a new GetPointsParams object
// with the default values initialized.
func NewGetPointsParams() *GetPointsParams {
	var ()
	return &GetPointsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPointsParamsWithTimeout creates a new GetPointsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPointsParamsWithTimeout(timeout time.Duration) *GetPointsParams {
	var ()
	return &GetPointsParams{

		timeout: timeout,
	}
}

// NewGetPointsParamsWithContext creates a new GetPointsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPointsParamsWithContext(ctx context.Context) *GetPointsParams {
	var ()
	return &GetPointsParams{

		Context: ctx,
	}
}

// NewGetPointsParamsWithHTTPClient creates a new GetPointsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPointsParamsWithHTTPClient(client *http.Client) *GetPointsParams {
	var ()
	return &GetPointsParams{
		HTTPClient: client,
	}
}

/*GetPointsParams contains all the parameters to send to the API endpoint
for the get points operation typically these are written to a http.Request
*/
type GetPointsParams struct {

	/*APIVersion
	  Version of the API to use.  This should be set to '5.1-preview.2' to use this version of the api.

	*/
	APIVersion string
	/*Organization
	  The name of the Azure DevOps organization.

	*/
	Organization string
	/*PlanID
	  ID of the test plan for which test points are requested.

	*/
	PlanID int32
	/*PointIds
	  ID of test points to be fetched.

	*/
	PointIds string
	/*Project
	  Project ID or project name

	*/
	Project string
	/*ReturnIdentityRef
	  If set to true, returns the AssignedTo field in TestCaseReference as IdentityRef object.

	*/
	ReturnIdentityRef *bool
	/*SuiteID
	  ID of the test suite for which test points are requested.

	*/
	SuiteID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get points params
func (o *GetPointsParams) WithTimeout(timeout time.Duration) *GetPointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get points params
func (o *GetPointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get points params
func (o *GetPointsParams) WithContext(ctx context.Context) *GetPointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get points params
func (o *GetPointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get points params
func (o *GetPointsParams) WithHTTPClient(client *http.Client) *GetPointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get points params
func (o *GetPointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get points params
func (o *GetPointsParams) WithAPIVersion(aPIVersion string) *GetPointsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get points params
func (o *GetPointsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithOrganization adds the organization to the get points params
func (o *GetPointsParams) WithOrganization(organization string) *GetPointsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get points params
func (o *GetPointsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPlanID adds the planID to the get points params
func (o *GetPointsParams) WithPlanID(planID int32) *GetPointsParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the get points params
func (o *GetPointsParams) SetPlanID(planID int32) {
	o.PlanID = planID
}

// WithPointIds adds the pointIds to the get points params
func (o *GetPointsParams) WithPointIds(pointIds string) *GetPointsParams {
	o.SetPointIds(pointIds)
	return o
}

// SetPointIds adds the pointIds to the get points params
func (o *GetPointsParams) SetPointIds(pointIds string) {
	o.PointIds = pointIds
}

// WithProject adds the project to the get points params
func (o *GetPointsParams) WithProject(project string) *GetPointsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get points params
func (o *GetPointsParams) SetProject(project string) {
	o.Project = project
}

// WithReturnIdentityRef adds the returnIdentityRef to the get points params
func (o *GetPointsParams) WithReturnIdentityRef(returnIdentityRef *bool) *GetPointsParams {
	o.SetReturnIdentityRef(returnIdentityRef)
	return o
}

// SetReturnIdentityRef adds the returnIdentityRef to the get points params
func (o *GetPointsParams) SetReturnIdentityRef(returnIdentityRef *bool) {
	o.ReturnIdentityRef = returnIdentityRef
}

// WithSuiteID adds the suiteID to the get points params
func (o *GetPointsParams) WithSuiteID(suiteID int32) *GetPointsParams {
	o.SetSuiteID(suiteID)
	return o
}

// SetSuiteID adds the suiteId to the get points params
func (o *GetPointsParams) SetSuiteID(suiteID int32) {
	o.SuiteID = suiteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param planId
	if err := r.SetPathParam("planId", swag.FormatInt32(o.PlanID)); err != nil {
		return err
	}

	// path param pointIds
	if err := r.SetPathParam("pointIds", o.PointIds); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if o.ReturnIdentityRef != nil {

		// query param returnIdentityRef
		var qrReturnIdentityRef bool
		if o.ReturnIdentityRef != nil {
			qrReturnIdentityRef = *o.ReturnIdentityRef
		}
		qReturnIdentityRef := swag.FormatBool(qrReturnIdentityRef)
		if qReturnIdentityRef != "" {
			if err := r.SetQueryParam("returnIdentityRef", qReturnIdentityRef); err != nil {
				return err
			}
		}

	}

	// path param suiteId
	if err := r.SetPathParam("suiteId", swag.FormatInt32(o.SuiteID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
