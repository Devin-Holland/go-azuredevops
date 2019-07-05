// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetReleaseTaskAttachmentsReader is a Reader for the GetReleaseTaskAttachments structure.
type GetReleaseTaskAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseTaskAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReleaseTaskAttachmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReleaseTaskAttachmentsOK creates a GetReleaseTaskAttachmentsOK with default headers values
func NewGetReleaseTaskAttachmentsOK() *GetReleaseTaskAttachmentsOK {
	return &GetReleaseTaskAttachmentsOK{}
}

/*GetReleaseTaskAttachmentsOK handles this case with default header values.

successful operation
*/
type GetReleaseTaskAttachmentsOK struct {
	Payload []*models.ReleaseTaskAttachment
}

func (o *GetReleaseTaskAttachmentsOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/release/releases/{releaseId}/environments/{environmentId}/attempts/{attemptId}/plan/{planId}/attachments/{type}][%d] getReleaseTaskAttachmentsOK  %+v", 200, o.Payload)
}

func (o *GetReleaseTaskAttachmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
