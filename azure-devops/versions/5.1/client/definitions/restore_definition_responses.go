// Code generated by go-swagger; DO NOT EDIT.

package definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// RestoreDefinitionReader is a Reader for the RestoreDefinition structure.
type RestoreDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRestoreDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRestoreDefinitionOK creates a RestoreDefinitionOK with default headers values
func NewRestoreDefinitionOK() *RestoreDefinitionOK {
	return &RestoreDefinitionOK{}
}

/*RestoreDefinitionOK handles this case with default header values.

successful operation
*/
type RestoreDefinitionOK struct {
	Payload *models.BuildDefinition
}

func (o *RestoreDefinitionOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/{project}/_apis/build/definitions/{definitionId}][%d] restoreDefinitionOK  %+v", 200, o.Payload)
}

func (o *RestoreDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
