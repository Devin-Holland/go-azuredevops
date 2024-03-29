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

// GetTestResultAttachmentZipReader is a Reader for the GetTestResultAttachmentZip structure.
type GetTestResultAttachmentZipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestResultAttachmentZipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTestResultAttachmentZipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTestResultAttachmentZipOK creates a GetTestResultAttachmentZipOK with default headers values
func NewGetTestResultAttachmentZipOK() *GetTestResultAttachmentZipOK {
	return &GetTestResultAttachmentZipOK{}
}

/*GetTestResultAttachmentZipOK handles this case with default header values.

successful operation
*/
type GetTestResultAttachmentZipOK struct {
	Payload string
}

func (o *GetTestResultAttachmentZipOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/test/Runs/{runId}/Results/{testCaseResultId}/attachments/{attachmentId}][%d] getTestResultAttachmentZipOK  %+v", 200, o.Payload)
}

func (o *GetTestResultAttachmentZipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
