// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetReader is a Reader for the Get structure.
type GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOK creates a GetOK with default headers values
func NewGetOK() *GetOK {
	return &GetOK{}
}

/*GetOK handles this case with default header values.

successful operation
*/
type GetOK struct {
	Payload *models.WebAPITeam
}

func (o *GetOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/_apis/projects/{projectId}/teams/{teamId}][%d] getOK  %+v", 200, o.Payload)
}

func (o *GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebAPITeam)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
