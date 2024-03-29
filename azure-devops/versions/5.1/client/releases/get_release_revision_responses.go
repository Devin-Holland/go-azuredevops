// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetReleaseRevisionReader is a Reader for the GetReleaseRevision structure.
type GetReleaseRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReleaseRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReleaseRevisionOK creates a GetReleaseRevisionOK with default headers values
func NewGetReleaseRevisionOK() *GetReleaseRevisionOK {
	return &GetReleaseRevisionOK{}
}

/*GetReleaseRevisionOK handles this case with default header values.

successful operation
*/
type GetReleaseRevisionOK struct {
	Payload string
}

func (o *GetReleaseRevisionOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/release/releases/{releaseId}][%d] getReleaseRevisionOK  %+v", 200, o.Payload)
}

func (o *GetReleaseRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
