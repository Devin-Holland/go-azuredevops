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

// GetRequestsRequestNameReader is a Reader for the GetRequestsRequestName structure.
type GetRequestsRequestNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestsRequestNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRequestsRequestNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetRequestsRequestNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRequestsRequestNameOK creates a GetRequestsRequestNameOK with default headers values
func NewGetRequestsRequestNameOK() *GetRequestsRequestNameOK {
	return &GetRequestsRequestNameOK{}
}

/*GetRequestsRequestNameOK handles this case with default header values.

Successfully retrieved the symbol request as specified by the request name.
*/
type GetRequestsRequestNameOK struct {
	Payload *models.Request
}

func (o *GetRequestsRequestNameOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/_apis/symbol/requests][%d] getRequestsRequestNameOK  %+v", 200, o.Payload)
}

func (o *GetRequestsRequestNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestsRequestNameNotFound creates a GetRequestsRequestNameNotFound with default headers values
func NewGetRequestsRequestNameNotFound() *GetRequestsRequestNameNotFound {
	return &GetRequestsRequestNameNotFound{}
}

/*GetRequestsRequestNameNotFound handles this case with default header values.

The symbol request as specified by the request name was not found.
*/
type GetRequestsRequestNameNotFound struct {
	Payload *models.Request
}

func (o *GetRequestsRequestNameNotFound) Error() string {
	return fmt.Sprintf("[GET /{organization}/_apis/symbol/requests][%d] getRequestsRequestNameNotFound  %+v", 404, o.Payload)
}

func (o *GetRequestsRequestNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
