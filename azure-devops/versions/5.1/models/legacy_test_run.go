// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LegacyTestRun legacy test run
// swagger:model LegacyTestRun
type LegacyTestRun struct {

	// bugs count
	BugsCount int32 `json:"bugsCount,omitempty"`

	// build configuration Id
	BuildConfigurationID int32 `json:"buildConfigurationId,omitempty"`

	// build flavor
	BuildFlavor string `json:"buildFlavor,omitempty"`

	// build number
	BuildNumber string `json:"buildNumber,omitempty"`

	// build platform
	BuildPlatform string `json:"buildPlatform,omitempty"`

	// build reference
	BuildReference *LegacyBuildConfiguration `json:"buildReference,omitempty"`

	// build Uri
	BuildURI string `json:"buildUri,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// complete date
	// Format: date-time
	CompleteDate strfmt.DateTime `json:"completeDate,omitempty"`

	// configuration ids
	ConfigurationIds []int32 `json:"configurationIds"`

	// controller
	Controller string `json:"controller,omitempty"`

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// csm content
	CsmContent string `json:"csmContent,omitempty"`

	// csm parameters
	CsmParameters string `json:"csmParameters,omitempty"`

	// custom fields
	CustomFields []*TestExtensionField `json:"customFields"`

	// drop location
	DropLocation string `json:"dropLocation,omitempty"`

	// dtl aut environment
	DtlAutEnvironment *ShallowReference `json:"dtlAutEnvironment,omitempty"`

	// dtl test environment
	DtlTestEnvironment *ShallowReference `json:"dtlTestEnvironment,omitempty"`

	// due date
	// Format: date-time
	DueDate strfmt.DateTime `json:"dueDate,omitempty"`

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// filter
	Filter *RunFilter `json:"filter,omitempty"`

	// incomplete tests
	IncompleteTests int32 `json:"incompleteTests,omitempty"`

	// is automated
	IsAutomated bool `json:"isAutomated,omitempty"`

	// is bvt
	IsBvt bool `json:"isBvt,omitempty"`

	// iteration
	Iteration string `json:"iteration,omitempty"`

	// iteration Id
	IterationID int32 `json:"iterationId,omitempty"`

	// last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// last updated by
	// Format: uuid
	LastUpdatedBy strfmt.UUID `json:"lastUpdatedBy,omitempty"`

	// last updated by name
	LastUpdatedByName string `json:"lastUpdatedByName,omitempty"`

	// legacy share path
	LegacySharePath string `json:"legacySharePath,omitempty"`

	// not applicable tests
	NotApplicableTests int32 `json:"notApplicableTests,omitempty"`

	// owner
	// Format: uuid
	Owner strfmt.UUID `json:"owner,omitempty"`

	// owner name
	OwnerName string `json:"ownerName,omitempty"`

	// passed tests
	PassedTests int32 `json:"passedTests,omitempty"`

	// post process state
	// Format: byte
	PostProcessState strfmt.Base64 `json:"postProcessState,omitempty"`

	// public test settings Id
	PublicTestSettingsID int32 `json:"publicTestSettingsId,omitempty"`

	// release environment Uri
	ReleaseEnvironmentURI string `json:"releaseEnvironmentUri,omitempty"`

	// release reference
	ReleaseReference *LegacyReleaseReference `json:"releaseReference,omitempty"`

	// release Uri
	ReleaseURI string `json:"releaseUri,omitempty"`

	// revision
	Revision int32 `json:"revision,omitempty"`

	// row version
	RowVersion []strfmt.Base64 `json:"rowVersion"`

	// run has dtl environment
	RunHasDtlEnvironment bool `json:"runHasDtlEnvironment,omitempty"`

	// run timeout
	RunTimeout string `json:"runTimeout,omitempty"`

	// service version
	ServiceVersion string `json:"serviceVersion,omitempty"`

	// source workflow
	SourceWorkflow string `json:"sourceWorkflow,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// state
	// Format: byte
	State strfmt.Base64 `json:"state,omitempty"`

	// subscription name
	SubscriptionName string `json:"subscriptionName,omitempty"`

	// substate
	// Format: byte
	Substate strfmt.Base64 `json:"substate,omitempty"`

	// team project
	TeamProject string `json:"teamProject,omitempty"`

	// team project Uri
	TeamProjectURI string `json:"teamProjectUri,omitempty"`

	// test configurations mapping
	TestConfigurationsMapping string `json:"testConfigurationsMapping,omitempty"`

	// test environment Id
	// Format: uuid
	TestEnvironmentID strfmt.UUID `json:"testEnvironmentId,omitempty"`

	// test message log entries
	TestMessageLogEntries []*TestMessageLogDetails `json:"testMessageLogEntries"`

	// test message log Id
	TestMessageLogID int32 `json:"testMessageLogId,omitempty"`

	// test plan Id
	TestPlanID int32 `json:"testPlanId,omitempty"`

	// test run Id
	TestRunID int32 `json:"testRunId,omitempty"`

	// test run statistics
	TestRunStatistics []*LegacyTestRunStatistic `json:"testRunStatistics"`

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

// Validate validates this legacy test run
func (m *LegacyTestRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompleteDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDtlAutEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDtlTestEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
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

	if err := m.validateReleaseReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRowVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubstate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestEnvironmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestMessageLogEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestRunStatistics(formats); err != nil {
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

func (m *LegacyTestRun) validateBuildReference(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildReference) { // not required
		return nil
	}

	if m.BuildReference != nil {
		if err := m.BuildReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildReference")
			}
			return err
		}
	}

	return nil
}

