// Code generated by go-swagger; DO NOT EDIT.

package feed_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetFeedViewReader is a Reader for the GetFeedView structure.
type GetFeedViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFeedViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFeedViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFeedViewOK creates a GetFeedViewOK with default headers values
func NewGetFeedViewOK() *GetFeedViewOK {
	return &GetFeedViewOK{}
}

/*GetFeedViewOK handles this case with default header values.

successful operation
*/
type GetFeedViewOK struct {
	Payload *models.FeedView
}

func (o *GetFeedViewOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/packaging/Feeds/{feedId}/views/{viewId}][%d] getFeedViewOK  %+v", 200, o.Payload)
}

func (o *GetFeedViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FeedView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
