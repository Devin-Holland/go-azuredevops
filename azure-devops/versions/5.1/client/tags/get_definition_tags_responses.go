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

// GetDefinitionTagsReader is a Reader for the GetDefinitionTags structure.
type GetDefinitionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefinitionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDefinitionTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDefinitionTagsOK creates a GetDefinitionTagsOK with default headers values
func NewGetDefinitionTagsOK() *GetDefinitionTagsOK {
	return &GetDefinitionTagsOK{}
}

/*GetDefinitionTagsOK handles this case with default header values.

successful operation
*/
type GetDefinitionTagsOK struct {
	Payload []string
}

func (o *GetDefinitionTagsOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/build/definitions/{DefinitionId}/tags][%d] getDefinitionTagsOK  %+v", 200, o.Payload)
}

func (o *GetDefinitionTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
