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

// UpdateUserEntitlementReader is a Reader for the UpdateUserEntitlement structure.
type UpdateUserEntitlementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserEntitlementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateUserEntitlementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserEntitlementOK creates a UpdateUserEntitlementOK with default headers values
func NewUpdateUserEntitlementOK() *UpdateUserEntitlementOK {
	return &UpdateUserEntitlementOK{}
}

/*UpdateUserEntitlementOK handles this case with default header values.

successful operation
*/
type UpdateUserEntitlementOK struct {
	Payload *models.UserEntitlementsPatchResponse
}

func (o *UpdateUserEntitlementOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/_apis/userentitlements/{userId}][%d] updateUserEntitlementOK  %+v", 200, o.Payload)
}

func (o *UpdateUserEntitlementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntitlementsPatchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
