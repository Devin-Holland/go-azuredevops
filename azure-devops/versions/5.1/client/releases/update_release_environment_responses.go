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

// UpdateReleaseEnvironmentReader is a Reader for the UpdateReleaseEnvironment structure.
type UpdateReleaseEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReleaseEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateReleaseEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateReleaseEnvironmentOK creates a UpdateReleaseEnvironmentOK with default headers values
func NewUpdateReleaseEnvironmentOK() *UpdateReleaseEnvironmentOK {
	return &UpdateReleaseEnvironmentOK{}
}

/*UpdateReleaseEnvironmentOK handles this case with default header values.

successful operation
*/
type UpdateReleaseEnvironmentOK struct {
	Payload *models.ReleaseEnvironment
}

func (o *UpdateReleaseEnvironmentOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/{project}/_apis/Release/releases/{releaseId}/environments/{environmentId}][%d] updateReleaseEnvironmentOK  %+v", 200, o.Payload)
}

func (o *UpdateReleaseEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseEnvironment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
