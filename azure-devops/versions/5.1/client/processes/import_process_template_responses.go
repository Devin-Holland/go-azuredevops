// Code generated by go-swagger; DO NOT EDIT.

package processes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// ImportProcessTemplateReader is a Reader for the ImportProcessTemplate structure.
type ImportProcessTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportProcessTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImportProcessTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImportProcessTemplateOK creates a ImportProcessTemplateOK with default headers values
func NewImportProcessTemplateOK() *ImportProcessTemplateOK {
	return &ImportProcessTemplateOK{}
}

/*ImportProcessTemplateOK handles this case with default header values.

successful operation
*/
type ImportProcessTemplateOK struct {
	Payload *models.ProcessImportResult
}

func (o *ImportProcessTemplateOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/_apis/work/processadmin/processes/import][%d] importProcessTemplateOK  %+v", 200, o.Payload)
}

func (o *ImportProcessTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessImportResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
