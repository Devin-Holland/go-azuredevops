// Code generated by go-swagger; DO NOT EDIT.

package test_plan_clone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// CloneTestPlanReader is a Reader for the CloneTestPlan structure.
type CloneTestPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneTestPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCloneTestPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCloneTestPlanOK creates a CloneTestPlanOK with default headers values
func NewCloneTestPlanOK() *CloneTestPlanOK {
	return &CloneTestPlanOK{}
}

/*CloneTestPlanOK handles this case with default header values.

successful operation
*/
type CloneTestPlanOK struct {
	Payload *models.CloneTestPlanOperationInformation
}

func (o *CloneTestPlanOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/{project}/_apis/testplan/Plans/CloneOperation][%d] cloneTestPlanOK  %+v", 200, o.Payload)
}

func (o *CloneTestPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloneTestPlanOperationInformation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}