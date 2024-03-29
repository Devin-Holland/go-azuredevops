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

// RemovePermissionReader is a Reader for the RemovePermission structure.
type RemovePermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemovePermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemovePermissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemovePermissionOK creates a RemovePermissionOK with default headers values
func NewRemovePermissionOK() *RemovePermissionOK {
	return &RemovePermissionOK{}
}

/*RemovePermissionOK handles this case with default header values.

successful operation
*/
type RemovePermissionOK struct {
	Payload *models.AccessControlEntry
}

func (o *RemovePermissionOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/_apis/permissions/{securityNamespaceId}/{permissions}][%d] removePermissionOK  %+v", 200, o.Payload)
}

func (o *RemovePermissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessControlEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
