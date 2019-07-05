// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetTaskLogReader is a Reader for the GetTaskLog structure.
type GetTaskLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTaskLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTaskLogOK creates a GetTaskLogOK with default headers values
func NewGetTaskLogOK() *GetTaskLogOK {
	return &GetTaskLogOK{}
}

/*GetTaskLogOK handles this case with default header values.

successful operation
*/
type GetTaskLogOK struct {
	Payload string
}

func (o *GetTaskLogOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/release/releases/{releaseId}/environments/{environmentId}/deployPhases/{releaseDeployPhaseId}/tasks/{taskId}/logs][%d] getTaskLogOK  %+v", 200, o.Payload)
}

func (o *GetTaskLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