func (m *LegacyTestRun) validateCompleteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompleteDate) { // not required
		return nil
	}

	if err := validate.FormatOf("completeDate", "body", "date-time", m.CompleteDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestRun) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestRun) validateCustomFields(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomFields) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomFields); i++ {
		if swag.IsZero(m.CustomFields[i]) { // not required
			continue
		}

		if m.CustomFields[i] != nil {
			if err := m.CustomFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LegacyTestRun) validateDtlAutEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.DtlAutEnvironment) { // not required
		return nil
	}

	if m.DtlAutEnvironment != nil {
		if err := m.DtlAutEnvironment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dtlAutEnvironment")
			}
			return err
		}
	}

	return nil
}

func (m *LegacyTestRun) validateDtlTestEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.DtlTestEnvironment) { // not required
		return nil
	}

	if m.DtlTestEnvironment != nil {
		if err := m.DtlTestEnvironment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dtlTestEnvironment")
			}
			return err
		}
	}

	return nil
}

func (m *LegacyTestRun) validateDueDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("dueDate", "body", "date-time", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestRun) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *LegacyTestRun) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestRun) validateLastUpdatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedBy", "body", "uuid", m.LastUpdatedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestRun) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if err := validate.FormatOf("owner", "body", "uuid", m.Owner.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestRun) validatePostProcessState(formats strfmt.Registry) error {

	if swag.IsZero(m.PostProcessState) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *LegacyTestRun) validateReleaseReference(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseReference) { // not required
		return nil
	}

	if m.ReleaseReference != nil {
		if err := m.ReleaseReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releaseReference")
			}
			return err
		}
	}

	return nil
}

func (m *LegacyTestRun) validateRowVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.RowVersion) { // not required
		return nil
	}

	for i := 0; i < len(m.RowVersion); i++ {

		// Format "byte" (base64 string) is already validated when unmarshalled

	}

	return nil
}

func (m *LegacyTestRun) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestRun) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *LegacyTestRun) validateSubstate(formats strfmt.Registry) error {

	if swag.IsZero(m.Substate) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *LegacyTestRun) validateTestEnvironmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.TestEnvironmentID) { // not required
		return nil
	}

	if err := validate.FormatOf("testEnvironmentId", "body", "uuid", m.TestEnvironmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestRun) validateTestMessageLogEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.TestMessageLogEntries) { // not required
		return nil
	}

	for i := 0; i < len(m.TestMessageLogEntries); i++ {
		if swag.IsZero(m.TestMessageLogEntries[i]) { // not required
			continue
		}

		if m.TestMessageLogEntries[i] != nil {
			if err := m.TestMessageLogEntries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testMessageLogEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LegacyTestRun) validateTestRunStatistics(formats strfmt.Registry) error {

	if swag.IsZero(m.TestRunStatistics) { // not required
		return nil
	}

	for i := 0; i < len(m.TestRunStatistics); i++ {
		if swag.IsZero(m.TestRunStatistics[i]) { // not required
			continue
		}

		if m.TestRunStatistics[i] != nil {
			if err := m.TestRunStatistics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testRunStatistics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LegacyTestRun) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (m *LegacyTestRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LegacyTestRun) UnmarshalBinary(b []byte) error {
	var res LegacyTestRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}