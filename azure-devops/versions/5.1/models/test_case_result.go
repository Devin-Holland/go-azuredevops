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

// TestCaseResult Represents a test result.
// swagger:model TestCaseResult
type TestCaseResult struct {

	// Test attachment ID of action recording.
	AfnStripID int32 `json:"afnStripId,omitempty"`

	// Reference to area path of test.
	Area *ShallowReference `json:"area,omitempty"`

	// Reference to bugs linked to test result.
	AssociatedBugs []*ShallowReference `json:"associatedBugs"`

	// ID representing test method in a dll.
	AutomatedTestID string `json:"automatedTestId,omitempty"`

	// Fully qualified name of test executed.
	AutomatedTestName string `json:"automatedTestName,omitempty"`

	// Container to which test belongs.
	AutomatedTestStorage string `json:"automatedTestStorage,omitempty"`

	// Type of automated test.
	AutomatedTestType string `json:"automatedTestType,omitempty"`

	// TypeId of automated test.
	AutomatedTestTypeID string `json:"automatedTestTypeId,omitempty"`

	// Shallow reference to build associated with test result.
	Build *ShallowReference `json:"build,omitempty"`

	// Reference to build associated with test result.
	BuildReference *BuildReference `json:"buildReference,omitempty"`

	// Comment in a test result with maxSize= 1000 chars.
	Comment string `json:"comment,omitempty"`

	// Time when test execution completed. Completed date should be greater than StartedDate.
	// Format: date-time
	CompletedDate strfmt.DateTime `json:"completedDate,omitempty"`

	// Machine name where test executed.
	ComputerName string `json:"computerName,omitempty"`

	// Reference to test configuration. Type ShallowReference.
	Configuration *ShallowReference `json:"configuration,omitempty"`

	// Timestamp when test result created.
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// Additional properties of test result.
	CustomFields []*CustomTestField `json:"customFields"`

	// Duration of test execution in milliseconds. If not provided value will be set as CompletedDate - StartedDate
	DurationInMs float64 `json:"durationInMs,omitempty"`

	// Error message in test execution.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Information when test results started failing.
	FailingSince *FailingSince `json:"failingSince,omitempty"`

	// Failure type of test result. Valid Value= (Known Issue, New Issue, Regression, Unknown, None)
	FailureType string `json:"failureType,omitempty"`

	// ID of a test result.
	ID int32 `json:"id,omitempty"`

	// Test result details of test iterations used only for Manual Testing.
	IterationDetails []*TestIterationDetailsModel `json:"iterationDetails"`

	// Reference to identity last updated test result.
	LastUpdatedBy *IdentityRef `json:"lastUpdatedBy,omitempty"`

	// Last updated datetime of test result.
	// Format: date-time
	LastUpdatedDate strfmt.DateTime `json:"lastUpdatedDate,omitempty"`

	// Test outcome of test result. Valid values = (Unspecified, None, Passed, Failed, Inconclusive, Timeout, Aborted, Blocked, NotExecuted, Warning, Error, NotApplicable, Paused, InProgress, NotImpacted)
	Outcome string `json:"outcome,omitempty"`

	// Reference to test owner.
	Owner *IdentityRef `json:"owner,omitempty"`

	// Priority of test executed.
	Priority int32 `json:"priority,omitempty"`

	// Reference to team project.
	Project *ShallowReference `json:"project,omitempty"`

	// Shallow reference to release associated with test result.
	Release *ShallowReference `json:"release,omitempty"`

	// Reference to release associated with test result.
	ReleaseReference *ReleaseReference `json:"releaseReference,omitempty"`

	// ResetCount.
	ResetCount int32 `json:"resetCount,omitempty"`

	// Resolution state of test result.
	ResolutionState string `json:"resolutionState,omitempty"`

	// ID of resolution state.
	ResolutionStateID int32 `json:"resolutionStateId,omitempty"`

	// Hierarchy type of the result, default value of None means its leaf node.
	// Enum: [none rerun dataDriven orderedTest generic]
	ResultGroupType interface{} `json:"resultGroupType,omitempty"`

	// Revision number of test result.
	Revision int32 `json:"revision,omitempty"`

	// Reference to identity executed the test.
	RunBy *IdentityRef `json:"runBy,omitempty"`

	// Stacktrace with maxSize= 1000 chars.
	StackTrace string `json:"stackTrace,omitempty"`

	// Time when test execution started.
	// Format: date-time
	StartedDate strfmt.DateTime `json:"startedDate,omitempty"`

	// State of test result. Type TestRunState.
	State string `json:"state,omitempty"`

	// List of sub results inside a test result, if ResultGroupType is not None, it holds corresponding type sub results.
	SubResults []*TestSubResult `json:"subResults"`

	// Reference to the test executed.
	TestCase *ShallowReference `json:"testCase,omitempty"`

	// Reference ID of test used by test result. Type TestResultMetaData
	TestCaseReferenceID int32 `json:"testCaseReferenceId,omitempty"`

	// TestCaseRevision Number.
	TestCaseRevision int32 `json:"testCaseRevision,omitempty"`

	// Name of test.
	TestCaseTitle string `json:"testCaseTitle,omitempty"`

	// Reference to test plan test case workitem is part of.
	TestPlan *ShallowReference `json:"testPlan,omitempty"`

	// Reference to the test point executed.
	TestPoint *ShallowReference `json:"testPoint,omitempty"`

	// Reference to test run.
	TestRun *ShallowReference `json:"testRun,omitempty"`

	// Reference to test suite test case workitem is part of.
	TestSuite *ShallowReference `json:"testSuite,omitempty"`

	// Url of test result.
	URL string `json:"url,omitempty"`
}

