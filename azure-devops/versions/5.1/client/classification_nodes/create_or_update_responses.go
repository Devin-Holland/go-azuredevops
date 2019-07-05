// Code generated by go-swagger; DO NOT EDIT.

package classification_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// CreateOrUpdateReader is a Reader for the CreateOrUpdate structure.
type CreateOrUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateOrUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOrUpdateOK creates a CreateOrUpdateOK with default headers values
func NewCreateOrUpdateOK() *CreateOrUpdateOK {
	return &CreateOrUpdateOK{}
}

/*CreateOrUpdateOK handles this case with default header values.

successful operation
*/
type CreateOrUpdateOK struct {
	Payload *models.WorkItemClassificationNode
}

func (o *CreateOrUpdateOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/{project}/_apis/wit/classificationnodes/{structureGroup}/{path}][%d] createOrUpdateOK  %+v", 200, o.Payload)
}

func (o *CreateOrUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkItemClassificationNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}