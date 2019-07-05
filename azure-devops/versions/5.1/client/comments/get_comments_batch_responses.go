// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetCommentsBatchReader is a Reader for the GetCommentsBatch structure.
type GetCommentsBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCommentsBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCommentsBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCommentsBatchOK creates a GetCommentsBatchOK with default headers values
func NewGetCommentsBatchOK() *GetCommentsBatchOK {
	return &GetCommentsBatchOK{}
}

/*GetCommentsBatchOK handles this case with default header values.

successful operation
*/
type GetCommentsBatchOK struct {
	Payload *models.CommentList
}

func (o *GetCommentsBatchOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/wit/workItems/{workItemId}/comments][%d] getCommentsBatchOK  %+v", 200, o.Payload)
}

func (o *GetCommentsBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommentList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}