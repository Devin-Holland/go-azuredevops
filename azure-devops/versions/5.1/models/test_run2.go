// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestRun2 test run2
// swagger:model TestRun2
type TestRun2 struct {

	// build configuration Id
	BuildConfigurationID int32 `json:"buildConfigurationId,omitempty"`

	// build number
	BuildNumber string `json:"buildNumber,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// complete date
	// Format: date-time
	CompleteDate strfmt.DateTime `json:"completeDate,omitempty"`

	// controller
	Controller string `json:"controller,omitempty"`

	// coverage Id
	CoverageID int32 `json:"coverageId,omitempty"`

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// deleted on
	// Format: date-time
	DeletedOn strfmt.DateTime `json:"deletedOn,omitempty"`

	// drop location
	DropLocation string `json:"dropLocation,omitempty"`

	// due date
	// Format: date-time
	DueDate strfmt.DateTime `json:"dueDate,omitempty"`

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// incomplete tests
	IncompleteTests int32 `json:"incompleteTests,omitempty"`

	// is automated
	IsAutomated bool `json:"isAutomated,omitempty"`

	// is bvt
	IsBvt bool `json:"isBvt,omitempty"`

	// is migrated
	IsMigrated bool `json:"isMigrated,omitempty"`

	// iteration Id
	IterationID int32 `json:"iterationId,omitempty"`

	// last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// last updated by
	// Format: uuid
	LastUpdatedBy strfmt.UUID `json:"lastUpdatedBy,omitempty"`

	// legacy share path
	LegacySharePath string `json:"legacySharePath,omitempty"`

	// max reserved result Id
	MaxReservedResultID int32 `json:"maxReservedResultId,omitempty"`

	// not applicable tests
	NotApplicableTests int32 `json:"notApplicableTests,omitempty"`

	// owner
	// Format: uuid
	Owner strfmt.UUID `json:"owner,omitempty"`

	// passed tests
	PassedTests int32 `json:"passedTests,omitempty"`

	// post process state
	// Format: byte
	PostProcessState strfmt.Base64 `json:"postProcessState,omitempty"`

	// project Id
	// Format: uuid
	ProjectID strfmt.UUID `json:"projectId,omitempty"`

	// public test settings Id
	PublicTestSettingsID int32 `json:"publicTestSettingsId,omitempty"`

	// release environment Uri
	ReleaseEnvironmentURI string `json:"releaseEnvironmentUri,omitempty"`

	// release Uri
	ReleaseURI string `json:"releaseUri,omitempty"`

	// revision
	Revision int32 `json:"revision,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// state
	// Format: byte
	State strfmt.Base64 `json:"state,omitempty"`

	// test environment Id
	// Format: uuid
	TestEnvironmentID strfmt.UUID `json:"testEnvironmentId,omitempty"`

	// test message log Id
	TestMessageLogID int32 `json:"testMessageLogId,omitempty"`

	// test plan Id
	TestPlanID int32 `json:"testPlanId,omitempty"`

	// test run context Id
	TestRunContextID int32 `json:"testRunContextId,omitempty"`

	// test run Id
	TestRunID int32 `json:"testRunId,omitempty"`

	// test settings Id
	TestSettingsID int32 `json:"testSettingsId,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// total tests
	TotalTests int32 `json:"totalTests,omitempty"`

	// type
	// Format: byte
	Type strfmt.Base64 `json:"type,omitempty"`

	// unanalyzed tests
	UnanalyzedTests int32 `json:"unanalyzedTests,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this test run2
func (m *TestRun2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompleteDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostProcessState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestEnvironmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestRun2) validateCompleteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompleteDate) { // not required
		return nil
	}

	if err := validate.FormatOf("completeDate", "body", "date-time", m.CompleteDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validateDeletedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("deletedOn", "body", "date-time", m.DeletedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validateDueDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("dueDate", "body", "date-time", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validateLastUpdatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedBy", "body", "uuid", m.LastUpdatedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if err := validate.FormatOf("owner", "body", "uuid", m.Owner.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validatePostProcessState(formats strfmt.Registry) error {

	if swag.IsZero(m.PostProcessState) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *TestRun2) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectId", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *TestRun2) validateTestEnvironmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.TestEnvironmentID) { // not required
		return nil
	}

	if err := validate.FormatOf("testEnvironmentId", "body", "uuid", m.TestEnvironmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestRun2) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (m *TestRun2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestRun2) UnmarshalBinary(b []byte) error {
	var res TestRun2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
