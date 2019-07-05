// Code generated by go-swagger; DO NOT EDIT.

package widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetWidgetsReader is a Reader for the GetWidgets structure.
type GetWidgetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWidgetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWidgetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWidgetsOK creates a GetWidgetsOK with default headers values
func NewGetWidgetsOK() *GetWidgetsOK {
	return &GetWidgetsOK{}
}

/*GetWidgetsOK handles this case with default header values.

successful operation
*/
type GetWidgetsOK struct {
	ETag string

	Payload []*models.Widget
}

func (o *GetWidgetsOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/{team}/_apis/dashboard/dashboards/{dashboardId}/widgets][%d] getWidgetsOK  %+v", 200, o.Payload)
}

func (o *GetWidgetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
