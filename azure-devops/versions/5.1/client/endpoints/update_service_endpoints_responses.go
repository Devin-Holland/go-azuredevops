// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// UpdateServiceEndpointsReader is a Reader for the UpdateServiceEndpoints structure.
type UpdateServiceEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateServiceEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateServiceEndpointsOK creates a UpdateServiceEndpointsOK with default headers values
func NewUpdateServiceEndpointsOK() *UpdateServiceEndpointsOK {
	return &UpdateServiceEndpointsOK{}
}

/*UpdateServiceEndpointsOK handles this case with default header values.

successful operation
*/
type UpdateServiceEndpointsOK struct {
	Payload []*models.ServiceEndpoint
}

func (o *UpdateServiceEndpointsOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/{project}/_apis/serviceendpoint/endpoints][%d] updateServiceEndpointsOK  %+v", 200, o.Payload)
}

func (o *UpdateServiceEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
