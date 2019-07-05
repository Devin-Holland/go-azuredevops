// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetBuildTagsReader is a Reader for the GetBuildTags structure.
type GetBuildTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBuildTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBuildTagsOK creates a GetBuildTagsOK with default headers values
func NewGetBuildTagsOK() *GetBuildTagsOK {
	return &GetBuildTagsOK{}
}

/*GetBuildTagsOK handles this case with default header values.

successful operation
*/
type GetBuildTagsOK struct {
	Payload []string
}

func (o *GetBuildTagsOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/build/builds/{buildId}/tags][%d] getBuildTagsOK  %+v", 200, o.Payload)
}

func (o *GetBuildTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
