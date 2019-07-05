// Code generated by go-swagger; DO NOT EDIT.

package source_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// ListRepositoriesReader is a Reader for the ListRepositories structure.
type ListRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRepositoriesOK creates a ListRepositoriesOK with default headers values
func NewListRepositoriesOK() *ListRepositoriesOK {
	return &ListRepositoriesOK{}
}

/*ListRepositoriesOK handles this case with default header values.

successful operation
*/
type ListRepositoriesOK struct {
	Payload *models.SourceRepositories
}

func (o *ListRepositoriesOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/sourceProviders/{providerName}/repositories][%d] listRepositoriesOK  %+v", 200, o.Payload)
}

func (o *ListRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SourceRepositories)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}