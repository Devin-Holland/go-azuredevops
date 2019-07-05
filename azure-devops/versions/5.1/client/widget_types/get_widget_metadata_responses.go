// Code generated by go-swagger; DO NOT EDIT.

package widget_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetWidgetMetadataReader is a Reader for the GetWidgetMetadata structure.
type GetWidgetMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWidgetMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWidgetMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWidgetMetadataOK creates a GetWidgetMetadataOK with default headers values
func NewGetWidgetMetadataOK() *GetWidgetMetadataOK {
	return &GetWidgetMetadataOK{}
}

/*GetWidgetMetadataOK handles this case with default header values.

successful operation
*/
type GetWidgetMetadataOK struct {
	Payload *models.WidgetMetadataResponse
}

func (o *GetWidgetMetadataOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/dashboard/widgettypes/{contributionId}][%d] getWidgetMetadataOK  %+v", 200, o.Payload)
}

func (o *GetWidgetMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WidgetMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
