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

// UploadChunkReader is a Reader for the UploadChunk structure.
type UploadChunkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadChunkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadChunkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadChunkOK creates a UploadChunkOK with default headers values
func NewUploadChunkOK() *UploadChunkOK {
	return &UploadChunkOK{}
}

/*UploadChunkOK handles this case with default header values.

successful operation
*/
type UploadChunkOK struct {
	Payload *models.AttachmentReference
}

func (o *UploadChunkOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/{project}/_apis/wit/attachments/{id}][%d] uploadChunkOK  %+v", 200, o.Payload)
}

func (o *UploadChunkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttachmentReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
