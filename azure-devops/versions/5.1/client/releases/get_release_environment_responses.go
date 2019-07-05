// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetReleaseEnvironmentReader is a Reader for the GetReleaseEnvironment structure.
type GetReleaseEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReleaseEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReleaseEnvironmentOK creates a GetReleaseEnvironmentOK with default headers values
func NewGetReleaseEnvironmentOK() *GetReleaseEnvironmentOK {
	return &GetReleaseEnvironmentOK{}
}

/*GetReleaseEnvironmentOK handles this case with default header values.

successful operation
*/
type GetReleaseEnvironmentOK struct {
	Payload *models.ReleaseEnvironment
}

func (o *GetReleaseEnvironmentOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/Release/releases/{releaseId}/environments/{environmentId}][%d] getReleaseEnvironmentOK  %+v", 200, o.Payload)
}

func (o *GetReleaseEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseEnvironment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
