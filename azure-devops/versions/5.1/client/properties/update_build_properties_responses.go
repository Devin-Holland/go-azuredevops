// Code generated by go-swagger; DO NOT EDIT.

package properties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// UpdateBuildPropertiesReader is a Reader for the UpdateBuildProperties structure.
type UpdateBuildPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBuildPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateBuildPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateBuildPropertiesOK creates a UpdateBuildPropertiesOK with default headers values
func NewUpdateBuildPropertiesOK() *UpdateBuildPropertiesOK {
	return &UpdateBuildPropertiesOK{}
}

/*UpdateBuildPropertiesOK handles this case with default header values.

successful operation
*/
type UpdateBuildPropertiesOK struct {
	Payload *models.PropertiesCollection
}

func (o *UpdateBuildPropertiesOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/{project}/_apis/build/builds/{buildId}/properties][%d] updateBuildPropertiesOK  %+v", 200, o.Payload)
}

func (o *UpdateBuildPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertiesCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
