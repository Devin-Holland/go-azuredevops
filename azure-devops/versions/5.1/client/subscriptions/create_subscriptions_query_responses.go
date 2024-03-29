// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// CreateSubscriptionsQueryReader is a Reader for the CreateSubscriptionsQuery structure.
type CreateSubscriptionsQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionsQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSubscriptionsQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSubscriptionsQueryOK creates a CreateSubscriptionsQueryOK with default headers values
func NewCreateSubscriptionsQueryOK() *CreateSubscriptionsQueryOK {
	return &CreateSubscriptionsQueryOK{}
}

/*CreateSubscriptionsQueryOK handles this case with default header values.

successful operation
*/
type CreateSubscriptionsQueryOK struct {
	Payload *models.SubscriptionsQuery
}

func (o *CreateSubscriptionsQueryOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/_apis/hooks/subscriptionsquery][%d] createSubscriptionsQueryOK  %+v", 200, o.Payload)
}

func (o *CreateSubscriptionsQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionsQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
