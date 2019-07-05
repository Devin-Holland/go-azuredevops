// Code generated by go-swagger; DO NOT EDIT.

package authorizedresources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// AuthorizeProjectResourcesReader is a Reader for the AuthorizeProjectResources structure.
type AuthorizeProjectResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizeProjectResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAuthorizeProjectResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthorizeProjectResourcesOK creates a AuthorizeProjectResourcesOK with default headers values
func NewAuthorizeProjectResourcesOK() *AuthorizeProjectResourcesOK {
	return &AuthorizeProjectResourcesOK{}
}

/*AuthorizeProjectResourcesOK handles this case with default header values.

successful operation
*/
type AuthorizeProjectResourcesOK struct {
	Payload []*models.DefinitionResourceReference
}

func (o *AuthorizeProjectResourcesOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/{project}/_apis/build/authorizedresources][%d] authorizeProjectResourcesOK  %+v", 200, o.Payload)
}

func (o *AuthorizeProjectResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}