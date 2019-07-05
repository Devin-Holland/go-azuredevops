// Code generated by go-swagger; DO NOT EDIT.

package test_suite_clone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// CloneTestSuiteReader is a Reader for the CloneTestSuite structure.
type CloneTestSuiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneTestSuiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCloneTestSuiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCloneTestSuiteOK creates a CloneTestSuiteOK with default headers values
func NewCloneTestSuiteOK() *CloneTestSuiteOK {
	return &CloneTestSuiteOK{}
}

/*CloneTestSuiteOK handles this case with default header values.

successful operation
*/
type CloneTestSuiteOK struct {
	Payload *models.CloneTestSuiteOperationInformation
}

func (o *CloneTestSuiteOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/{project}/_apis/testplan/Suites/CloneOperation][%d] cloneTestSuiteOK  %+v", 200, o.Payload)
}

func (o *CloneTestSuiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloneTestSuiteOperationInformation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
