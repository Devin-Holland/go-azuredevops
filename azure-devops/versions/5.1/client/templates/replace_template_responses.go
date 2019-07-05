// Code generated by go-swagger; DO NOT EDIT.

package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// ReplaceTemplateReader is a Reader for the ReplaceTemplate structure.
type ReplaceTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceTemplateOK creates a ReplaceTemplateOK with default headers values
func NewReplaceTemplateOK() *ReplaceTemplateOK {
	return &ReplaceTemplateOK{}
}

/*ReplaceTemplateOK handles this case with default header values.

successful operation
*/
type ReplaceTemplateOK struct {
	Payload *models.WorkItemTemplate
}

func (o *ReplaceTemplateOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/{project}/{team}/_apis/wit/templates/{templateId}][%d] replaceTemplateOK  %+v", 200, o.Payload)
}

func (o *ReplaceTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkItemTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
