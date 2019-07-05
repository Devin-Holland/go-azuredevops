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

// GetReleaseTaskAttachmentContentReader is a Reader for the GetReleaseTaskAttachmentContent structure.
type GetReleaseTaskAttachmentContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseTaskAttachmentContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReleaseTaskAttachmentContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReleaseTaskAttachmentContentOK creates a GetReleaseTaskAttachmentContentOK with default headers values
func NewGetReleaseTaskAttachmentContentOK() *GetReleaseTaskAttachmentContentOK {
	return &GetReleaseTaskAttachmentContentOK{}
}

/*GetReleaseTaskAttachmentContentOK handles this case with default header values.

successful operation
*/
type GetReleaseTaskAttachmentContentOK struct {
	Payload string
}

func (o *GetReleaseTaskAttachmentContentOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/release/releases/{releaseId}/environments/{environmentId}/attempts/{attemptId}/plan/{planId}/timelines/{timelineId}/records/{recordId}/attachments/{type}/{name}][%d] getReleaseTaskAttachmentContentOK  %+v", 200, o.Payload)
}

func (o *GetReleaseTaskAttachmentContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}