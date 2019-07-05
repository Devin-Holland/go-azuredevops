// Code generated by go-swagger; DO NOT EDIT.

package group_entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// UpdateReader is a Reader for the Update structure.
type UpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOK creates a UpdateOK with default headers values
func NewUpdateOK() *UpdateOK {
	return &UpdateOK{}
}

/*UpdateOK handles this case with default header values.

successful operation
*/
type UpdateOK struct {
	Payload *models.GroupEntitlementOperationReference
}

func (o *UpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/_apis/groupentitlements/{groupId}][%d] updateOK  %+v", 200, o.Payload)
}

func (o *UpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupEntitlementOperationReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
