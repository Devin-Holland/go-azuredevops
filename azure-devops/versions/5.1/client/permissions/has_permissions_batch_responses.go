// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// HasPermissionsBatchReader is a Reader for the HasPermissionsBatch structure.
type HasPermissionsBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HasPermissionsBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHasPermissionsBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHasPermissionsBatchOK creates a HasPermissionsBatchOK with default headers values
func NewHasPermissionsBatchOK() *HasPermissionsBatchOK {
	return &HasPermissionsBatchOK{}
}

/*HasPermissionsBatchOK handles this case with default header values.

successful operation
*/
type HasPermissionsBatchOK struct {
	Payload *models.PermissionEvaluationBatch
}

func (o *HasPermissionsBatchOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/_apis/security/permissionevaluationbatch][%d] hasPermissionsBatchOK  %+v", 200, o.Payload)
}

func (o *HasPermissionsBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PermissionEvaluationBatch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
