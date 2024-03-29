// Code generated by go-swagger; DO NOT EDIT.

package feed_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteFeedViewReader is a Reader for the DeleteFeedView structure.
type DeleteFeedViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFeedViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteFeedViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFeedViewOK creates a DeleteFeedViewOK with default headers values
func NewDeleteFeedViewOK() *DeleteFeedViewOK {
	return &DeleteFeedViewOK{}
}

/*DeleteFeedViewOK handles this case with default header values.

successful operation
*/
type DeleteFeedViewOK struct {
}

func (o *DeleteFeedViewOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/{project}/_apis/packaging/Feeds/{feedId}/views/{viewId}][%d] deleteFeedViewOK ", 200)
}

func (o *DeleteFeedViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
