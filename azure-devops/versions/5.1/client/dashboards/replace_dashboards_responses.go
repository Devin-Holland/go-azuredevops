// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// ReplaceDashboardsReader is a Reader for the ReplaceDashboards structure.
type ReplaceDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceDashboardsOK creates a ReplaceDashboardsOK with default headers values
func NewReplaceDashboardsOK() *ReplaceDashboardsOK {
	return &ReplaceDashboardsOK{}
}

/*ReplaceDashboardsOK handles this case with default header values.

successful operation
*/
type ReplaceDashboardsOK struct {
	Payload *models.DashboardGroup
}

func (o *ReplaceDashboardsOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/{project}/{team}/_apis/dashboard/dashboards][%d] replaceDashboardsOK  %+v", 200, o.Payload)
}

func (o *ReplaceDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