// Validate validates this test case result
func (m *TestCaseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArea(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssociatedBugs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailingSince(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIterationDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestCase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestRun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestSuite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCaseResult) validateArea(formats strfmt.Registry) error {

	if swag.IsZero(m.Area) { // not required
		return nil
	}

	if m.Area != nil {
		if err := m.Area.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("area")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateAssociatedBugs(formats strfmt.Registry) error {

	if swag.IsZero(m.AssociatedBugs) { // not required
		return nil
	}

	for i := 0; i < len(m.AssociatedBugs); i++ {
		if swag.IsZero(m.AssociatedBugs[i]) { // not required
			continue
		}

		if m.AssociatedBugs[i] != nil {
			if err := m.AssociatedBugs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("associatedBugs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestCaseResult) validateBuild(formats strfmt.Registry) error {

	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {
		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateBuildReference(formats strfmt.Registry) error {

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

func (m *TestCaseResult) validateCompletedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("completedDate", "body", "date-time", m.CompletedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestCaseResult) validateConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestCaseResult) validateCustomFields(formats strfmt.Registry) error {

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

func (m *TestCaseResult) validateFailingSince(formats strfmt.Registry) error {

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

func (m *TestCaseResult) validateIterationDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.IterationDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.IterationDetails); i++ {
		if swag.IsZero(m.IterationDetails[i]) { // not required
			continue
		}

		if m.IterationDetails[i] != nil {
			if err := m.IterationDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("iterationDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestCaseResult) validateLastUpdatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedBy) { // not required
		return nil
	}

	if m.LastUpdatedBy != nil {
		if err := m.LastUpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUpdatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateLastUpdatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedDate", "body", "date-time", m.LastUpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestCaseResult) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateRelease(formats strfmt.Registry) error {

	if swag.IsZero(m.Release) { // not required
		return nil
	}

	if m.Release != nil {
		if err := m.Release.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("release")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateReleaseReference(formats strfmt.Registry) error {

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

func (m *TestCaseResult) validateRunBy(formats strfmt.Registry) error {

	if swag.IsZero(m.RunBy) { // not required
		return nil
	}

	if m.RunBy != nil {
		if err := m.RunBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runBy")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateStartedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startedDate", "body", "date-time", m.StartedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestCaseResult) validateSubResults(formats strfmt.Registry) error {

	if swag.IsZero(m.SubResults) { // not required
		return nil
	}

	for i := 0; i < len(m.SubResults); i++ {
		if swag.IsZero(m.SubResults[i]) { // not required
			continue
		}

		if m.SubResults[i] != nil {
			if err := m.SubResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestCaseResult) validateTestCase(formats strfmt.Registry) error {

	if swag.IsZero(m.TestCase) { // not required
		return nil
	}

	if m.TestCase != nil {
		if err := m.TestCase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testCase")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateTestPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.TestPlan) { // not required
		return nil
	}

	if m.TestPlan != nil {
		if err := m.TestPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testPlan")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateTestPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.TestPoint) { // not required
		return nil
	}

	if m.TestPoint != nil {
		if err := m.TestPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testPoint")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateTestRun(formats strfmt.Registry) error {

	if swag.IsZero(m.TestRun) { // not required
		return nil
	}

	if m.TestRun != nil {
		if err := m.TestRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testRun")
			}
			return err
		}
	}

	return nil
}

func (m *TestCaseResult) validateTestSuite(formats strfmt.Registry) error {

	if swag.IsZero(m.TestSuite) { // not required
		return nil
	}

	if m.TestSuite != nil {
		if err := m.TestSuite.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testSuite")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCaseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCaseResult) UnmarshalBinary(b []byte) error {
	var res TestCaseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}