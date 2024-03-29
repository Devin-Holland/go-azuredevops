// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// ReplaceAgentReader is a Reader for the ReplaceAgent structure.
type ReplaceAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceAgentOK creates a ReplaceAgentOK with default headers values
func NewReplaceAgentOK() *ReplaceAgentOK {
	return &ReplaceAgentOK{}
}

/*ReplaceAgentOK handles this case with default header values.

successful operation
*/
type ReplaceAgentOK struct {
	Payload *models.TaskAgent
}

func (o *ReplaceAgentOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/_apis/distributedtask/pools/{poolId}/agents/{agentId}][%d] replaceAgentOK  %+v", 200, o.Payload)
}

func (o *ReplaceAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
