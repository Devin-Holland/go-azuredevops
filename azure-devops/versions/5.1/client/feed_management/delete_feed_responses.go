// Code generated by go-swagger; DO NOT EDIT.

package feed_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteFeedReader is a Reader for the DeleteFeed structure.
type DeleteFeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteFeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFeedOK creates a DeleteFeedOK with default headers values
func NewDeleteFeedOK() *DeleteFeedOK {
	return &DeleteFeedOK{}
}

/*DeleteFeedOK handles this case with default header values.

successful operation
*/
type DeleteFeedOK struct {
}

func (o *DeleteFeedOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/{project}/_apis/packaging/feeds/{feedId}][%d] deleteFeedOK ", 200)
}

func (o *DeleteFeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}