// Code generated by go-swagger; DO NOT EDIT.

package user_entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// UpdateUserEntitlementsReader is a Reader for the UpdateUserEntitlements structure.
type UpdateUserEntitlementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserEntitlementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateUserEntitlementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserEntitlementsOK creates a UpdateUserEntitlementsOK with default headers values
func NewUpdateUserEntitlementsOK() *UpdateUserEntitlementsOK {
	return &UpdateUserEntitlementsOK{}
}

/*UpdateUserEntitlementsOK handles this case with default header values.

successful operation
*/
type UpdateUserEntitlementsOK struct {
	Payload *models.UserEntitlementOperationReference
}

func (o *UpdateUserEntitlementsOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/_apis/userentitlements][%d] updateUserEntitlementsOK  %+v", 200, o.Payload)
}

func (o *UpdateUserEntitlementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntitlementOperationReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}