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

// LegacyTestCaseResult legacy test case result
// swagger:model LegacyTestCaseResult
type LegacyTestCaseResult struct {

	// afn strip Id
	AfnStripID int32 `json:"afnStripId,omitempty"`

	// area Id
	AreaID int32 `json:"areaId,omitempty"`

	// area Uri
	AreaURI string `json:"areaUri,omitempty"`

	// automated test Id
	AutomatedTestID string `json:"automatedTestId,omitempty"`

	// automated test name
	AutomatedTestName string `json:"automatedTestName,omitempty"`

	// automated test storage
	AutomatedTestStorage string `json:"automatedTestStorage,omitempty"`

	// automated test type
	AutomatedTestType string `json:"automatedTestType,omitempty"`

	// automated test type Id
	AutomatedTestTypeID string `json:"automatedTestTypeId,omitempty"`

	// build number
	BuildNumber string `json:"buildNumber,omitempty"`

	// build reference
	BuildReference *LegacyBuildConfiguration `json:"buildReference,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// computer name
	ComputerName string `json:"computerName,omitempty"`

	// configuration Id
	ConfigurationID int32 `json:"configurationId,omitempty"`

	// configuration name
	ConfigurationName string `json:"configurationName,omitempty"`

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// custom fields
	CustomFields []*TestExtensionField `json:"customFields"`

	// date completed
	// Format: date-time
	DateCompleted strfmt.DateTime `json:"dateCompleted,omitempty"`

	// date started
	// Format: date-time
	DateStarted strfmt.DateTime `json:"dateStarted,omitempty"`

	// duration
	Duration int64 `json:"duration,omitempty"`

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// failing since
	FailingSince *FailingSince `json:"failingSince,omitempty"`

	// failure type
	// Format: byte
	FailureType strfmt.Base64 `json:"failureType,omitempty"`

	// id
	ID *LegacyTestCaseResultIdentifier `json:"id,omitempty"`

	// is rerun
	IsRerun bool `json:"isRerun,omitempty"`

	// last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// last updated by
	// Format: uuid
	LastUpdatedBy strfmt.UUID `json:"lastUpdatedBy,omitempty"`

	// last updated by name
	LastUpdatedByName string `json:"lastUpdatedByName,omitempty"`

	// outcome
	// Format: byte
	Outcome strfmt.Base64 `json:"outcome,omitempty"`

	// owner
	// Format: uuid
	Owner strfmt.UUID `json:"owner,omitempty"`

	// owner name
	OwnerName string `json:"ownerName,omitempty"`

	// priority
	// Format: byte
	Priority strfmt.Base64 `json:"priority,omitempty"`

	// release reference
	ReleaseReference *LegacyReleaseReference `json:"releaseReference,omitempty"`

	// reset count
	ResetCount int32 `json:"resetCount,omitempty"`

	// resolution state Id
	ResolutionStateID int32 `json:"resolutionStateId,omitempty"`

	// result group type
	// Enum: [none rerun dataDriven orderedTest generic]
	ResultGroupType interface{} `json:"resultGroupType,omitempty"`

	// revision
	Revision int32 `json:"revision,omitempty"`

	// run by
	// Format: uuid
	RunBy strfmt.UUID `json:"runBy,omitempty"`

	// run by name
	RunByName string `json:"runByName,omitempty"`

	// sequence Id
	SequenceID int32 `json:"sequenceId,omitempty"`

	// stack trace
	StackTrace *TestExtensionField `json:"stackTrace,omitempty"`

	// state
	// Format: byte
	State strfmt.Base64 `json:"state,omitempty"`

	// sub result count
	SubResultCount int32 `json:"subResultCount,omitempty"`

	// suite name
	SuiteName string `json:"suiteName,omitempty"`

	// test case area
	TestCaseArea string `json:"testCaseArea,omitempty"`

	// test case area Uri
	TestCaseAreaURI string `json:"testCaseAreaUri,omitempty"`

	// test case Id
	TestCaseID int32 `json:"testCaseId,omitempty"`

	// test case reference Id
	TestCaseReferenceID int32 `json:"testCaseReferenceId,omitempty"`

	// test case revision
	TestCaseRevision int32 `json:"testCaseRevision,omitempty"`

	// test case title
	TestCaseTitle string `json:"testCaseTitle,omitempty"`

	// test plan Id
	TestPlanID int32 `json:"testPlanId,omitempty"`

	// test point Id
	TestPointID int32 `json:"testPointId,omitempty"`

	// test result Id
	TestResultID int32 `json:"testResultId,omitempty"`

	// test run Id
	TestRunID int32 `json:"testRunId,omitempty"`

	// test run title
	TestRunTitle string `json:"testRunTitle,omitempty"`

	// test suite Id
	TestSuiteID int32 `json:"testSuiteId,omitempty"`
}

// Validate validates this legacy test case result
func (m *LegacyTestCaseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailingSince(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailureType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutcome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackTrace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LegacyTestCaseResult) validateBuildReference(formats strfmt.Registry) error {

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

func (m *LegacyTestCaseResult) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestCaseResult) validateCustomFields(formats strfmt.Registry) error {

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

func (m *LegacyTestCaseResult) validateDateCompleted(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCompleted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCompleted", "body", "date-time", m.DateCompleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestCaseResult) validateDateStarted(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStarted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStarted", "body", "date-time", m.DateStarted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestCaseResult) validateFailingSince(formats strfmt.Registry) error {

	if swag.IsZero(m.FailingSince) { // not required
		return nil
	}

	if m.FailingSince != nil {
		if err := m.FailingSince.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failingSince")
			}
			return err
		}
	}

	return nil
}

func (m *LegacyTestCaseResult) validateFailureType(formats strfmt.Registry) error {

	if swag.IsZero(m.FailureType) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *LegacyTestCaseResult) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *LegacyTestCaseResult) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestCaseResult) validateLastUpdatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedBy", "body", "uuid", m.LastUpdatedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestCaseResult) validateOutcome(formats strfmt.Registry) error {

	if swag.IsZero(m.Outcome) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *LegacyTestCaseResult) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if err := validate.FormatOf("owner", "body", "uuid", m.Owner.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestCaseResult) validatePriority(formats strfmt.Registry) error {

	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *LegacyTestCaseResult) validateReleaseReference(formats strfmt.Registry) error {

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

func (m *LegacyTestCaseResult) validateRunBy(formats strfmt.Registry) error {

	if swag.IsZero(m.RunBy) { // not required
		return nil
	}

	if err := validate.FormatOf("runBy", "body", "uuid", m.RunBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyTestCaseResult) validateStackTrace(formats strfmt.Registry) error {

	if swag.IsZero(m.StackTrace) { // not required
		return nil
	}

	if m.StackTrace != nil {
		if err := m.StackTrace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackTrace")
			}
			return err
		}
	}

	return nil
}

func (m *LegacyTestCaseResult) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (m *LegacyTestCaseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LegacyTestCaseResult) UnmarshalBinary(b []byte) error {
	var res LegacyTestCaseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
