// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HasPermissionsReader is a Reader for the HasPermissions structure.
type HasPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HasPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHasPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHasPermissionsOK creates a HasPermissionsOK with default headers values
func NewHasPermissionsOK() *HasPermissionsOK {
	return &HasPermissionsOK{}
}

/*HasPermissionsOK handles this case with default header values.

successful operation
*/
type HasPermissionsOK struct {
	Payload []bool
}

func (o *HasPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/_apis/permissions/{securityNamespaceId}/{permissions}][%d] hasPermissionsOK  %+v", 200, o.Payload)
}

func (o *HasPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
