// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRequestsRequestNameReader is a Reader for the DeleteRequestsRequestName structure.
type DeleteRequestsRequestNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRequestsRequestNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteRequestsRequestNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRequestsRequestNameOK creates a DeleteRequestsRequestNameOK with default headers values
func NewDeleteRequestsRequestNameOK() *DeleteRequestsRequestNameOK {
	return &DeleteRequestsRequestNameOK{}
}

/*DeleteRequestsRequestNameOK handles this case with default header values.

The request no longer exists.
*/
type DeleteRequestsRequestNameOK struct {
}

func (o *DeleteRequestsRequestNameOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/_apis/symbol/requests][%d] deleteRequestsRequestNameOK ", 200)
}

func (o *DeleteRequestsRequestNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}