// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// CreateRequestsReader is a Reader for the CreateRequests structure.
type CreateRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateRequestsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewCreateRequestsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRequestsCreated creates a CreateRequestsCreated with default headers values
func NewCreateRequestsCreated() *CreateRequestsCreated {
	return &CreateRequestsCreated{}
}

/*CreateRequestsCreated handles this case with default header values.

Successfully created a new symbol request.
*/
type CreateRequestsCreated struct {
	Payload *models.Request
}

func (o *CreateRequestsCreated) Error() string {
	return fmt.Sprintf("[POST /{organization}/_apis/symbol/requests][%d] createRequestsCreated  %+v", 201, o.Payload)
}

func (o *CreateRequestsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRequestsConflict creates a CreateRequestsConflict with default headers values
func NewCreateRequestsConflict() *CreateRequestsConflict {
	return &CreateRequestsConflict{}
}

/*CreateRequestsConflict handles this case with default header values.

The symbol request couldn't be created due to a server-side conflict.
*/
type CreateRequestsConflict struct {
	Payload *models.Request
}

func (o *CreateRequestsConflict) Error() string {
	return fmt.Sprintf("[POST /{organization}/_apis/symbol/requests][%d] createRequestsConflict  %+v", 409, o.Payload)
}

func (o *CreateRequestsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
