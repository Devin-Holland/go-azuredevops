// Code generated by go-swagger; DO NOT EDIT.

package test_suites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetTestSuitesForPlanReader is a Reader for the GetTestSuitesForPlan structure.
type GetTestSuitesForPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestSuitesForPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTestSuitesForPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTestSuitesForPlanOK creates a GetTestSuitesForPlanOK with default headers values
func NewGetTestSuitesForPlanOK() *GetTestSuitesForPlanOK {
	return &GetTestSuitesForPlanOK{}
}

/*GetTestSuitesForPlanOK handles this case with default header values.

successful operation
*/
type GetTestSuitesForPlanOK struct {
	Payload []*models.TestSuite
}

func (o *GetTestSuitesForPlanOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/testplan/Plans/{planId}/suites][%d] getTestSuitesForPlanOK  %+v", 200, o.Payload)
}

func (o *GetTestSuitesForPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
