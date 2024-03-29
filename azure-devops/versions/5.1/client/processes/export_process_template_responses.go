// Code generated by go-swagger; DO NOT EDIT.

package processes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ExportProcessTemplateReader is a Reader for the ExportProcessTemplate structure.
type ExportProcessTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportProcessTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExportProcessTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExportProcessTemplateOK creates a ExportProcessTemplateOK with default headers values
func NewExportProcessTemplateOK() *ExportProcessTemplateOK {
	return &ExportProcessTemplateOK{}
}

/*ExportProcessTemplateOK handles this case with default header values.

successful operation
*/
type ExportProcessTemplateOK struct {
	Payload string
}

func (o *ExportProcessTemplateOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/_apis/work/processadmin/processes/export/{id}][%d] exportProcessTemplateOK  %+v", 200, o.Payload)
}

func (o *ExportProcessTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
