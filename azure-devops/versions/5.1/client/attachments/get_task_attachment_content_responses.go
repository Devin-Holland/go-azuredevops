// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetTaskAttachmentContentReader is a Reader for the GetTaskAttachmentContent structure.
type GetTaskAttachmentContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskAttachmentContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTaskAttachmentContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTaskAttachmentContentOK creates a GetTaskAttachmentContentOK with default headers values
func NewGetTaskAttachmentContentOK() *GetTaskAttachmentContentOK {
	return &GetTaskAttachmentContentOK{}
}

/*GetTaskAttachmentContentOK handles this case with default header values.

successful operation
*/
type GetTaskAttachmentContentOK struct {
	Payload string
}

func (o *GetTaskAttachmentContentOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/release/releases/{releaseId}/environments/{environmentId}/attempts/{attemptId}/timelines/{timelineId}/records/{recordId}/attachments/{type}/{name}][%d] getTaskAttachmentContentOK  %+v", 200, o.Payload)
}

func (o *GetTaskAttachmentContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
